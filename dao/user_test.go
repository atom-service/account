package dao

import (
	"testing"

	"github.com/yinxulai/goutils/easysql"
)

func TestCreateUserTable(t *testing.T) {
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8")
	createUserTable()
	t.Error("test")
}

func TestCreateUser(t *testing.T) {
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8")
	createUser("test", "test", "test", "test")
	t.Error("test")
}
