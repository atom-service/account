package model

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/package/proto"
	"github.com/atom-service/common/config"
	"github.com/atom-service/common/logger"
	"github.com/yinxulai/sqls"
)

func init() {
	config.Declare("init_admin_password", helper.GenerateRandomString(12, nil), true, "admin 用户的初始密码")
	config.Declare("init_admin_username", helper.GenerateRandomString(12, nil), true, "admin 用户的初始账号名")
}

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

func (srv *User) LoadProto(user *proto.User) (err error) {
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

// ToProto ToProto
func (srv *User) ToProto() *proto.User {
	user := new(proto.User)
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

func (srv *UserSelector) LoadProto(data *proto.UserSelector) {
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

var UserTable = &userTable{}

func (s *userTable) initData(ctx context.Context) (err error) {
	adminUserID := int64(1)
	userSelector := UserSelector{ID: &adminUserID}
	queryResult, err := s.QueryUsers(ctx, userSelector, nil, nil)
	if err != nil {
		return err
	}

	adminUsername := config.MustGet("init_admin_username")
	adminPassword := config.MustGet("init_admin_password")
	adminPasswordHash := Password.Hash(adminPassword)
	adminUser := &User{Username: &adminUsername, Password: &adminPasswordHash}

	if len(queryResult) > 0 {
		adminUsername = *queryResult[0].Username
		err = s.UpdateUser(ctx, UserSelector{ID: &adminUserID}, adminUser)
		if err != nil {
			return err
		}
	} else {
		err = s.CreateUser(ctx, *adminUser)
		if err != nil {
			return err
		}
	}

	logger.Info("admin user are upsert:")
	logger.Infof("username: %s", adminUsername)
	logger.Infof("password: %s", adminPassword)
	return nil
}

func (t *userTable) InitTable(ctx context.Context) error {
	tx, err := Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
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

	return t.initData(ctx)
}

func (r *userTable) CreateUser(ctx context.Context, newUser User) (err error) {
	s := sqls.INSERT_INTO(userTableName)
	s.VALUES("parent_id", s.Param(newUser.ParentID))
	s.VALUES("username", s.Param(newUser.Username))
	s.VALUES("password", s.Param(newUser.Password))

	logger.Debug(s.String(), s.Params())
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *userTable) DeleteUser(ctx context.Context, selector UserSelector) (err error) {
	s := sqls.UPDATE(userTableName)

	if selector.ID == nil && selector.Username == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
	}

	// 特殊的 admin 账号禁止删除
	if selector.ID != nil && *selector.ID == 1 {
		return fmt.Errorf("selector conditions cannot be admin user")
	}

	if selector.Username != nil && *selector.Username == "admin" {
		return fmt.Errorf("selector conditions cannot be admin user")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Username != nil {
		s.WHERE("username=" + s.Param(selector.Username))
	}

	s.SET("deleted_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *userTable) UpdateUser(ctx context.Context, selector UserSelector, user *User) (err error) {
	s := sqls.UPDATE(userTableName)

	if selector.ID == nil && selector.Username == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
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
	_, err = Database.ExecContext(ctx, s.String(), s.Params()...)
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
	rowQuery := Database.QueryRowContext(ctx, s.String(), s.Params()...)
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
	queryResult, err := Database.QueryContext(ctx, s.String(), s.Params()...)
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
