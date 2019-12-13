package packtest

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/grpcbrick/account/dao"
	"github.com/grpcbrick/account/provider"
	"github.com/grpcbrick/account/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/sqldb"
)

func TestMain(m *testing.M) {
	fmt.Println("准备测试环境")                                                                       // 日志
	config.Set("encrypt-password", "encrypt-password")                                          // 密码加密中会用到的
	sqldb.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4&parseTime=true") // 测试数据库
	dao.InitTables()                                                                            // 初始化测试数据库
	fmt.Println("开始执行测试")                                                                       // 日志
	exitCode := m.Run()                                                                         // 执行测试
	fmt.Println("测试执行完成,清理测试数据")                                                                // 日志
	dao.TruncateTables()                                                                        // 重置测试数据库
	os.Exit(exitCode)                                                                           // 推出
}

func TestCreateUser(t *testing.T) {
	srv := provider.NewService()
	tests := []struct {
		name string
		args *standard.CreateUserRequest
	}{
		{"正常创建", &standard.CreateUserRequest{
			Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password",
		}},
		{"正常创建", &standard.CreateUserRequest{
			Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password",
		}},
		{"正常创建", &standard.CreateUserRequest{
			Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.CreateUser(context.Background(), tt.args)
			if err != nil {
				t.Errorf("Service.CreateGroup() error = %v", err)
				return
			}

			// 查询校验
			queryResp, err := srv.QueryUserByID(
				context.Background(),
				&standard.QueryUserByIDRequest{ID: gotResp.Data.ID},
			)

			if err != nil {
				t.Errorf("Service.QueryUserByID() error = %v", err)
				return
			}

			if queryResp.Data.Class != tt.args.Class {
				t.Errorf("Service.QueryUserByID() Class = %s expect = %s", queryResp.Data.Class, tt.args.Class)
				return
			}

			if queryResp.Data.Nickname != "Nickname" {
				t.Errorf("Service.QueryUserByID() Nickname = %s expect = %s", queryResp.Data.Nickname, tt.args.Nickname)
				return
			}

			if queryResp.Data.Username != "Username" {
				t.Errorf("Service.QueryUserByID() Username = %s expect = %s", queryResp.Data.Username, tt.args.Username)
				return
			}

			if queryResp.Data.Inviter != 1 {
				t.Errorf("Service.QueryUserByID() Inviter = %d expect = %d", queryResp.Data.Inviter, tt.args.Inviter)
				return
			}

		})
	}
}

func CreateLabelTest() {

}

func LabelMappingTest() {

}

func CreateGroupTest() {

}

func GroupMappingTest() {

}
