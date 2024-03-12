package model

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"math/big"
	"time"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/logger"
	"github.com/atom-service/common/sqls"
)

var secretTableName = "\"secret\".\"secrets\""

const (
	UserSecretType   = "user"
	SystemSecretType = "system"
)

type Secret struct {
	Key         *string    `json:"key"`
	Type        *string    `json:"type"`
	Value       *string    `json:"value"`
	OwnerID     *int64     `json:"owner_id"`
	Description *string    `json:"description"`
	CreatedTime *time.Time `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
	DeletedTime *time.Time `json:"deleted_time"`
}

func (srv *Secret) LoadProtoStruct(secret *protos.Secret) (err error) {
	srv.Key = &secret.Key
	srv.Value = &secret.Value
	srv.OwnerID = &secret.OwnerID
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

// OutProtoStruct OutProtoStruct
func (srv *Secret) OutProtoStruct() *protos.Secret {
	secret := new(protos.Secret)
	secret.Key = *srv.Key
	secret.Value = *srv.Value
	secret.OwnerID = *srv.OwnerID
	secret.Description = *srv.Description
	secret.CreatedTime = srv.CreatedTime.String()
	secret.UpdatedTime = srv.UpdatedTime.String()

	if srv.DeletedTime != nil {
		timeString := srv.DeletedTime.String()
		secret.DeletedTime = &timeString
	}

	return secret
}

type SecretSelector struct {
	Key     *string
	Type    string
	OwnerID *int64
}

func (srv *SecretSelector) LoadProtoStruct(data *protos.SecretSelector) {
	srv.Key = data.Key
	srv.OwnerID = data.OwnerID
}

// OutProtoStruct OutProtoStruct
func (srv *SecretSelector) OutProtoStruct() *protos.SecretSelector {
	result := new(protos.SecretSelector)
	result.Key = srv.Key
	result.OwnerID = srv.OwnerID
	return result
}

type secretTable struct{}

var SecretTable = &secretTable{}

func (t *secretTable) CreateTable(ctx context.Context) error {
	tx, err := db.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	tx.ExecContext(ctx, "CREATE SCHEMA IF NOT EXISTS \"secret\"")

	// 创建 table
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("CREATE TABLE IF NOT EXISTS", secretTableName, "(")
	s.Append("key character varying(128) NOT NULL,")
	s.Append("type character varying(128) NOT NULL,")
	s.Append("value character varying(128) NOT NULL,")
	s.Append("owner_id integer NOT NULL,")
	s.Append("description character varying(64) NULL,")
	s.Append("created_time timestamp without time zone NULL DEFAULT now(),")
	s.Append("updated_time timestamp without time zone NULL DEFAULT now(),")
	s.Append("deleted_time timestamp without time zone NULL")
	s.Append(");")
	logger.Debug(s.String())
	tx.ExecContext(ctx, s.String())
	if err := tx.Commit(); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Join(err)
		}
		return err
	}

	return nil
}

func (t *secretTable) TruncateTable(ctx context.Context) error {
	_, err := db.Database.ExecContext(ctx, "TRUNCATE TABLE", secretTableName, ";")
	return err
}

type CreateSecretParams struct {
	Type        string    `json:"type"`
	OwnerID     int64     `json:"owner_id"`
	Description *string    `json:"description"`
}

func (r *secretTable) CreateSecret(ctx context.Context, createParams CreateSecretParams) (err error) {
	generateRandomString := func(length int) (string, error) {
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		charsetLength := big.NewInt(int64(len(charset)))

		randomString := make([]byte, length)
		for i := 0; i < length; i++ {
			index, err := rand.Int(rand.Reader, charsetLength)
			if err != nil {
				return "", err
			}
			randomString[i] = charset[index.Int64()]
		}

		return string(randomString), nil
	}

	var key string

	// TODO 并发加快速度
	for {
		key, err = generateRandomString(128)
		if err != nil {
			return err
		}

		count, err := r.CountSecrets(ctx, SecretSelector{Key: &key})
		if err != nil {
			return err
		}

		if (count > 0) {
			continue
		} else {
			break
		}
	}

	value, err := generateRandomString(128)
	if err != nil {
		return err
	}


	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("INSERT INTO", userTableName, "(key,value,type,owner_id,description)")
	s.Append("VALUES", "(", s.Param(key), ",", s.Param(value), ",", s.Param(createParams.Type), ",", s.Param(createParams.OwnerID), ",", s.Param(createParams.Description), ")")

	logger.Debug(s.String())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *secretTable) CountSecrets(ctx context.Context, selector SecretSelector) (result uint64, err error) {
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("SELECT COUNT(*) AS count FORM", secretTableName)

	s.Append("WHERE")
	if selector.Key != nil {
		s.Append("key=", s.Param(selector.Key), "AND")
	}
	if selector.OwnerID != nil {
		s.Append("owner_id=", s.Param(selector.OwnerID), "AND")
	}

	s.Append("type=", s.Param(selector.Type), "AND")
	s.Append("deleted_time<=", s.Param(time.Now()))
	s.TrimSuffix("AND")

	logger.Debug(s.String())
	rowQuery := db.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *secretTable) QuerySecrets(ctx context.Context, selector SecretSelector, pagination *Pagination, sort *Sort) (result []*Secret, err error) {
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("SELECT")
	s.Append("key,")
	s.Append("type,")
	s.Append("value,")
	s.Append("owner_id,")
	s.Append("description,")
	s.Append("created_time,")
	s.Append("updated_time,")
	s.Append("deleted_time")
	s.Append("FORM", userTableName)

	s.Append("WHERE")
	if selector.Key != nil {
		s.Append("key=", s.Param(selector.Key), "AND")
	}
	if selector.OwnerID != nil {
		s.Append("owner_id=", s.Param(selector.OwnerID), "AND")
	}

	s.Append("type=", s.Param(selector.Type), "AND")
	s.Append("deleted_time<=", s.Param(time.Now()))
	s.TrimSuffix("AND")

	if pagination == nil {
		pagination = &Pagination{}
	}

	if pagination.Limit == nil {
		// 个人的 ak sk 默认上限 10
		defaultLimit := uint64(10)
		pagination.Limit = &defaultLimit
	}

	s.Append("LIMIT", s.Param(pagination.Limit))

	if pagination.Offset != nil {
		s.Append("OFFSET", s.Param(pagination.Offset))
	}

	if sort != nil {
		s.Append("ORDER BY", s.Param(sort.Key))
		if sort.Type == SortAsc {
			s.Append("ASC")
		}
		if sort.Type == SortDesc {
			s.Append("DESC")
		}
	}

	queryResult, err := db.Database.QueryContext(ctx, s.String(), s.Params()...)
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
			&secret.OwnerID,
			&secret.Description,
			&secret.CreatedTime,
			&secret.UpdatedTime,
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
