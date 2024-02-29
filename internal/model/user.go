package model

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/logger"
	"github.com/atom-service/common/sqls"
)

var tableName = "\"user\".\"users\""

type userTable struct{}

type User struct {
	ID          sql.NullInt64  `json:"id"`
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
	tx.ExecContext(ctx, strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS \"user\".\"users\" (",
		"id serial NOT NULL,",
		"parent_id integer NULL,",
		"username character varying(64) NOT NULL,",
		"password character varying(256) NOT NULL,",
		"created_time timestamp without time zone NULL DEFAULT now(),",
		"updated_time timestamp without time zone NULL DEFAULT now(),",
		"deleted_time timestamp without time zone NULL",
		");",
	}, " "))

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

func (r *userTable) CreateUser() {}

func (r *userTable) QueryUsers(ctx context.Context, selector UserSelector, pagination Pagination, sort Sort) (PaginationResult[User], error) {
	result := PaginationResult[User]{
		Total: 0,
		Data:  []User{},
	}

	whereCond := map[string]any{}
	if selector.ID != nil {
		whereCond["id="] = selector.ID
	}
	if selector.Username != nil {
		whereCond["username="] = selector.Username
	}

	selectSql := sqls.Select(tableName, "COUNT(*)").Where(whereCond)
	logger.Debugf("ready to execute sql %s, %v", selectSql.String(), selectSql.Params())
	_, err := db.Database.QueryContext(ctx, selectSql.String(), selectSql.Params()...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return result, nil
}
