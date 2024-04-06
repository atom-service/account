package model

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/atom-service/account/package/proto"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/yinxulai/goconf"
)

var Database *sql.DB

func init() {
	goconf.Declare("postgres_uri", "postgresql://postgres:password@localhost/account", true, "postgres 的数据库连接 uri")
}

func InitDB(ctx context.Context) error {
	newDB, err := sql.Open("pgx", goconf.MustGet("postgres_uri"))
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	var version string
	versionQuery := newDB.QueryRowContext(ctx, "SELECT version()")
	if err = versionQuery.Scan(&version); err != nil {
		return fmt.Errorf("failed to query database version: %v", err)
	}

	slog.DebugContext(ctx, "Server run on database: %s\n", slog.String("version", version))

	newDB.SetMaxOpenConns(10)
	newDB.SetMaxIdleConns(3)
	Database = newDB
	return nil
}

func Init(ctx context.Context) error {
	if err := InitDB(ctx); err != nil {
		return err
	}

	// 先初始化用户表
	if err := UserTable.InitTable(ctx); err != nil {
		return err
	}

	// 再初始化权限表
	if err := RoleTable.InitTable(ctx); err != nil {
		return err
	}
	if err := UserRoleTable.InitTable(ctx); err != nil {
		return err
	}
	if err := ResourceTable.InitTable(ctx); err != nil {
		return err
	}
	if err := RoleResourceTable.InitTable(ctx); err != nil {
		return err
	}
	if err := RoleResourceRuleTable.InitTable(ctx); err != nil {
		return err
	}

	if err := Permission.InitDefaultPermissions(ctx); err != nil {
		return err
	}

	// 最后初始化用户信息表
	if err := SettingTable.InitTable(ctx); err != nil {
		return err
	}
	if err := SecretTable.InitTable(ctx); err != nil {
		return err
	}
	if err := LabelTable.InitTable(ctx); err != nil {
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
