package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/atom-service/account/package/proto"
	"github.com/atom-service/common/config"
	"github.com/atom-service/common/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Database *sql.DB

func init() {
	config.Declare("postgres_uri", "postgresql://postgres:password@localhost/account", true, "postgres 的数据库连接 uri")
}

func InitDB(ctx context.Context) error {
	newDB, err := sql.Open("pgx", config.MustGet("postgres_uri"))
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	var version string
	versionQuery := newDB.QueryRowContext(ctx, "SELECT version()")
	if versionQuery.Scan(&version); err != nil {
		return fmt.Errorf( "failed to query database version: %v", err)
	}

	logger.Debugf("Server run on database: %s\n", version)

	newDB.SetMaxOpenConns(10)
	newDB.SetMaxIdleConns(3)
	Database = newDB
	return nil
}

func Init(ctx context.Context) error {
	if err := InitDB(ctx); err != nil {
		return err
	}
	if err := UserTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := UserTable.InitAdminUser(ctx); err != nil {
		return err
	}
	if err := SettingTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := SecretTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := LabelTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := RoleTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := UserRoleTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := ResourceTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := RoleResourceTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := RoleResourceRuleTable.CreateTable(ctx); err != nil {
		return err
	}
	if err := Permission.InitDefaultPermissions(ctx); err != nil {
		return err
	}
	return nil
}

type Pagination struct {
	Limit  *int64
	Offset *int64
}

func (srv *Pagination) LoadProto(data *proto.PaginationOption) {
	if data != nil {
		srv.Limit = data.Limit
		srv.Offset = data.Offset
	}
}

type Sort struct {
	Key  string
	Type SortType
}

type SortType = int32

const (
	SortAsc  = SortType(0)
	SortDesc = SortType(1)
)

func (srv *Sort) LoadProto(data *proto.SortOption) {
	if data != nil {
		srv.Key = data.Key
		srv.Type = int32(data.Type.Number())
	}
}
