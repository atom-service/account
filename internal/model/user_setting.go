package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/atom-service/account/package/proto"
	"github.com/yinxulai/sqls"
)

var userSettingTableName = userSchemaName + ".\"settings\""
var SettingTable = &settingTable{}

type Setting struct {
	ID          *int64
	Key         string
	Value       *string
	UserID      int64
	Description *string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	DeletedTime *time.Time
}

func (srv *Setting) ToProto() *proto.Setting {
	setting := new(proto.Setting)
	setting.ID = *srv.ID
	setting.Key = srv.Key
	setting.UserID = srv.UserID

	if srv.ID != nil {
		setting.ID = *srv.ID
	}

	if srv.Value != nil {
		setting.Value = *srv.Value
	}

	if srv.DeletedTime != nil {
		timeString := srv.DeletedTime.UTC().String()
		setting.DeletedTime = &timeString
	}

	if srv.CreatedTime != nil {
		timeString := srv.CreatedTime.UTC().String()
		setting.CreatedTime = timeString
	}

	if srv.UpdatedTime != nil {
		timeString := srv.UpdatedTime.UTC().String()
		setting.UpdatedTime = timeString
	}

	return setting
}

type settingTable struct{}

type SettingSelector struct {
	ID     *int64
	Key    *string
	UserID *int64
}

func (sel *SettingSelector) LoadProto(data *proto.SettingSelector) {
	if data == nil {
		return
	}

	sel.ID = data.ID
	sel.Key = data.Key
	sel.UserID = data.UserID
}

func (t *settingTable) InitTable(ctx context.Context) error {
	tx, err := Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	cs := sqls.CREATE_SCHEMA(userSchemaName).IF_NOT_EXISTS()
	if _, err = tx.ExecContext(ctx, cs.String()); err != nil {
		tx.Rollback()
		return err
	}
	// 创建 table
	s := sqls.CREATE_TABLE(userSettingTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial PRIMARY KEY NOT NULL")
	s.COLUMN("user_id int NOT NULL")
	s.COLUMN("key character varying(64) NOT NULL")
	s.COLUMN("value character varying(513) NOT NULL")
	s.COLUMN("description character varying(128) NULL")
	s.COLUMN("created_time timestamp with time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp with time zone NULL DEFAULT now()")
	s.COLUMN("deleted_time timestamp with time zone NULL")
	s.OPTIONS("CONSTRAINT user_setting_union_unique_keys UNIQUE (user_id, key)")
	slog.DebugContext(ctx, s.String())

	if _, err = tx.ExecContext(ctx, s.String()); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Join(err)
		}
		return err
	}

	return nil
}

func (r *settingTable) CreateSetting(ctx context.Context, newSetting Setting) (err error) {
	s := sqls.INSERT_INTO(userSettingTableName)
	s.VALUES("key", s.Param(newSetting.Key))
	s.VALUES("value", s.Param(newSetting.Value))
	s.VALUES("user_id", s.Param(newSetting.UserID))
	s.VALUES("description", s.Param(newSetting.Description))

	slog.DebugContext(ctx, s.String())
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		slog.ErrorContext(ctx, "CreateSetting failed", err)
		return
	}

	return
}

func (r *settingTable) UpdateSetting(ctx context.Context, selector SettingSelector, role *Setting) (err error) {
	s := sqls.UPDATE(userSettingTableName)

	if selector.ID == nil && selector.UserID == nil && selector.Key == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}

	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	s.SET("key", s.Param(role.Key))

	if role.Value != nil {
		s.SET("value", s.Param(*role.Value))
	}

	if role.Description != nil {
		s.SET("description", s.Param(*role.Description))
	}

	if role.DeletedTime != nil {
		s.SET("deleted_time", s.Param(*role.DeletedTime))
	}

	s.SET("updated_time", "CURRENT_TIMESTAMP")

	slog.DebugContext(ctx, s.String(), slog.Any("params", s.Params()))
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		slog.ErrorContext(ctx, "UpdateSetting failed", err)
	}

	return
}

func (r *settingTable) DeleteSetting(ctx context.Context, selector SettingSelector) (err error) {
	s := sqls.UPDATE(userSettingTableName)

	if selector.ID == nil && selector.UserID == nil && selector.Key == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}

	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	s.SET("deleted_time", "CURRENT_TIMESTAMP")

	slog.DebugContext(ctx, s.String(), slog.Any("params", s.Params()))
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		slog.ErrorContext(ctx, "DeleteSetting failed", err)
	}

	return
}

func (r *settingTable) CountSettings(ctx context.Context, selector SettingSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(userSettingTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}

	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	s.WHERE("(deleted_time>=CURRENT_TIMESTAMP OR deleted_time IS NULL)")

	slog.DebugContext(ctx, s.String())
	rowQuery := Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		slog.ErrorContext(ctx, "CountSettings failed", err)
	}

	return
}

func (r *settingTable) QuerySettings(ctx context.Context, selector SettingSelector, pagination *Pagination, sort *Sort) (result []*Setting, err error) {
	s := sqls.SELECT(
		"id",
		"key",
		"value",
		"user_id",
		"description",
		"created_time",
		"updated_time",
		"deleted_time",
	).FROM(userSettingTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}

	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	s.WHERE("(deleted_time>=CURRENT_TIMESTAMP OR deleted_time IS NULL)")

	if pagination == nil {
		pagination = &Pagination{}
	}

	if pagination.Limit == nil {
		// 默认为 100，防止刷爆
		defaultLimit := int64(100)
		pagination.Limit = &defaultLimit
	}

	s.LIMIT(s.Param(pagination.Limit))

	if pagination.Offset != nil {
		s.OFFSET(s.Param(pagination.Offset))
	}

	if sort != nil {
		var sortType = "ASC"
		if sort.Type == SortDesc {
			sortType = "DESC"
		}

		s.ORDER_BY(s.Param(sort.Key) + " " + sortType)
	}

	queryResult, err := Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		slog.ErrorContext(ctx, "QuerySettings failed", err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		setting := Setting{}
		if err = queryResult.Scan(
			&setting.ID,
			&setting.Key,
			&setting.Value,
			&setting.UserID,
			&setting.Description,
			&setting.CreatedTime,
			&setting.UpdatedTime,
			&setting.DeletedTime,
		); err != nil {
			slog.ErrorContext(ctx, "QuerySettings failed", err)
			return
		}
		result = append(result, &setting)
	}
	if err = queryResult.Err(); err != nil {
		slog.ErrorContext(ctx, "QuerySettings failed", err)
		return
	}
	return
}
