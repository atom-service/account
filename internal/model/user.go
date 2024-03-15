package model

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/logger"
	"github.com/yinxulai/sqls"
)

var userTableName = "\"user\".\"users\""

type userTable struct{}

var Password = &password{}

type password struct{}

func (p *password) Hash(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

type User struct {
	ID           *int64     `json:"id"`
	ParentID     *int64     `json:"parent-id"`
	Username     *string    `json:"username"`
	Password     *string    `json:"password"`
	CreatedTime  *time.Time `json:"created_time"`
	UpdatedTime  *time.Time `json:"updated_time"`
	DeletedTime  *time.Time `json:"deleted_time"`
	DisabledTime *time.Time `json:"disabled_time"`
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

	if user.DisabledTime != nil {
		disabledTime, err := time.Parse(time.RFC3339Nano, *user.DisabledTime)
		if err != nil {
			return err
		}

		srv.DisabledTime = &disabledTime
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

	if srv.DisabledTime != nil {
		timeString := srv.DisabledTime.String()
		user.DisabledTime = &timeString
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
	tx.ExecContext(ctx, sqls.CREATE_SCHEMA("user").IF_NOT_EXISTS().String())

	// 创建 table
	s := sqls.CREATE_TABLE(userTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial NOT NULL")
	s.COLUMN("parent_id integer NULL")
	s.COLUMN("username character varying(64) NOT NULL")
	s.COLUMN("password character varying(256) NOT NULL")
	s.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("disabled_time timestamp without time zone NULL")
	s.COLUMN("deleted_time timestamp without time zone NULL")

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
	_, err := db.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(userTableName).String())
	return err
}

func (r *userTable) CreateUser(ctx context.Context, newUser User) (err error) {
	s := sqls.INSERT_INTO(userTableName)
	s.VALUES("parent_id", s.Param(newUser.ParentID))
	s.VALUES("username", s.Param(newUser.Username))
	s.VALUES("password", s.Param(newUser.Password))

	logger.Debug(s.String())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *userTable) CountUsers(ctx context.Context, selector UserSelector) (result uint64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(userTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Username != nil {
		s.WHERE("username=" + s.Param(selector.Username))
	}

	s.WHERE("deleted_time=" + s.Param(time.Now()))

	logger.Debug(s.String())
	rowQuery := db.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *userTable) QueryUsers(ctx context.Context, selector UserSelector, pagination *Pagination, sort *Sort) (result []*User, err error) {
	s := sqls.SELECT(
		"id",
		"parent_id",
		"username",
		"password",
		"created_time",
		"updated_time",
		"deleted_time",
		"disabled_time",
	).FROM(userTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Username != nil {
		s.WHERE("username=" + s.Param(selector.Username))
	}

	s.WHERE("deleted_time=" + s.Param(time.Now()))

	if pagination == nil {
		pagination = &Pagination{}
	}

	if pagination.Limit == nil {
		// 默认为 100，防止刷爆
		defaultLimit := uint64(100)
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
			&user.DisabledTime,
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
