package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/atom-service/account/internal/database"
	"github.com/atom-service/account/package/proto"
	"github.com/atom-service/common/logger"
	"github.com/yinxulai/sqls"
)

var permissionSchema = "\"permission\""
var roleTableName = permissionSchema + ".\"roles\""
var userRoleTableName = permissionSchema + ".\"user_roles\""
var resourceTableName = permissionSchema + ".\"resources\""
var resourceRuleTableName = permissionSchema + ".\"resource_rules\""
var roleResourceTableName = permissionSchema + ".\"role_resources\""

var RoleTable = &roleTable{}
var UserRoleTable = &userRoleTable{}
var ResourceTable = &resourceTable{}
var RoleResourceTable = &roleResourceTable{}
var RoleResourceRuleTable = &resourceRuleTable{}

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

func (r *RoleSelector) LoadProto(data *proto.RoleSelector) {
	if data == nil {
		return
	}

	if data.ID != nil {
		r.ID = data.ID
	}

	if data.Name != nil {
		r.Name = data.Name
	}
}

func (srv *Role) ToProto() *proto.Role {
	role := new(proto.Role)
	if srv.ID != nil {
		role.ID = *srv.ID
	}
	if srv.Name != nil {
		role.Name = *srv.Name
	}
	if srv.Description != nil {
		role.Description = *srv.Description
	}

	if srv.Description != nil {
		role.Description = *srv.Description
	}

	if srv.CreatedTime != nil {
		role.CreatedTime = srv.CreatedTime.String()
	}

	if srv.UpdatedTime != nil {
		role.UpdatedTime = srv.UpdatedTime.String()
	}

	if srv.DeletedTime != nil {
		timeString := srv.DeletedTime.String()
		role.DeletedTime = &timeString
	}

	return role
}

type roleTable struct{}

func (t *roleTable) CreateTable(ctx context.Context) error {
	tx, err := database.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
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
	s.COLUMN("id serial PRIMARY KEY NOT NULL")
	s.COLUMN("name character varying(64) UNIQUE NOT NULL")
	s.COLUMN("description character varying(128) NULL")
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
	_, err := database.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(roleTableName).String())
	return err
}

func (r *roleTable) CreateRole(ctx context.Context, newRole Role) (err error) {
	s := sqls.INSERT_INTO(roleTableName)
	s.VALUES("name", s.Param(newRole.Name))
	s.VALUES("description", s.Param(newRole.Description))

	logger.Debug(s.String())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *roleTable) UpdateRole(ctx context.Context, selector RoleSelector, role *Role) (err error) {
	s := sqls.UPDATE(roleTableName)

	if selector.ID == nil && selector.Name == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
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
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleTable) DeleteRole(ctx context.Context, selector RoleSelector) (err error) {
	s := sqls.UPDATE(roleTableName)

	if selector.ID == nil && selector.Name == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	s.SET("deleted_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
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
	rowQuery := database.Database.QueryRowContext(ctx, s.String(), s.Params()...)
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

	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
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

func (srv *Resource) ToProto() *proto.Resource {
	resource := new(proto.Resource)
	resource.ID = *srv.ID

	if srv.Name != nil {
		resource.Name = *srv.Name
	}

	resource.CreatedTime = srv.CreatedTime.String()
	resource.UpdatedTime = srv.UpdatedTime.String()

	if srv.DeletedTime != nil {
		timeString := srv.DeletedTime.String()
		resource.DeletedTime = &timeString
	}

	return resource
}

type ResourceSelector struct {
	ID   *int64
	Name *string
}

func (r *ResourceSelector) LoadProto(data *proto.ResourceSelector) {
	if data == nil {
		return
	}

	r.ID = data.ID
	r.Name = data.Name
}

type resourceTable struct{}

func (t *resourceTable) CreateTable(ctx context.Context) error {
	tx, err := database.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
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
	s.COLUMN("id serial PRIMARY KEY NOT NULL")
	s.COLUMN("name character varying(64) UNIQUE NOT NULL")
	s.COLUMN("description character varying(128) NULL")
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
	_, err := database.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(resourceTableName).String())
	return err
}

func (r *resourceTable) CreateResource(ctx context.Context, newResource Resource) (err error) {
	s := sqls.INSERT_INTO(resourceTableName)
	s.VALUES("name", s.Param(newResource.Name))
	s.VALUES("description", s.Param(newResource.Description))

	logger.Debug(s.String())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *resourceTable) UpdateResource(ctx context.Context, selector ResourceSelector, resource *Resource) (err error) {
	s := sqls.UPDATE(resourceTableName)

	if selector.ID == nil && selector.Name == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
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
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *resourceTable) DeleteResource(ctx context.Context, selector ResourceSelector) (err error) {
	s := sqls.UPDATE(resourceTableName)

	if selector.ID == nil && selector.Name == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
	}

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	s.SET("deleted_time", s.Param(time.Now()))

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
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
	rowQuery := database.Database.QueryRowContext(ctx, s.String(), s.Params()...)
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

	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
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
	ActionInsert = "Insert"
	ActionDelete = "Delete"
	ActionUpdate = "Update"
	ActionQuery  = "Query"
)

type RoleResource struct {
	ID         *int64
	Action     string
	RoleID     int64
	ResourceID int64
}

func (srv *RoleResource) LoadProto(data *proto.RoleResource) {
	if data == nil {
		return
	}

	srv.RoleID = data.ResourceID

	if data.Action == proto.ResourceAction_Insert {
		srv.Action = ActionInsert
	}

	if data.Action == proto.ResourceAction_Delete {
		srv.Action = ActionDelete
	}

	if data.Action == proto.ResourceAction_Update {
		srv.Action = ActionUpdate
	}

	if data.Action == proto.ResourceAction_Query {
		srv.Action = ActionQuery
	}
}

type RoleResourceSelector struct {
	ID         *int64
	Action     *string
	RoleID     *int64
	ResourceID *int64
}

func (srv *RoleResourceSelector) LoadProtoAction(action proto.ResourceAction) {
	if action == proto.ResourceAction_Insert {
		var temp = ActionInsert
		srv.Action = &temp
	}

	if action == proto.ResourceAction_Delete {
		var temp = ActionDelete
		srv.Action = &temp
	}

	if action == proto.ResourceAction_Update {
		var temp = ActionUpdate
		srv.Action = &temp
	}

	if action == proto.ResourceAction_Query {
		var temp = ActionQuery
		srv.Action = &temp
	}
}

type roleResourceTable struct{}

func (t *roleResourceTable) CreateTable(ctx context.Context) error {
	tx, err := database.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
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
	s.COLUMN("id serial PRIMARY KEY NOT NULL")
	s.COLUMN("action character varying(32) NOT NULL")
	s.COLUMN("resource_id int NOT NULL")
	s.COLUMN("role_id int NOT NULL")
	s.OPTIONS("CONSTRAINT role_resource_union_unique_keys UNIQUE (action, resource_id, role_id)")
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
	_, err := database.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(roleResourceTableName).String())
	return err
}

func (r *roleResourceTable) CreateRoleResource(ctx context.Context, newResource RoleResource) (err error) {
	s := sqls.INSERT_INTO(roleResourceTableName)
	s.VALUES("action", s.Param(newResource.Action))
	s.VALUES("role_id", s.Param(newResource.RoleID))
	s.VALUES("resource_id", s.Param(newResource.ResourceID))

	logger.Debug(s.String())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleResourceTable) DeleteRoleResource(ctx context.Context, selector RoleResourceSelector) (err error) {
	s := sqls.DELETE_FROM(roleResourceTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}

	if selector.Action != nil {
		s.WHERE("action=" + s.Param(selector.Action))
	}

	if selector.RoleID != nil {
		s.WHERE("role_id=" + s.Param(selector.RoleID))
	}

	if selector.ResourceID != nil {
		s.WHERE("resource_id=" + s.Param(selector.ResourceID))
	}

	logger.Debug(s.String(), s.Params())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
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
	if selector.RoleID != nil {
		s.WHERE("role_id=" + s.Param(selector.RoleID))
	}
	if selector.ResourceID != nil {
		s.WHERE("resource_id=" + s.Param(selector.ResourceID))
	}

	logger.Debug(s.String())
	rowQuery := database.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *roleResourceTable) QueryRoleResources(ctx context.Context, selector RoleResourceSelector, pagination *Pagination, sort *Sort) (result []*RoleResource, err error) {
	s := sqls.SELECT("id")
	s.SELECT("action")
	s.SELECT("role_id")
	s.SELECT("resource_id")
	s.FROM(roleResourceTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Action != nil {
		s.WHERE("action=" + s.Param(selector.Action))
	}
	if selector.RoleID != nil {
		s.WHERE("role_id=" + s.Param(selector.RoleID))
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

	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
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
			&roleResource.RoleID,
			&roleResource.ResourceID,
		); err != nil {
			logger.Error(err)
			return
		}
		result = append(result, &roleResource)
	}
	if err = queryResult.Err(); err != nil {
		logger.Error(err)
	}
	return
}

type ResourceRule struct {
	ID             *int64
	Key            string
	Value          string
	RoleResourceID int64
}

const (
	// 代表匹配任何 key
	ResourceRuleKeyOfAny = "*"
	// 代表匹配任何 value
	ResourceRuleValueOfAny = "*"
)

type ResourceRuleSelector struct {
	ID             *int64
	Key            *string
	RoleResourceID *int64
}

type resourceRuleTable struct {
}

func (t *resourceRuleTable) CreateTable(ctx context.Context) error {
	tx, err := database.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
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
	s := sqls.CREATE_TABLE(resourceRuleTableName).IF_NOT_EXISTS()
	s.COLUMN("id serial PRIMARY KEY NOT NULL")
	s.COLUMN("role_resource_id int NOT NULL")
	s.COLUMN("key character varying(64) NOT NULL")
	s.COLUMN("value character varying(128) NOT NULL")
	s.OPTIONS("CONSTRAINT resource_rule_union_unique_keys UNIQUE (role_resource_id, key)")
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

func (t *resourceRuleTable) TruncateTable(ctx context.Context) error {
	_, err := database.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(resourceRuleTableName).String())
	return err
}

func (r *resourceRuleTable) CreateResourceRule(ctx context.Context, newRule ResourceRule) (err error) {
	s := sqls.INSERT_INTO(resourceRuleTableName)
	s.VALUES("key", s.Param(newRule.Key))
	s.VALUES("value", s.Param(newRule.Value))
	s.VALUES("role_resource_id", s.Param(newRule.RoleResourceID))

	logger.Debug(s.String())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *resourceRuleTable) DeleteResourceRule(ctx context.Context, selector ResourceRuleSelector) (err error) {
	s := sqls.DELETE_FROM(resourceRuleTableName)

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
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *resourceRuleTable) CountResourceRules(ctx context.Context, selector ResourceRuleSelector) (result int64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(resourceRuleTableName)

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
	rowQuery := database.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *resourceRuleTable) QueryResourceRules(ctx context.Context, selector ResourceRuleSelector, pagination *Pagination, sort *Sort) (result []*ResourceRule, err error) {
	s := sqls.SELECT("id")
	s.SELECT("key")
	s.SELECT("value")
	s.SELECT("role_resource_id")
	s.FROM(resourceRuleTableName)

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

	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	defer queryResult.Close()
	for queryResult.Next() {
		roleResourceRule := ResourceRule{}
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
	tx, err := database.Database.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
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
	s.COLUMN("id serial PRIMARY KEY NOT NULL")
	s.COLUMN("user_id int NOT NULL")
	s.COLUMN("role_id int NOT NULL")
	s.COLUMN("created_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("updated_time timestamp without time zone NULL DEFAULT now()")
	s.COLUMN("disabled_time timestamp without time zone NULL")
	s.OPTIONS("CONSTRAINT user_role_union_unique_keys UNIQUE (user_id, role_id)")
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
	_, err := database.Database.ExecContext(ctx, sqls.TRUNCATE_TABLE(userRoleTableName).String())
	return err
}

func (r *userRoleTable) CreateUserRole(ctx context.Context, newRule UserRole) (err error) {
	s := sqls.INSERT_INTO(userRoleTableName)
	s.VALUES("user_id", s.Param(newRule.UserID))
	s.VALUES("role_id", s.Param(newRule.RoleID))

	logger.Debug(s.String())
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return
	}

	return
}

func (r *userRoleTable) DeleteUserRole(ctx context.Context, selector UserRoleSelector) (err error) {
	s := sqls.DELETE_FROM(userRoleTableName)

	if selector.ID == nil && selector.UserID == nil && selector.RoleID == nil {
		return fmt.Errorf("selector conditions cannot all be empty")
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
	_, err = database.Database.ExecContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
	}

	return
}

func (r *userRoleTable) CountUserRoles(ctx context.Context, selector UserRoleSelector) (result int64, err error) {
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
	rowQuery := database.Database.QueryRowContext(ctx, s.String(), s.Params()...)
	if err = rowQuery.Scan(&result); err != nil {
		logger.Error(err)
	}

	return
}

func (r *userRoleTable) QueryUserRoles(ctx context.Context, selector UserRoleSelector, pagination *Pagination, sort *Sort) (result []*UserRole, err error) {
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

	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
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

type permission struct {
}

var Permission = &permission{}

type UserResourcePermissionRule struct {
	Key   string
	Value string
}

type UserResourcePermissionSummary struct {
	Name   string
	Action string
	Rules  []*UserResourcePermissionRule
}

func (srv *UserResourcePermissionSummary) ToProto() *proto.UserResourceSummary {
	result := &proto.UserResourceSummary{}
	result.Name = srv.Name

	if srv.Action == ActionInsert {
		result.Action = proto.ResourceAction_Insert
	}
	if srv.Action == ActionDelete {
		result.Action = proto.ResourceAction_Delete
	}
	if srv.Action == ActionUpdate {
		result.Action = proto.ResourceAction_Update
	}
	if srv.Action == ActionQuery {
		result.Action = proto.ResourceAction_Query
	}

	for _, rule := range srv.Rules {
		result.Rules = append(result.Rules, &proto.UserResourceRule{
			Key:   rule.Key,
			Value: rule.Value,
		})
	}

	return result
}

func (p *UserResourcePermissionSummary) HasOwner() bool {
	return p.Name == "owner"
}

func (p *UserResourcePermissionSummary) MatchRules(name string, action string, rules ...UserResourcePermissionRule) bool {
	if p.Name != name || p.Action != action {
		return false
	}

	for _, targetRule := range rules {
		for _, sourceRule := range p.Rules {
			matchKey := targetRule.Key == sourceRule.Key
			matchValue := targetRule.Value == sourceRule.Value
			if matchKey && matchValue {
				return true
			}
		}
	}

	return false
}

func (src *UserResourcePermissionSummary) LoadProto(data *proto.UserResourceSummary) {
	// 加载 data 信息到 src 上
	src.Name = data.GetName()

	if data.Action == proto.ResourceAction_Insert {
		src.Action = ActionInsert
	}

	if data.Action == proto.ResourceAction_Delete {
		src.Action = ActionDelete
	}

	if data.Action == proto.ResourceAction_Update {
		src.Action = ActionUpdate
	}

	if data.Action == proto.ResourceAction_Query {
		src.Action = ActionQuery
	}

	src.Rules = make([]*UserResourcePermissionRule, len(data.GetRules()))
	for i, rule := range data.GetRules() {
		src.Rules[i] = &UserResourcePermissionRule{
			Key:   rule.GetKey(),
			Value: rule.GetValue(),
		}
	}
}

type UserResourceSummarySelector struct {
	RoleID *int64
	UserID *int64
}

// 初始化管理员权限以及用户默认的配置
func (r *permission) InitDefaultPermissions(ctx context.Context) (err error) {
	adminName := "all"
	adminDescription := "This represents all resources"
	adminResource := &Resource{Name: &adminName, Description: &adminDescription}

	ownerName := "owner"
	ownerDescription := "This represents the user’s own resources"
	ownerResource := &Resource{Name: &ownerName, Description: &ownerDescription}

	createResource := func(resource *Resource) (*Resource, error) {
		selector := ResourceSelector{Name: resource.Name}
		count, err := ResourceTable.CountResources(ctx, selector)
		if err != nil {
			return nil, err
		}
		if count <= 0 {
			if err := ResourceTable.CreateResource(ctx, *resource); err != nil {
				return nil, err
			}
		}
		result, err := ResourceTable.QueryResources(ctx, selector, nil, nil)
		if err != nil {
			return nil, err
		}

		return result[0], nil
	}

	adminResource, err = createResource(adminResource)
	if err != nil {
		return err
	}

	ownerResource, err = createResource(ownerResource)
	if err != nil {
		return err
	}

	adminName = "admin"
	adminDescription = "This represents all role"
	adminRole := &Role{Name: &adminName, Description: &adminDescription}

	ownerName = "owner"
	ownerDescription = "This represents the user’s own role"
	ownerRole := &Role{Name: &ownerName, Description: &ownerDescription}

	createRole := func(role *Role) (*Role, error) {
		selector := RoleSelector{Name: role.Name}
		count, err := RoleTable.CountRoles(ctx, selector)
		if err != nil {
			return nil, err
		}
		if count <= 0 {
			if err := RoleTable.CreateRole(ctx, *role); err != nil {
				return nil, err
			}
		}
		result, err := RoleTable.QueryRoles(ctx, selector, nil, nil)
		if err != nil {
			return nil, err
		}

		return result[0], nil
	}

	adminRole, err = createRole(adminRole)
	if err != nil {
		return err
	}

	ownerRole, err = createRole(ownerRole)
	if err != nil {
		return err
	}

	adminRoleResource := &RoleResource{RoleID: *adminRole.ID, ResourceID: *adminResource.ID}
	ownerRoleResource := &RoleResource{RoleID: *ownerRole.ID, ResourceID: *ownerResource.ID}

	createRoleResource := func(role *RoleResource) (*RoleResource, error) {
		selector := RoleResourceSelector{RoleID: &role.RoleID, ResourceID: &role.ResourceID, Action: &role.Action}
		count, err := RoleResourceTable.CountRoleResources(ctx, selector)
		if err != nil {
			return nil, err
		}
		if count <= 0 {
			if err := RoleResourceTable.CreateRoleResource(ctx, *role); err != nil {
				return nil, err
			}
		}
		result, err := RoleResourceTable.QueryRoleResources(ctx, selector, nil, nil)
		if err != nil {
			return nil, err
		}

		return result[0], nil
	}

	adminRoleResource.Action = ActionInsert
	_, err = createRoleResource(adminRoleResource)
	if err != nil {
		return err
	}

	adminRoleResource.Action = ActionDelete
	_, err = createRoleResource(adminRoleResource)
	if err != nil {
		return err
	}

	adminRoleResource.Action = ActionUpdate
	_, err = createRoleResource(adminRoleResource)
	if err != nil {
		return err
	}

	adminRoleResource.Action = ActionQuery
	_, err = createRoleResource(adminRoleResource)
	if err != nil {
		return err
	}

	ownerRoleResource.Action = ActionInsert
	_, err = createRoleResource(ownerRoleResource)
	if err != nil {
		return err
	}

	ownerRoleResource.Action = ActionDelete
	_, err = createRoleResource(ownerRoleResource)
	if err != nil {
		return err
	}

	ownerRoleResource.Action = ActionUpdate
	_, err = createRoleResource(ownerRoleResource)
	if err != nil {
		return err
	}

	ownerRoleResource.Action = ActionQuery
	_, err = createRoleResource(ownerRoleResource)
	if err != nil {
		return err
	}

	createUserRole := func(userRole *UserRole) (*UserRole, error) {
		selector := UserRoleSelector{UserID: userRole.ID, RoleID: &userRole.RoleID}
		count, err := UserRoleTable.CountUserRoles(ctx, selector)
		if err != nil {
			return nil, err
		}
		if count <= 0 {
			if err := UserRoleTable.CreateUserRole(ctx, *userRole); err != nil {
				return nil, err
			}
		}
		result, err := UserRoleTable.QueryUserRoles(ctx, selector, nil, nil)
		if err != nil {
			return nil, err
		}

		return result[0], nil
	}

	adminUserRole := &UserRole{UserID: 1, RoleID: *adminRole.ID}
	if _, err := createUserRole(adminUserRole); err != nil {
		return err
	}

	ownerUserRole := &UserRole{UserID: 1, RoleID: *ownerRole.ID}
	if _, err := createUserRole(ownerUserRole); err != nil {
		return err
	}
	return nil
}

func (r *permission) QueryUserResourceSummaries(ctx context.Context, selector UserResourceSummarySelector) (result []*UserResourcePermissionSummary, err error) {
	// Build the SQL query to retrieve user resource summaries from the database.
	s := sqls.SELECT()
	s.SELECT("d.name AS name")
	s.SELECT("c.action AS action")
	s.SELECT("e.key AS key")
	s.SELECT("e.value AS value")
	s.FROM(fmt.Sprintf("%s AS a", userRoleTableName))
	s.LEFT_OUTER_JOIN(fmt.Sprintf("%s AS b ON a.role_id=b.id", roleTableName))
	s.LEFT_OUTER_JOIN(fmt.Sprintf("%s AS c ON b.id=c.role_id", roleResourceTableName))
	s.LEFT_OUTER_JOIN(fmt.Sprintf("%s AS d ON c.resource_id=d.id", resourceTableName))
	s.LEFT_OUTER_JOIN(fmt.Sprintf("%s AS e ON c.id=e.role_resource_id", resourceRuleTableName))

	if selector.UserID != nil {
		s.WHERE(fmt.Sprintf("a.user_id=%s", s.Param(selector.UserID)))
	}

	if selector.RoleID != nil {
		s.WHERE(fmt.Sprintf("a.role_id=%s", s.Param(selector.RoleID)))
	}

	// Log the query string for debugging purposes.
	logger.Debug(s.String())

	// Execute the query and retrieve the result set.
	queryResult, err := database.Database.QueryContext(ctx, s.String(), s.Params()...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// Close the result set once we are done.
	defer queryResult.Close()

	// Map to store the user resource summaries, keyed by the cache key.
	userResourceSummaryMap := make(map[string]*UserResourcePermissionSummary)

	// Iterate over the result set and populate the user resource summary map.
	for queryResult.Next() {
		cacheRule := struct {
			Name   string
			Action string
			Key    *string
			Value  *string
		}{}

		// Scan the result set row into the cacheRule struct.
		if err = queryResult.Scan(
			&cacheRule.Name,
			&cacheRule.Action,
			&cacheRule.Key,
			&cacheRule.Value,
		); err != nil {
			logger.Error(err)
			return nil, err
		}

		// Generate the cache key for the user resource summary.
		cacheKey := fmt.Sprintf("%s-%s", cacheRule.Action, cacheRule.Name)

		// If the cache key doesn't exist in the map, create a new UserResourceSummary object.
		if _, has := userResourceSummaryMap[cacheKey]; !has {
			userResourceSummaryMap[cacheKey] = &UserResourcePermissionSummary{
				Name:   cacheRule.Name,
				Action: cacheRule.Action,
				Rules:  []*UserResourcePermissionRule{},
			}
		}

		// Add the user resource rule to the user resource summary.
		userResourceSummary := userResourceSummaryMap[cacheKey]
		rule := &UserResourcePermissionRule{}

		if cacheRule.Key != nil {
			rule.Key = *cacheRule.Key
		}

		if cacheRule.Value != nil {
			rule.Value = *cacheRule.Value
		}

		if cacheRule.Key != nil || cacheRule.Value != nil {
			userResourceSummary.Rules = append(userResourceSummary.Rules, rule)
		}
	}

	// Convert the user resource summary map to a list and return.
	for _, data := range userResourceSummaryMap {
		result = append(result, data)
	}

	return
}
