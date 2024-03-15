package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/atom-service/account/internal/db"
	"github.com/atom-service/common/logger"
	"github.com/yinxulai/sqls"
)

var permissionSchema = "\"permission\""
var roleTableName = permissionSchema + ".\"roles\""
var resourceTableName = permissionSchema + ".\"resources\""
var roleResourceTableName = permissionSchema + ".\"role—resource\""

type RoleResource struct {
	ResourceID int64
	Action     int64
	Rules      interface{}
}

type Role struct {
	ID          *int64
	Name        *string
	Description *string
	CreatedTime *time.Time `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
	DeletedTime *time.Time `json:"deleted_time"`
}

type Resource struct {
	ID          *int64
	Name        *string
	Description *string
	CreatedTime *time.Time `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
	DeletedTime *time.Time `json:"deleted_time"`
}

type RoleSelector struct {
	ID   *int64
	Name *string
}

type ResourceSelector struct {
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
	tx.ExecContext(ctx, sqls.CREATE_SCHEMA(permissionSchema).String())

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
	tx.ExecContext(ctx, s.String())
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

func (r *roleTable) CountRoles(ctx context.Context, selector RoleSelector) (result uint64, err error) {
	s := sqls.SELECT("COUNT(*) AS count").FROM(roleTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	s.WHERE("deleted_time<=" + s.Param(time.Now()))

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
	s.SELECT("deleted_time")
	s.SELECT("disabled_time")
	s.FROM(roleTableName)

	if selector.ID != nil {
		s.WHERE("id=" + s.Param(selector.ID))
	}
	if selector.Name != nil {
		s.WHERE("name=" + s.Param(selector.Name))
	}

	s.WHERE("deleted_time<=" + s.Param(time.Now()))
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
		role := Role{}
		if err = queryResult.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedTime,
			&role.UpdatedTime,
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

type resourceTable struct{}


type roleResourceTable struct{}
