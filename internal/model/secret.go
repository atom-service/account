package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/common/logger"
	"github.com/atom-service/common/sqls"
)

var secretTableName = "\"secret\".\"secrets\""

type Secret struct {
	Key         *string    `json:"key"`
	Value       *string    `json:"value"`
	OwnerID     *string    `json:"owner_id"`
	Description *string    `json:"description"`
	CreatedTime *time.Time `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
	DeletedTime *time.Time `json:"deleted_time"`
}

type SecretSelector struct {
	Key     *int64
	OwnerID *int64
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

func (r *secretTable) CreateSecret(ctx context.Context, secret Secret) (err error) {
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("INSERT INTO", userTableName, "(key,value,owner_id,description)")
	s.Append("VALUES", "(", s.Param(secret.Key), ",", s.Param(secret.Value), ",", s.Param(secret.OwnerID), ",", s.Param(secret.Description), ")")

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

	if selector.Key != nil || selector.OwnerID != nil {
		s.Append("WHERE")
		if selector.Key != nil {
			s.Append("key=", s.Param(selector.Key), ",")
		}
		if selector.OwnerID != nil {
			s.Append("owner_id=", s.Param(selector.OwnerID), ",")
		}

		s.TrimSuffix(",")
	}

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
	s.Append("value,")
	s.Append("owner_id,")
	s.Append("description,")
	s.Append("created_time,")
	s.Append("updated_time,")
	s.Append("deleted_time")
	s.Append("FORM", userTableName)

	if selector.Key != nil || selector.OwnerID != nil {
		s.Append("WHERE")
		if selector.Key != nil {
			s.Append("key=", s.Param(selector.Key), ",")
		}
		if selector.OwnerID != nil {
			s.Append("owner_id=", s.Param(selector.OwnerID), ",")
		}

		s.TrimSuffix(",")
	}

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
