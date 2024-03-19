package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/atom-service/account/internal/database"
	"github.com/atom-service/common/logger"
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

type SettingSelector struct {
	ID     *int64
	Key    *string
	UserID *int64
}

type settingTable struct{}

func (t *settingTable) CreateTable(ctx context.Context) error {
	tx, err := database.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
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
	s.COLUMN("value character varying(128) NOT NULL")
	s.COLUMN("description character varying(128) NOT NULL")
	s.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("deleted_time timestamp without time zone NULL")
	s.OPTIONS("CONSTRAINT user_setting_union_unique_keys UNIQUE (user_id, key)")
	logger.Debug(s.String())

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

func (t *settingTable) TruncateTable(ctx context.Context) error {
	_, err := database.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(userSettingTableName).String())
	return err
}

func (r *settingTable) CreateSetting(ctx context.Context, newSetting Setting) (err error) {
	s := sqls.INSERT_INTO(userSettingTableName)
	s.VALUES("key", s.Param(newSetting.Key))
	s.VALUES("value", s.Param(newSetting.Value))
	s.VALUES("user_id", s.Param(newSetting.UserID))
	s.VALUES("description", s.Param(newSetting.Description))

	logger.Debug(s.String())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *settingTable) UpdateSetting(ctx context.Context, selector SettingSelector, role *Setting) (err error) {
	s := sqls.UPDATE(userSettingTableName)

	if selector.ID == nil && selector.UserID == nil && selector.Key == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
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

	if role.Value != nil {
		s.SET("value", s.Param(*role.Value))
	}

	if role.Description != nil {
		s.SET("description", s.Param(*role.Description))
	}

	if role.DeletedTime != nil {
		s.SET("disabled_time", s.Param(*role.DeletedTime))
	}

	s.SET("updated_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *settingTable) DeleteSetting(ctx context.Context, selector SettingSelector) (err error) {
	s := sqls.UPDATE(userSettingTableName)

	if selector.ID == nil && selector.UserID == nil && selector.Key == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
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

	s.SET("deleted_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
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

	s.WHERE("(deleted_time<CURRENT_TIMESTAMP OR deleted_time IS NULL)")

	logger.Debug(s.String())
	rowQuery := database.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
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

	s.WHERE("(deleted_time<CURRENT_TIMESTAMP OR deleted_time IS NULL)")

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

	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
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
			logger.Error(err)
			return
		}
		result = append(result, &setting)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}
