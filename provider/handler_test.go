package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/grpcbrick/account/dao"
	"github.com/grpcbrick/account/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/easysql"
)

func TestMain(m *testing.M) {
	fmt.Println("准备测试环境")                                                          // 日志
	config.Set("encrypt-password", "encrypt-password")                             // 密码加密中会用到的
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4") // 测试数据库
	dao.InitTables()                                                               // 初始化测试数据库
	fmt.Println("开始执行测试")                                                          // 日志
	exitCode := m.Run()                                                            // 执行测试
	fmt.Println("测试执行完成,清理测试数据")                                                   // 日志
	dao.TruncateTables()                                                           // 重置测试数据库
	os.Exit(exitCode)                                                              // 推出
}

func TestService_CreateUser(t *testing.T) {
	fmt.Println("213123")
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.CreateUserRequest
		wantState standard.State
		wantErr   bool
	}{
		{"测试正常创建", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password"},
			standard.State_SUCCESS, false},

		{"重复的 Username", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password"},
			standard.State_USER_ALREADY_EXISTS, false},

		{"测试空的 Class", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username1", Password: "Password"},
			standard.State_SUCCESS, false},

		{"测试空的 Inviter", &standard.CreateUserRequest{Class: "TEST", Nickname: "Nickname", Username: "Username2", Password: "Password"},
			standard.State_SUCCESS, false},

		{"测试空的 Nickname", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Username: "Username3", Password: "Password"},
			standard.State_PARAMS_INVALID, false},

		{"测试空的 Username", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Password: "Username4"},
			standard.State_PARAMS_INVALID, false},

		{"测试空的 Password", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username5"},
			standard.State_PARAMS_INVALID, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.CreateUser(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.CreateUser() = %v, want %v", gotResp, tt.wantState)
			}
		})
	}
}
