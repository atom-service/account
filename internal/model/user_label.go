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

var userLabelTableName = userSchemaName + ".\"labels\""
var LabelTable = &labelTable{}

const (
	LabelLastSignInTime = "last_sign_in_time"     // 最近一次登录时间
	LabelLastVerifyTime = "last_verify_auth_time" // 最近一次身份验证时间
)

type Label struct {
	ID          *int64
	Key         string
	Value       *string
	UserID      int64
	Description *string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	DeletedTime *time.Time
}

func (srv *Label) ToProto() *proto.Label {
	secret := new(proto.Label)
	secret.Key = srv.Key
	secret.UserID = &srv.UserID

	if srv.Value != nil {
		secret.Value = *srv.Value
	}
	if srv.Description != nil {
		secret.Description = srv.Description
	}

	if srv.CreatedTime != nil {
		timeString := srv.CreatedTime.String()
		secret.CreatedTime = &timeString
	}

	if srv.UpdatedTime != nil {
		timeString := srv.UpdatedTime.String()
		secret.UpdatedTime = &timeString
	}

	if srv.DeletedTime != nil {
		timeString := srv.DeletedTime.String()
		secret.DeletedTime = &timeString
	}

	return secret
}

type LabelSelector struct {
	ID     *int64
	Key    *string
	UserID *int64
}

func (srv *LabelSelector) LoadProto(data *proto.LabelSelector) {
	if data != nil {
		srv.Key = data.Key
		srv.UserID = data.UserID
	}
}

type labelTable struct{}

func (t *labelTable) InitTable(ctx context.Context) error {
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
	s := sqls.CREATE_TABLE(userLabelTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial PRIMARY KEY NOT NULL")
	s.COLUMN("user_id int NOT NULL")
	s.COLUMN("key character varying(64) NOT NULL")
	s.COLUMN("value character varying(128) NOT NULL")
	s.COLUMN("description character varying(128) NULL")
	s.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("deleted_time timestamp without time zone NULL")
	s.OPTIONS("CONSTRAINT user_label_union_unique_keys UNIQUE (user_id, key)")
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

func (r *labelTable) UpsertLabel(ctx context.Context, newLabel Label) (err error) {
	s := sqls.INSERT_INTO(userLabelTableName)
	s.VALUES("key", s.Param(newLabel.Key))
	s.VALUES("value", s.Param(newLabel.Value))
	s.VALUES("user_id", s.Param(newLabel.UserID))
	s.VALUES("description", s.Param(newLabel.Description))
	s.ON_CONFLICT("user_id, key")
	s.DO_UPDATE_SET("value", s.Param(newLabel.Value))
	s.DO_UPDATE_SET("description", s.Param(newLabel.Description))

	slog.DebugContext(ctx, s.String(), s.Params())
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		slog.ErrorContext(ctx, "UpsertLabel failed", err)
		return
	}

	return
}

func (r *labelTable) DeleteLabel(ctx context.Context, selector LabelSelector) (err error) {
	s := sqls.UPDATE(userLabelTableName)

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

	s.SET("deleted_time", s.Param(time.Now()))

	slog.DebugContext(ctx, s.String(), s.Params())
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		slog.ErrorContext(ctx, "DeleteLabel failed", err)
	}

	return
}

func (r *labelTable) CountLabels(ctx context.Context, selector LabelSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(userLabelTableName)

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

	slog.DebugContext(ctx, s.String())
	rowQuery := Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		slog.ErrorContext(ctx, "CountLabels failed", err)
	}

	return
}

func (r *labelTable) QueryLabels(ctx context.Context, selector LabelSelector, pagination *Pagination, sort *Sort) (result []*Label, err error) {
	s := sqls.SELECT(
		"id",
		"key",
		"value",
		"user_id",
		"description",
		"created_time",
		"updated_time",
		"deleted_time",
	).FROM(userLabelTableName)

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

	queryResult, err := Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		slog.ErrorContext(ctx, "QueryLabels failed", err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		label := Label{}
		if err = queryResult.Scan(
			&label.ID,
			&label.Key,
			&label.Value,
			&label.UserID,
			&label.Description,
			&label.CreatedTime,
			&label.UpdatedTime,
			&label.DeletedTime,
		); err != nil {
			slog.ErrorContext(ctx, "QueryLabels failed", err)
			return
		}
		result = append(result, &label)
	}
	if err = queryResult.Err(); err != nil {
		slog.ErrorContext(ctx, "QueryLabels failed", err)
		return
	}
	return
}
