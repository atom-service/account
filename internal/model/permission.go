package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/common/logger"
	"github.com/yinxulai/sqls"
)

var permissionSchema = "\"permission\""
var roleTableName = permissionSchema + ".\"roles\""
var userRoleTableName = permissionSchema + ".\"user_roles\""
var resourceTableName = permissionSchema + ".\"resources\""
var roleResourceTableName = permissionSchema + ".\"role_resources\""
var roleResourceRuleTableName = permissionSchema + ".\"role_resource_rules\""

type Role struct {
	ID           *int64
	Name         *string
	Description  *string
	CreatedTime  *time.Time
	UpdatedTime  *time.Time
	DeletedTime  *time.Time
	DisabledTime *time.Time
}

type RoleSelector struct {
	ID   *int64
	Name *string
}

type roleTable struct{}

func (t *roleTable) CreateTable(ctx context.Context) error {
	tx, err := db.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	cs := sqls.CREATE_SCHEMA(permissionSchema).IF_NOT_EXISTS()
	if _, err = tx.ExecContext(ctx, cs.String()); err != nil {
		tx.Rollback()
		return err
	}
	// 创建 table
	s := sqls.CREATE_TABLE(roleTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial NOT NULL")
	s.COLUMN("name character varying(64) NOT NULL")
	s.COLUMN("description character varying(128) NOT NULL")
	s.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("disabled_time timestamp without time zone NULL")
	s.COLUMN("deleted_time timestamp without time zone NULL")
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

func (t *roleTable) TruncateTable(ctx context.Context) error {
	_, err := db.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(roleTableName).String())
	return err
}

func (r *roleTable) CreateRole(ctx context.Context, newRole Role) (err error) {
	s := sqls.INSERT_INTO(roleTableName)
	s.VALUES("name", s.Param(newRole.Name))
	s.VALUES("description", s.Param(newRole.Description))

	logger.Debug(s.String())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *roleTable) UpdateRole(ctx context.Context, selector RoleSelector, role *Role) (err error) {
	s := sqls.UPDATE(roleTableName)

	if selector.ID == nil && selector.Name == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	if role.Name != nil {
		s.SET("name", s.Param(*role.Name))
	}

	if role.Description != nil {
		s.SET("description", s.Param(*role.Description))
	}

	if role.DeletedTime != nil {
		s.SET("disabled_time", s.Param(*role.DeletedTime))
	}

	s.SET("updated_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleTable) DeleteRole(ctx context.Context, selector RoleSelector) (err error) {
	s := sqls.UPDATE(roleTableName)

	if selector.ID == nil && selector.Name == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	s.SET("deleted_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleTable) CountRoles(ctx context.Context, selector RoleSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(roleTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	s.WHERE("(deleted_time<CURRENT_TIMESTAMP OR deleted_time IS NULL)")

	logger.Debug(s.String())
	rowQuery := db.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleTable) QueryRoles(ctx context.Context, selector RoleSelector, pagination *Pagination, sort *Sort) (result []*Role, err error) {
	s := sqls.SELECT("id")
	s.SELECT("name")
	s.SELECT("description")
	s.SELECT("created_time")
	s.SELECT("updated_time")
	s.SELECT("disabled_time")
	s.SELECT("deleted_time")
	s.FROM(roleTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
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

	queryResult, err := db.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		role := Role{}
		if err = queryResult.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedTime,
			&role.UpdatedTime,
			&role.DisabledTime,
			&role.DeletedTime,
		); err != nil {
			logger.Error(err)
			return
		}
		result = append(result, &role)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}

type Resource struct {
	ID          *int64
	Name        *string
	Description *string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	DeletedTime *time.Time
}
type ResourceSelector struct {
	ID   *int64
	Name *string
}

type resourceTable struct{}

func (t *resourceTable) CreateTable(ctx context.Context) error {
	tx, err := db.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	cs := sqls.CREATE_SCHEMA(permissionSchema).IF_NOT_EXISTS()
	if _, err = tx.ExecContext(ctx, cs.String()); err != nil {
		tx.Rollback()
		return err
	}
	// 创建 table
	s := sqls.CREATE_TABLE(resourceTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial NOT NULL")
	s.COLUMN("name character varying(64) NOT NULL")
	s.COLUMN("description character varying(128) NOT NULL")
	s.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("deleted_time timestamp without time zone NULL")
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

func (t *resourceTable) TruncateTable(ctx context.Context) error {
	_, err := db.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(resourceTableName).String())
	return err
}

func (r *resourceTable) CreateResource(ctx context.Context, newResource Resource) (err error) {
	s := sqls.INSERT_INTO(resourceTableName)
	s.VALUES("name", s.Param(newResource.Name))
	s.VALUES("description", s.Param(newResource.Description))

	logger.Debug(s.String())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *resourceTable) UpdateResource(ctx context.Context, selector ResourceSelector, resource *Resource) (err error) {
	s := sqls.UPDATE(resourceTableName)

	if selector.ID == nil && selector.Name == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	if resource.Name != nil {
		s.SET("name", s.Param(*resource.Name))
	}

	if resource.Description != nil {
		s.SET("description", s.Param(*resource.Description))
	}

	if resource.DeletedTime != nil {
		s.SET("disabled_time", s.Param(*resource.DeletedTime))
	}

	s.SET("updated_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *resourceTable) DeleteResource(ctx context.Context, selector ResourceSelector) (err error) {
	s := sqls.UPDATE(resourceTableName)

	if selector.ID == nil && selector.Name == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	s.SET("deleted_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *resourceTable) CountResources(ctx context.Context, selector ResourceSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(resourceTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	s.WHERE("(deleted_time<CURRENT_TIMESTAMP OR deleted_time IS NULL)")

	logger.Debug(s.String())
	rowQuery := db.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *resourceTable) QueryResources(ctx context.Context, selector ResourceSelector, pagination *Pagination, sort *Sort) (result []*Resource, err error) {
	s := sqls.SELECT("id")
	s.SELECT("name")
	s.SELECT("description")
	s.SELECT("created_time")
	s.SELECT("updated_time")
	s.SELECT("deleted_time")
	s.FROM(resourceTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
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

	queryResult, err := db.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		resource := Resource{}
		if err = queryResult.Scan(
			&resource.ID,
			&resource.Name,
			&resource.Description,
			&resource.CreatedTime,
			&resource.UpdatedTime,
			&resource.DeletedTime,
		); err != nil {
			logger.Error(err)
			return
		}
		result = append(result, &resource)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}

const (
	RoleResourceInsertAction = "Insert"
	RoleResourceDeleteAction = "Delete"
	RoleResourceUpdateAction = "Update"
	RoleResourceQueryAction  = "Query"
)

type RoleResource struct {
	ID         *int64
	Action     string
	ResourceID int64
}

type RoleResourceSelector struct {
	ID         *int64
	Action     *string
	ResourceID *int64
}

type roleResourceTable struct{}

func (t *roleResourceTable) CreateTable(ctx context.Context) error {
	tx, err := db.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	cs := sqls.CREATE_SCHEMA(permissionSchema).IF_NOT_EXISTS()
	if _, err = tx.ExecContext(ctx, cs.String()); err != nil {
		tx.Rollback()
		return err
	}

	// 创建 table
	s := sqls.CREATE_TABLE(roleResourceTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial NOT NULL")
	s.COLUMN("action character varying(32) NOT NULL")
	s.COLUMN("resource_id int NOT NULL")
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

func (t *roleResourceTable) TruncateTable(ctx context.Context) error {
	_, err := db.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(roleResourceTableName).String())
	return err
}

func (r *roleResourceTable) CreateRoleResource(ctx context.Context, newResource RoleResource) (err error) {
	s := sqls.INSERT_INTO(roleResourceTableName)
	s.VALUES("action", s.Param(newResource.Action))
	s.VALUES("resource_id", s.Param(newResource.ResourceID))

	logger.Debug(s.String())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *roleResourceTable) DeleteRoleResource(ctx context.Context, selector RoleResourceSelector) (err error) {
	s := sqls.DELETE_FROM(roleResourceTableName)

	if selector.ID == nil && selector.Action == nil && selector.ResourceID == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Action != nil {
		s.WHERE("action=" + s.Param(selector.Action))
	}

	if selector.ResourceID != nil {
		s.WHERE("resource_id=" + s.Param(selector.ResourceID))
	}

	logger.Debug(s.String(), s.Params())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleResourceTable) CountRoleResources(ctx context.Context, selector RoleResourceSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(roleResourceTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Action != nil {
		s.WHERE("action=" + s.Param(selector.Action))
	}
	if selector.ResourceID != nil {
		s.WHERE("resource_id=" + s.Param(selector.ResourceID))
	}

	logger.Debug(s.String())
	rowQuery := db.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleResourceTable) QueryRoleResources(ctx context.Context, selector RoleResourceSelector, pagination *Pagination, sort *Sort) (result []*RoleResource, err error) {
	s := sqls.SELECT("id")
	s.SELECT("action")
	s.SELECT("resource_id")
	s.FROM(roleResourceTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Action != nil {
		s.WHERE("action=" + s.Param(selector.Action))
	}
	if selector.ResourceID != nil {
		s.WHERE("resource_id=" + s.Param(selector.ResourceID))
	}

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

	queryResult, err := db.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		roleResource := RoleResource{}
		if err = queryResult.Scan(
			&roleResource.ID,
			&roleResource.Action,
			&roleResource.ResourceID,
		); err != nil {
			logger.Error(err)
			return
		}
		result = append(result, &roleResource)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}

type RoleResourceRule struct {
	ID             *int64
	Key            string
	Value          string
	RoleResourceID int64
}

type RoleResourceRuleSelector struct {
	ID             *int64
	Key            *string
	RoleResourceID *int64
}

type roleResourceRuleTable struct {
}

func (t *roleResourceRuleTable) CreateTable(ctx context.Context) error {
	tx, err := db.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	cs := sqls.CREATE_SCHEMA(permissionSchema).IF_NOT_EXISTS()
	if _, err = tx.ExecContext(ctx, cs.String()); err != nil {
		tx.Rollback()
		return err
	}

	// 创建 table
	s := sqls.CREATE_TABLE(roleResourceRuleTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial NOT NULL")
	s.COLUMN("key character varying(64) NOT NULL")
	s.COLUMN("value character varying(128) NOT NULL")
	s.COLUMN("role_resource_id int NOT NULL")
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

func (t *roleResourceRuleTable) TruncateTable(ctx context.Context) error {
	_, err := db.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(roleResourceRuleTableName).String())
	return err
}

func (r *roleResourceRuleTable) CreateRoleResourceRule(ctx context.Context, newRule RoleResourceRule) (err error) {
	s := sqls.INSERT_INTO(roleResourceRuleTableName)
	s.VALUES("key", s.Param(newRule.Key))
	s.VALUES("value", s.Param(newRule.Value))
	s.VALUES("role_resource_id", s.Param(newRule.RoleResourceID))

	logger.Debug(s.String())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *roleResourceRuleTable) DeleteRoleResourceRule(ctx context.Context, selector RoleResourceRuleSelector) (err error) {
	s := sqls.DELETE_FROM(roleResourceRuleTableName)

	if selector.ID == nil && selector.RoleResourceID == nil && selector.Key == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}

	if selector.RoleResourceID != nil {
		s.WHERE("role_resource_id=" + s.Param(selector.RoleResourceID))
	}

	logger.Debug(s.String(), s.Params())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleResourceRuleTable) CountRoleResourceRules(ctx context.Context, selector RoleResourceRuleSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(roleResourceRuleTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}

	if selector.RoleResourceID != nil {
		s.WHERE("role_resource_id=" + s.Param(selector.RoleResourceID))
	}

	logger.Debug(s.String())
	rowQuery := db.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleResourceRuleTable) QueryRoleResourceRules(ctx context.Context, selector RoleResourceRuleSelector, pagination *Pagination, sort *Sort) (result []*RoleResourceRule, err error) {
	s := sqls.SELECT("id")
	s.SELECT("key")
	s.SELECT("value")
	s.SELECT("role_resource_id")
	s.FROM(roleResourceRuleTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Key != nil {
		s.WHERE("key=" + s.Param(selector.Key))
	}

	if selector.RoleResourceID != nil {
		s.WHERE("role_resource_id=" + s.Param(selector.RoleResourceID))
	}

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

	queryResult, err := db.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		roleResourceRule := RoleResourceRule{}
		if err = queryResult.Scan(
			&roleResourceRule.ID,
			&roleResourceRule.Key,
			&roleResourceRule.Value,
			&roleResourceRule.RoleResourceID,
		); err != nil {
			logger.Error(err)
			return
		}
		result = append(result, &roleResourceRule)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}

type UserRole struct {
	ID           *int64
	UserID       int64
	RoleID       int64
	CreatedTime  *time.Time
	UpdatedTime  *time.Time
	DisabledTime *time.Time
}

type UserRoleSelector struct {
	ID     *int64
	UserID *int64
	RoleID *int64
}

type userRoleTable struct {
}

func (t *userRoleTable) CreateTable(ctx context.Context) error {
	tx, err := db.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return err
	}

	// 创建 schema
	cs := sqls.CREATE_SCHEMA(permissionSchema).IF_NOT_EXISTS()
	if _, err = tx.ExecContext(ctx, cs.String()); err != nil {
		tx.Rollback()
		return err
	}

	// 创建 table
	s := sqls.CREATE_TABLE(userRoleTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial NOT NULL")
	s.COLUMN("user_id int NOT NULL")
	s.COLUMN("role_id int NOT NULL")
	s.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("disabled_time timestamp without time zone NULL")
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

func (t *userRoleTable) TruncateTable(ctx context.Context) error {
	_, err := db.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(userRoleTableName).String())
	return err
}

func (r *userRoleTable) CreateUserRole(ctx context.Context, newRule UserRole) (err error) {
	s := sqls.INSERT_INTO(userRoleTableName)
	s.VALUES("user_id", s.Param(newRule.UserID))
	s.VALUES("role_id", s.Param(newRule.RoleID))

	logger.Debug(s.String())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *userRoleTable) DeleteUserRole(ctx context.Context, selector UserRoleSelector) (err error) {
	s := sqls.DELETE_FROM(userRoleTableName)

	if selector.ID == nil && selector.UserID == nil && selector.RoleID == nil {
		return fmt.Errorf("elector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	if selector.RoleID != nil {
		s.WHERE("role_id=" + s.Param(selector.RoleID))
	}

	logger.Debug(s.String(), s.Params())
	_, err = db.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *userRoleTable) CountUserRole(ctx context.Context, selector UserRoleSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(userRoleTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	if selector.RoleID != nil {
		s.WHERE("role_id=" + s.Param(selector.RoleID))
	}

	logger.Debug(s.String())
	rowQuery := db.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *userRoleTable) QueryUserRole(ctx context.Context, selector UserRoleSelector, pagination *Pagination, sort *Sort) (result []*UserRole, err error) {
	s := sqls.SELECT(
		"id",
		"user_id",
		"role_id",
		"created_time",
		"updated_time",
		"disabled_time",
	).FROM(userRoleTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.UserID != nil {
		s.WHERE("user_id=" + s.Param(selector.UserID))
	}

	if selector.RoleID != nil {
		s.WHERE("role_id=" + s.Param(selector.RoleID))
	}

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

	queryResult, err := db.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		userRole := UserRole{}
		if err = queryResult.Scan(
			&userRole.ID,
			&userRole.UserID,
			&userRole.RoleID,
			&userRole.CreatedTime,
			&userRole.UpdatedTime,
			&userRole.DisabledTime,
		); err != nil {
			logger.Error(err)
			return
		}
		result = append(result, &userRole)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
		return
	}
	return
}
