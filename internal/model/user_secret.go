package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/package/proto"
	"github.com/atom-service/common/logger"
	"github.com/yinxulai/sqls"
)

var secretTableName = userSchemaName + ".\"secrets\""

var (
	UserSecretType   = "user"
	SystemSecretType = "system"
)

type Secret struct {
	Key          *string
	Type         *string
	Value        *string
	UserID       *int64
	Description  *string
	CreatedTime  *time.Time
	UpdatedTime  *time.Time
	DeletedTime  *time.Time
	DisabledTime *time.Time
}

func (srv *Secret) LoadProto(secret *proto.Secret) (err error) {
	if secret == nil {
		return nil
	}

	srv.Key = &secret.Key
	srv.Value = &secret.Value
	srv.UserID = &secret.UserID
	srv.Description = &secret.Description

	createdTime, err := time.Parse(time.RFC3339Nano, secret.CreatedTime)
	if err != nil {
		return err
	}

	srv.CreatedTime = &createdTime

	updatedTime, err := time.Parse(time.RFC3339Nano, secret.UpdatedTime)
	if err != nil {
		return err
	}

	srv.UpdatedTime = &updatedTime

	if secret.DeletedTime != nil {
		deletedTime, err := time.Parse(time.RFC3339Nano, *secret.DeletedTime)
		if err != nil {
			return err
		}

		srv.DeletedTime = &deletedTime
	}

	return
}

func (srv *Secret) ToProto() *proto.Secret {
	secret := new(proto.Secret)
	if srv.Key != nil {
		secret.Key = *srv.Key
	}
	if srv.Value != nil {
		secret.Value = *srv.Value
	}
	if srv.UserID != nil {
		secret.UserID = *srv.UserID
	}
	if srv.Description != nil {
		secret.Description = *srv.Description
	}

	if srv.CreatedTime != nil {
		secret.CreatedTime = srv.CreatedTime.String()
	}

	if srv.UpdatedTime != nil {
		secret.UpdatedTime = srv.UpdatedTime.String()
	}

	if srv.DeletedTime != nil {
		timeString := srv.DeletedTime.String()
		secret.DeletedTime = &timeString
	}

	return secret
}

func (srv *Secret) IsDisabled() bool {
	if srv.DisabledTime == nil {
		return false
	}

	if srv.DisabledTime.Before(time.Now()) {
		return true
	}

	return false
}

type SecretSelector struct {
	Key    *string
	Type   *string
	UserID *int64
}

func (srv *SecretSelector) LoadProto(data *proto.SecretSelector) {
	if data != nil {
		srv.Key = data.Key
		srv.UserID = data.UserID
	}
}

type secretTable struct{}

var SecretTable = &secretTable{}

func (s *secretTable) initData(ctx context.Context) (err error) {
	adminUserID := int64(1)
	selector := SecretSelector{UserID: &adminUserID}
	queryResult, err := s.QuerySecrets(ctx, selector, nil, nil)
	if err != nil {
		return err
	}

	if len(queryResult) == 0 {
		return s.CreateSecret(ctx, CreateSecretParams{
			UserID: adminUserID,
			Type:   SystemSecretType,
		})
	}

	return nil
}


func (t *secretTable) InitTable(ctx context.Context) error {
	tx, err := Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	cs := sqls.CREATE_SCHEMA(userSchemaName).IF_NOT_EXISTS()
	logger.Debug(cs.String())
	if _, err = tx.ExecContext(ctx, cs.String()); err != nil {
		tx.Rollback()
		return err
	}

	// 创建 table
	s := sqls.CREATE_TABLE(secretTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial PRIMARY KEY NOT NULL")
	s.COLUMN("key character varying(64) UNIQUE NOT NULL")
	s.COLUMN("type character varying(64) NOT NULL")
	s.COLUMN("value character varying(64) NOT NULL")
	s.COLUMN("user_id integer NOT NULL")
	s.COLUMN("description character varying(128) NULL")
	s.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("disabled_time timestamp without time zone NULL")
	s.COLUMN("deleted_time timestamp without time zone NULL")
	s.OPTIONS("CONSTRAINT user_secret_union_unique_keys UNIQUE (user_id, key, type)")
	logger.Debug(s.String())

	if _, err := tx.ExecContext(ctx, s.String()); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Join(err)
		}
		return err
	}

	return t.initData(ctx)
}

func (t *secretTable) TruncateTable(ctx context.Context) error {
	_, err := Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(secretTableName).String())
	return err
}

type CreateSecretParams struct {
	Type        string
	UserID      int64
	Description *string
}

func (r *secretTable) CreateSecret(ctx context.Context, createParams CreateSecretParams) (err error) {
	var key string

	// TODO 并发加快速度
	for {
		key = helper.GenerateRandomString(64, nil)

		count, err := r.CountSecrets(ctx, SecretSelector{Key: &key})
		if err != nil {
			return err
		}

		if count > 0 {
			continue
		} else {
			break
		}
	}

	value := helper.GenerateRandomString(64, nil)

	s := sqls.INSERT_INTO(secretTableName)
	s.VALUES("key", s.Param(key))
	s.VALUES("value", s.Param(value))
	s.VALUES("type", s.Param(createParams.Type))
	s.VALUES("user_id", s.Param(createParams.UserID))
	s.VALUES("description", s.Param(createParams.Description))

	logger.Debug(s.String(), s.Params())
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *secretTable) DeleteSecret(ctx context.Context, selector SecretSelector) (err error) {
	s := sqls.UPDATE(secretTableName)

	if selector.Key == nil && selector.UserID == nil && selector.Type == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
	}

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}
	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	if selector.Type != nil {
		s.WHERE("type=" + s.Param(selector.Type))
	}

	s.SET("deleted_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *secretTable) UpdateSecret(ctx context.Context, selector SecretSelector, user *Secret) (err error) {
	s := sqls.UPDATE(secretTableName)

	if selector.UserID == nil || selector.Key == nil {
		return fmt.Errorf("selector conditions cannot be empty")
	}

	s.WHERE("key=" + s.Param(selector.Key))
	s.WHERE("user_id=" + s.Param(selector.UserID))

	if user.DisabledTime != nil {
		s.SET("disabled_time", s.Param(*user.DisabledTime))
	}

	s.SET("updated_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *secretTable) CountSecrets(ctx context.Context, selector SecretSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(secretTableName)

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}
	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}
	if selector.Type != nil {
		s.WHERE("type=" + s.Param(selector.Type))
	}

	s.WHERE("(deleted_time<CURRENT_TIMESTAMP OR deleted_time IS NULL)")

	logger.Debug(s.String(), s.Params())
	rowQuery := Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *secretTable) QuerySecrets(ctx context.Context, selector SecretSelector, pagination *Pagination, sort *Sort) (result []*Secret, err error) {
	s := sqls.SELECT(
		"key",
		"type",
		"value",
		"user_id",
		"description",
		"created_time",
		"updated_time",
		"disabled_time",
		"deleted_time",
	).FROM(secretTableName)

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}
	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	if selector.Type != nil {
		s.WHERE("type=" + s.Param(selector.Type))
	}

	s.WHERE("(deleted_time<CURRENT_TIMESTAMP OR deleted_time IS NULL)")

	if pagination == nil {
		pagination = &Pagination{}
	}

	if pagination.Limit == nil {
		// 个人的 ak sk 默认上限 10
		defaultLimit := int64(10)
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

	logger.Debug(s.String(), s.Params())
	queryResult, err := Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		secret := Secret{}
		if err = queryResult.Scan(
			&secret.Key,
			&secret.Type,
			&secret.Value,
			&secret.UserID,
			&secret.Description,
			&secret.CreatedTime,
			&secret.UpdatedTime,
			&secret.DisabledTime,
			&secret.DeletedTime,
		); err != nil {
			logger.Error(err)
			return
		}
		result = append(result, &secret)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}
