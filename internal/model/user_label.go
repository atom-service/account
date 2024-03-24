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

var userLabelTableName = userSchemaName + ".\"labels\""
var LabelTable = &labelTable{}

const (
	LabelLastSignInTime = "last_sign_in_time" // 最近一次登录时间
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

type LabelSelector struct {
	ID     *int64
	Key    *string
	UserID *int64
}

type labelTable struct{}

func (t *labelTable) CreateTable(ctx context.Context) error {
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

func (t *labelTable) TruncateTable(ctx context.Context) error {
	_, err := database.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(userLabelTableName).String())
	return err
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

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
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

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
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

	logger.Debug(s.String())
	rowQuery := database.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
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

	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
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
			logger.Error(err)
			return
		}
		result = append(result, &label)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}
