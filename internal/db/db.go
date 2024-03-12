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
	config.Declare("postgres_uri", "postgresql://postgres:password@postgresql/account", true, "postgres 的数据库连接 uri")
}

func Init() {
	newDB, err := sql.Open("pgx", config.MustGet("postgres_uri"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	newDB.SetMaxOpenConns(10)
	newDB.SetMaxIdleConns(3)
	Database = newDB
}