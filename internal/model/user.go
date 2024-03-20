package model

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/atom-service/account/internal/database"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/logger"
	"github.com/yinxulai/sqls"
)

var userSchemaName = "\"user\""
var userTableName = userSchemaName + ".\"users\""

type userTable struct{}

var Password = &password{}

type password struct{}

func (p *password) Hash(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

type User struct {
	ID           *int64
	ParentID     *int64
	Username     *string
	Password     *string
	CreatedTime  *time.Time
	UpdatedTime  *time.Time
	DeletedTime  *time.Time
	DisabledTime *time.Time
}

func (srv *User) LoadProtoStruct(user *protos.User) (err error) {
	if user == nil {
		return fmt.Errorf("user is nil")
	}

	srv.ID = &user.ID
	srv.Username = &user.Username
	srv.Password = &user.Password

	createdTime, err := time.Parse(time.RFC3339Nano, user.CreatedTime)
	if err != nil {
		return fmt.Errorf("failed to parse created time: %v", err)
	}
	srv.CreatedTime = &createdTime

	updatedTime, err := time.Parse(time.RFC3339Nano, user.UpdatedTime)
	if err != nil {
		return fmt.Errorf("failed to parse updated time: %v", err)
	}
	srv.UpdatedTime = &updatedTime

	if user.DeletedTime != nil {
		deletedTime, err := time.Parse(time.RFC3339Nano, *user.DeletedTime)
		if err != nil {
			return fmt.Errorf("failed to parse deleted time: %v", err)
		}
		srv.DeletedTime = &deletedTime
	}

	if user.DisabledTime != nil {
		disabledTime, err := time.Parse(time.RFC3339Nano, *user.DisabledTime)
		if err != nil {
			return fmt.Errorf("failed to parse disabled time: %v", err)
		}
		srv.DisabledTime = &disabledTime
	}

	return nil
}

// OutProtoStruct OutProtoStruct
func (srv *User) OutProtoStruct() *protos.User {
	user := new(protos.User)
	user.ID = *srv.ID

	if srv.Username != nil {
		user.Username = *srv.Username
	}
	
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
	if data == nil {
		return
	}

	if data.ID != nil {
		srv.ID = data.ID
	}

	if data.Username != nil {
		srv.Username = data.Username
	}
}

// OutProtoStruct OutProtoStruct
func (srv *UserSelector) OutProtoStruct() *protos.UserSelector {
	result := &protos.UserSelector{
		Username: srv.Username,
		ID:       srv.ID,
	}
	
	return result
}

var UserTable = &userTable{}

func (t *userTable) CreateTable(ctx context.Context) error {
	tx, err := database.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// 创建 schema
	cs := sqls.CREATE_SCHEMA(userSchemaName).IF_NOT_EXISTS()
	logger.Debug(cs.String())
	_, err = tx.ExecContext(ctx, cs.String())
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	// 创建 table
	ct := sqls.CREATE_TABLE(userTableName).IF_NOT_EXISTS()
	ct.COLUMN("id serial PRIMARY KEY NOT NULL")
	ct.COLUMN("parent_id integer NULL")
	ct.COLUMN("username character varying(64) UNIQUE NOT NULL")
	ct.COLUMN("password character varying(256) NOT NULL")
	ct.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	ct.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	ct.COLUMN("disabled_time timestamp without time zone NULL")
	ct.COLUMN("deleted_time timestamp without time zone NULL")

	logger.Debug(ct.String())
	_, err = tx.ExecContext(ctx, ct.String())
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	if err := tx.Commit(); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Join(err)
		}
		return err
	}

	return nil
}

func (t *userTable) TruncateTable(ctx context.Context) error {
	_, err := database.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(userTableName).String())
	return err
}

func (r *userTable) CreateUser(ctx context.Context, newUser User) (err error) {
	s := sqls.INSERT_INTO(userTableName)
	s.VALUES("parent_id", s.Param(newUser.ParentID))
	s.VALUES("username", s.Param(newUser.Username))
	s.VALUES("password", s.Param(newUser.Password))

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *userTable) DeleteUser(ctx context.Context, selector UserSelector) (err error) {
	s := sqls.UPDATE(userTableName)

	if selector.ID == nil && selector.Username == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Username != nil {
		s.WHERE("username=" + s.Param(selector.Username))
	}

	s.SET("deleted_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *userTable) UpdateUser(ctx context.Context, selector UserSelector, user *User) (err error) {
	s := sqls.UPDATE(userTableName)

	if selector.ID == nil && selector.Username == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Username != nil {
		s.WHERE("username=" + s.Param(selector.Username))
	}

	if user.Password != nil {
		s.SET("password", s.Param(*user.Password))
	}

	if user.DisabledTime != nil {
		s.SET("disabled_time", s.Param(*user.DisabledTime))
	}

	s.SET("updated_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *userTable) CountUsers(ctx context.Context, selector UserSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(userTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Username != nil {
		s.WHERE("username=" + s.Param(selector.Username))
	}

	s.WHERE("(deleted_time<CURRENT_TIMESTAMP OR deleted_time IS NULL)")

	logger.Debug(s.String(), s.Params())
	rowQuery := database.Database.QueryRowContext(ctx, s.String(), s.Params()...)
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

	logger.Debug(s.String(), s.Params())
	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
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
