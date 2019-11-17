package provider

import (
	"context"
	"reflect"
	"testing"

	"github.com/grpcbrick/account/dao"
	"github.com/grpcbrick/account/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/easysql"
)

func TestService_CreateUser(t *testing.T) {

	config.Set("encrypt-password", "encrypt-password")                             // 密码加密中会用到的
	easysql.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4") // 测试数据库
	dao.InitTables()                                                               // 初始化测试数据库

	srv := NewService()
	tests := []struct {
		name     string
		args     *standard.CreateUserRequest
		wantResp *standard.CreateUserResponse
		wantErr  bool
	}{
		{"测试正常创建", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Username"}, nil, false},
		{"重复的 Username", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Username"}, nil, false},
		{"测试空的 Class", &standard.CreateUserRequest{Class: "TEST1", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Username"}, nil, false},
		{"测试空的 Inviter", &standard.CreateUserRequest{Class: "TEST2", Nickname: "Nickname", Username: "Username", Password: "Username"}, nil, false},
		{"测试空的 Nickname", &standard.CreateUserRequest{Class: "TEST3", Inviter: 1, Username: "Username", Password: "Username"}, nil, false},
		{"测试空的 Username", &standard.CreateUserRequest{Class: "TEST4", Inviter: 1, Nickname: "Nickname", Password: "Username"}, nil, false},
		{"测试空的 Password", &standard.CreateUserRequest{Class: "TEST5", Inviter: 1, Nickname: "Nickname", Username: "Username"}, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.CreateUser(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Service.CreateUser() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}

	dao.TruncateTables() // 重置测试数据库
}
