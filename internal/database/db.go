package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/atom-service/common/config"
	"github.com/atom-service/common/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Database *sql.DB

func init() {
	config.Declare("postgres_uri", "postgresql://postgres:password@localhost/account", true, "postgres 的数据库连接 uri")
}

func Init(ctx context.Context) error {
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
