package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/logger"
	"github.com/atom-service/common/sqls"
)

var tableName = "\"user\".\"users\""

type userTable struct{}

type User struct {
	ID          *int64     `json:"id"`
	ParentID    *int64     `json:"parent_id"`
	Username    *string    `json:"username"`
	Password    *string    `json:"password"`
	CreatedTime *time.Time `json:"created-time"`
	UpdatedTime *time.Time `json:"updated-time"`
	DeletedTime *time.Time `json:"deleted-time"`
}

func (srv *User) LoadProtoStruct(user *protos.User) (err error) {
	srv.ID = &user.ID
	srv.Username = &user.Username
	srv.Password = &user.Password

	createdTime, err := time.Parse(time.RFC3339Nano, user.CreatedTime)
	if err != nil {
		return err
	}

	srv.CreatedTime = &createdTime

	updatedTime, err := time.Parse(time.RFC3339Nano, user.UpdatedTime)
	if err != nil {
		return err
	}

	srv.UpdatedTime = &updatedTime

	if user.DeletedTime != nil {
		deletedTime, err := time.Parse(time.RFC3339Nano, *user.DeletedTime)
		if err != nil {
			return err
		}

		srv.DeletedTime = &deletedTime
	}

	return
}

// OutProtoStruct OutProtoStruct
func (srv *User) OutProtoStruct() *protos.User {
	user := new(protos.User)
	user.ID = *srv.ID
	// user.Password = srv.Password.String
	user.Username = *srv.Username
	user.CreatedTime = srv.CreatedTime.String()
	user.UpdatedTime = srv.UpdatedTime.String()

	if srv.DeletedTime != nil {
		timeString := srv.DeletedTime.String()
		user.DeletedTime = &timeString
	}

	return user
}

type UserSelector struct {
	ID       *int64
	Username *string
}

func (srv *UserSelector) LoadProtoStruct(data *protos.UserSelector) {
	srv.ID = data.ID
	srv.Username = data.Username
}

// OutProtoStruct OutProtoStruct
func (srv *UserSelector) OutProtoStruct() *protos.UserSelector {
	result := new(protos.UserSelector)
	result.Username = srv.Username
	result.ID = srv.ID
	return result
}

var UserTable = &userTable{}

func (t *userTable) CreateTable(ctx context.Context) error {
	tx, err := db.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	tx.ExecContext(ctx, "CREATE SCHEMA IF NOT EXISTS \"user\"")

	// 创建 table
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("CREATE TABLE IF NOT EXISTS \"user\".\"users\" (")
	s.Append("id serial NOT NULL,")
	s.Append("parent_id integer NULL,")
	s.Append("username character varying(64) NOT NULL,")
	s.Append("password character varying(256) NOT NULL,")
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

func (t *userTable) TruncateTable(ctx context.Context) error {
	_, err := db.Database.ExecContext(ctx, "TRUNCATE TABLE \"user\".\"users\";")
	return err
}

func (r *userTable) CreateUser(ctx context.Context, newUser User) (err error) {
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("INSERT INTO", tableName, "(parent_id,username,password)")
	s.Append("VALUES", "(", s.Param(newUser.ParentID), ",", s.Param(newUser.Username), ",", s.Param(newUser.Password), ")")

	logger.Debug(s.String())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *userTable) CountUsers(ctx context.Context, selector UserSelector) (result uint64, err error) {
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("SELECT COUNT(*) AS count FORM", tableName)

	if selector.ID != nil || selector.Username != nil {
		s.Append("WHERE")
		if selector.ID != nil {
			s.Append("id=", s.Param(selector.ID), ",")
		}
		if selector.Username != nil {
			s.Append("username=", s.Param(selector.Username), ",")
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

func (r *userTable) QueryUsers(ctx context.Context, selector UserSelector, pagination *Pagination, sort *Sort) (result []*User, err error) {
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("SELECT")
	s.Append("id,")
	s.Append("parent_id,")
	s.Append("username,")
	s.Append("password,")
	s.Append("created_time,")
	s.Append("updated_time,")
	s.Append("deleted_time")
	s.Append("FORM", tableName)

	if selector.ID != nil || selector.Username != nil {
		s.Append("WHERE")
		if selector.ID != nil {
			s.Append("id =", s.Param(selector.ID), ",")
		}

		if selector.Username != nil {
			s.Append("username=", s.Param(selector.Username), ",")
		}

		s.TrimSuffix(",")
	}

	if pagination != nil {
		if pagination.Limit == nil {
			// 限制最大
			defaultLimit := uint64(1000)
			pagination.Limit = &defaultLimit
		}

		s.Append("LIMIT", s.Param(pagination.Limit))

		if pagination.Offset != nil {
			s.Append("OFFSET", s.Param(pagination.Offset))
		}
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
		user := User{}
		if err = queryResult.Scan(
			&user.ID,
			&user.ParentID,
			&user.Username,
			&user.Password,
			&user.CreatedTime,
			&user.UpdatedTime,
			&user.DeletedTime,
		); err != nil {
			logger.Error(err)
			return
		}
		result = append(result, &user)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}
