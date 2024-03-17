package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/atom-service/common/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Database *sql.DB

func init() {
	config.Declare("postgres_uri", "postgresql://postgres:password@localhost/account", true, "postgres 的数据库连接 uri")
}

func Init() {
	newDB, err := sql.Open("pgx", config.MustGet("postgres_uri"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	var version string
	versionQuery := newDB.QueryRow("SELECT version()")
	if versionQuery.Scan(&version); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to query database version: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf(version)

	newDB.SetMaxOpenConns(10)
	newDB.SetMaxIdleConns(3)
	Database = newDB
}
