package dao

import (
	"testing"

	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/easysql"
)

func TestCreateUserTable(t *testing.T) {
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8")
	createUserTable()
	t.Error("test")
}

func TestCreateUser(t *testing.T) {
	config.Set("encrypt-password", "test")
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8")
	CreateUser("test", "test2", "test2", "test")
	t.Error("test")
}

func TestQueryUserByID(t *testing.T) {
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8")
	QueryUserByID(1)
	t.Error("test")
}

func TestQueryUserByUsername(t *testing.T) {
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8")
	QueryUserByUsername("")
	t.Error("test")
}
