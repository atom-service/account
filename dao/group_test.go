package dao

import (
	"fmt"
	"os"
	"testing"

	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/easysql"
	"github.com/yinxulai/goutils/sqldb"
)

func TestMain(m *testing.M) {
	fmt.Println("准备测试环境")                                                          // 日志
	config.Set("encrypt-password", "encrypt-password")                             // 密码加密中会用到的
	sqldb.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4")   // 测试数据库
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4") // 测试数据库
	MustInitTables()                                                               // 初始化测试数据库
	fmt.Println("开始执行测试")                                                          // 日志
	exitCode := m.Run()                                                            // 执行测试
	fmt.Println("测试执行完成,清理测试数据")                                                   // 日志
	TruncateTables()                                                               // 重置测试数据库
	os.Exit(exitCode)                                                              // 推出
}

func TestCreateGroup(t *testing.T) {
	type args struct {
		name        string
		class       string
		state       string
		description string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"正常创建", args{"test", "test", "state", "描述"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateGroup(tt.args.name, tt.args.class, tt.args.state, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
