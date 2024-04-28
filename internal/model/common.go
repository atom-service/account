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

var (
	UserResourceName                      = "user"
	UserResourceDescription               = "用户管理"
	UserLabelResourceName                 = "user.label"
	UserLabelResourceDescription          = "用户标签管理"
	UserSecretResourceName                = "user.secret"
	UserSecretResourceDescription         = "用户密钥管理"
	UserSettingResourceName               = "user.setting"
	UserSettingResourceDescription        = "用户设置管理"
	PermissionResourceName                = "permission"
	PermissionResourceDescription         = "权限管理"
	PermissionRoleResourceName            = "permission.role"
	PermissionRoleResourceDescription     = "权限角色管理"
	PermissionResourceResourceName        = "permission.resource"
	PermissionResourceResourceDescription = "权限资源管理"
)

var initPermissionResources = []Resource{
	{Name: &UserResourceName, Description: &UserResourceDescription},
	{Name: &UserLabelResourceName, Description: &UserLabelResourceDescription},
	{Name: &UserSecretResourceName, Description: &UserSecretResourceDescription},
	{Name: &UserSettingResourceName, Description: &UserSettingResourceDescription},
	{Name: &PermissionResourceName, Description: &PermissionResourceDescription},
	{Name: &PermissionRoleResourceName, Description: &PermissionRoleResourceDescription},
	{Name: &PermissionResourceResourceName, Description: &PermissionResourceResourceDescription},
}

func init() {
	goconf.Declare("postgres_uri", "postgresql://postgres:password@localhost/account", true, "Postgres database connection uri")
}

func InitDB(ctx context.Context) {
	newDB, err := sql.Open("pgx", goconf.MustGet("postgres_uri"))
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %v", err))
	}

	var version string
	versionQuery := newDB.QueryRowContext(ctx, "SELECT version()")
	if err = versionQuery.Scan(&version); err != nil {
		panic(fmt.Errorf("failed to query database version: %v", err))
	}

	slog.DebugContext(ctx, "Server run on database: %s\n", slog.String("version", version))

	newDB.SetMaxOpenConns(10)
	newDB.SetMaxIdleConns(3)
	Database = newDB
}

func Init(ctx context.Context) error {
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	InitDB(ctx)

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

	if err := Permission.InitDefaultPermissions(ctx); err != nil {
		return err
	}

	if err := Permission.AutoCreateResources(ctx, initPermissionResources); err != nil {
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
