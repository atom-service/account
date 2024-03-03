package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/logger"
	"github.com/atom-service/common/sqls"
)

var tableName = "\"user\".\"users\""

type userTable struct{}

type User struct {
	ID          sql.NullInt64  `json:"id"`
	ParentID    sql.NullInt64  `json:"parent_id"`
	Username    sql.NullString `json:"username"`
	Password    sql.NullString `json:"password"`
	CreatedTime sql.NullTime   `json:"created-time"`
	UpdatedTime sql.NullTime   `json:"updated-time"`
	DeletedTime sql.NullTime   `json:"deleted-time"`
}

func (srv *User) LoadProtoStruct(user *protos.User) {
	srv.ID.Scan(user.ID)
	srv.Username.Scan(user.Username)
	srv.Password.Scan(user.Password)
	srv.CreatedTime.Scan(user.CreatedTime)
	srv.UpdatedTime.Scan(user.UpdatedTime)

	if user.DeletedTime != nil {
		srv.DeletedTime.Scan(user.DeletedTime)
	}
}

// OutProtoStruct OutProtoStruct
func (srv *User) OutProtoStruct() *protos.User {
	user := new(protos.User)
	user.ID = srv.ID.Int64
	// user.Password = srv.Password.String
	user.Username = srv.Username.String
	user.CreatedTime = srv.CreatedTime.Time.String()
	user.UpdatedTime = srv.UpdatedTime.Time.String()

	if srv.DeletedTime.Valid {
		timeString := srv.DeletedTime.Time.String()
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
	tx.ExecContext(ctx, s.String())
	logger.Debug(s.String())
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

func (r *userTable) CreateUser(ctx context.Context, newUser User) (user User, err error) {
	// 检查昵称是否存在

	return
}

func (r *userTable) CountUsers(ctx context.Context, selector UserSelector) (result uint64, err error) {
	s := sqls.Begin(sqls.PostgreSQL)
	s.Append("SELECT COUNT(*) AS count FORM", tableName)
	if selector.ID != nil || selector.Username != nil {
		s.Append("WHERE")
		if selector.ID != nil {
			s.Append("id = ", s.Param(selector.ID), ",")
		}
		if selector.Username != nil {
			s.Append("username = ", s.Param(selector.Username), ",")
		}

		s.TrimSuffix(",")
	}

	rowQuery := db.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *userTable) QueryUsers(ctx context.Context, selector UserSelector, pagination Pagination, sort *Sort) (result PaginationResult[User], err error) {
	result.Data = []User{}

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
			s.Append("id = ", s.Param(selector.ID), ",")
		}
		if selector.Username != nil {
			s.Append("username = ", s.Param(selector.Username), ",")
		}

		s.TrimSuffix(",")
	}

	if pagination.Limit == nil {
		// 限制最大
		defaultLimit := uint64(1000)
		pagination.Limit = &defaultLimit
	}

	s.Append("LIMIT ", s.Param(pagination.Limit))

	if pagination.Offset != nil {
		s.Append("OFFSET ", s.Param(pagination.Offset))
	}

	if sort != nil {
		s.Append("ORDER BY ", s.Param(sort.Key), " ")
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
		result.Data = append(result.Data, user)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}
