package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/grpcbrick/account/dao"
	"github.com/grpcbrick/account/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/sqldb"
)

func TestMain(m *testing.M) {
	fmt.Println("准备测试环境")                                                                       // 日志
	config.Set("encrypt_password", "encrypt_password")                                          // 密码加密中会用到的
	sqldb.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4&parseTime=true") // 测试数据库
	dao.InitTables()                                                                            // 初始化测试数据库
	InitTestData()                                                                              // 预插入一条用户数据
	fmt.Println("开始执行测试")                                                                       // 日志
	exitCode := m.Run()                                                                         // 执行测试
	fmt.Println("测试执行完成,清理测试数据")                                                                // 日志
	dao.TruncateTables()                                                                        // 重置测试数据库
	os.Exit(exitCode)                                                                           // 推出
}

func InitTestData() {
	// 预创建一条测试用户数据
	// 方便 label、group 测试
	srv := NewService()
	resp, err := srv.CreateUser(
		context.Background(),
		&standard.CreateUserRequest{Class: "InitTestData", Inviter: 0, Nickname: "InitTestData", Username: "InitTestData", Password: "InitTestData"},
	)
	if err != nil {
		panic(err)
	}
	if resp.State != standard.State_SUCCESS {
		panic(fmt.Errorf("准备测试数据失败、用户创建失败: %v", resp))
	}

	fmt.Printf("预插入一条用户数据 ID 为: %d \n", resp.Data.ID)
}

func TestService_CreateUser(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.CreateUserRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常创建", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password"},
			standard.State_SUCCESS, false},

		{"重复的 Username", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password"},
			standard.State_USER_ALREADY_EXISTS, false},

		{"空的 Class", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username1", Password: "Password"},
			standard.State_SUCCESS, false},

		{"空的 Inviter", &standard.CreateUserRequest{Class: "TEST", Nickname: "Nickname", Username: "Username2", Password: "Password"},
			standard.State_SUCCESS, false},

		{"空的 Nickname", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Username: "Username3", Password: "Password"},
			standard.State_PARAMS_INVALID, false},

		{"空的 Username", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Password: "Username4"},
			standard.State_PARAMS_INVALID, false},

		{"空的 Password", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username5"},
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
				return
			}
			if gotResp.State == standard.State_SUCCESS {
				if gotResp.Data.Username != tt.args.Username {
					t.Errorf("Service.CreateUser() = %v, want %v", gotResp, tt.wantState)
					return
				}
			}
		})
	}
}

func TestService_QueryUsers(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name            string
		args            *standard.QueryUsersRequest
		wantState       standard.State
		wantDataSize    int64
		wantCurrentPage int64
		wantTotalPage   int64
		wantErr         bool
	}{
		{"正常查询", &standard.QueryUsersRequest{Page: 1, Limit: 90},
			standard.State_SUCCESS, 4, 1, 1, false},
		{"只查一条", &standard.QueryUsersRequest{Page: 1, Limit: 1},
			standard.State_SUCCESS, 1, 1, 4, false},
		{"第二页", &standard.QueryUsersRequest{Page: 1, Limit: 2},
			standard.State_SUCCESS, 2, 1, 2, false},
		{"空的 ID", &standard.QueryUsersRequest{Page: 0, Limit: 0},
			standard.State_PARAMS_INVALID, 0, 0, 0, false},
		{"不存在的 ID", &standard.QueryUsersRequest{Page: 0, Limit: 0},
			standard.State_PARAMS_INVALID, 0, 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.QueryUsers(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.QueryUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.QueryUsers() = %v, want %v", gotResp, tt.wantState)
				return
			}

			if tt.wantState == standard.State_SUCCESS {
				if gotResp.TotalPage != tt.wantTotalPage {
					t.Errorf("Service.QueryUsers() = %v, want %d", gotResp, tt.wantTotalPage)
					return
				}
				if int64(len(gotResp.Data)) != tt.wantDataSize {
					t.Errorf("Service.QueryUsers() = %v, want %d", gotResp, tt.wantDataSize)
					return
				}
				if gotResp.CurrentPage != tt.wantCurrentPage {
					t.Errorf("Service.QueryUsers() = %v, want %d", gotResp, tt.wantCurrentPage)
					return
				}
			}
		})
	}
}

func TestService_QueryUserByID(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name         string
		args         *standard.QueryUserByIDRequest
		wantUsername string
		wantState    standard.State
		wantErr      bool
	}{
		{"正常查询", &standard.QueryUserByIDRequest{ID: 2},
			"Username", standard.State_SUCCESS, false},
		{"正常查询", &standard.QueryUserByIDRequest{ID: 3},
			"Username1", standard.State_SUCCESS, false},
		{"正常查询", &standard.QueryUserByIDRequest{ID: 4},
			"Username2", standard.State_SUCCESS, false},
		{"空的 ID", &standard.QueryUserByIDRequest{ID: 0},
			"", standard.State_PARAMS_INVALID, false},
		{"不存在的 ID", &standard.QueryUserByIDRequest{ID: 999},
			"", standard.State_USER_NOT_EXIST, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.QueryUserByID(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.QueryUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.QueryUserByID() = %v, want %v", gotResp, tt.wantState)
				return
			}

			if tt.wantState == standard.State_SUCCESS {
				if gotResp.Data.Username != tt.wantUsername {
					t.Errorf("Service.QueryUserByID() = %v, want %v", gotResp, tt.wantUsername)
					return
				}
			}
		})
	}
}

func TestService_QueryUsersByInviter(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name            string
		args            *standard.QueryUsersByInviterRequest
		wantState       standard.State
		wantDataSize    int64
		wantCurrentPage int64
		wantTotalPage   int64
		wantErr         bool
	}{
		{"正常查询", &standard.QueryUsersByInviterRequest{Inviter: 1, Page: 1, Limit: 90},
			standard.State_SUCCESS, 2, 1, 1, false},
		{"只查一条", &standard.QueryUsersByInviterRequest{Inviter: 1, Page: 1, Limit: 1},
			standard.State_SUCCESS, 1, 1, 2, false},
		{"第二页", &standard.QueryUsersByInviterRequest{Inviter: 1, Page: 1, Limit: 2},
			standard.State_SUCCESS, 2, 1, 1, false},
		{"无效的翻页数据", &standard.QueryUsersByInviterRequest{Inviter: 1, Page: 0, Limit: 0},
			standard.State_PARAMS_INVALID, 0, 0, 0, false},
		{"不存在的 ID", &standard.QueryUsersByInviterRequest{Inviter: 999, Page: 1, Limit: 90},
			standard.State_SUCCESS, 0, 1, 0, false},
		{"空的 ID", &standard.QueryUsersByInviterRequest{Page: 1, Limit: 90},
			standard.State_PARAMS_INVALID, 0, 1, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.QueryUsersByInviter(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.QueryUsersByInviter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.QueryUsersByInviter() = %v, want %v", gotResp, tt.wantState)
				return
			}

			if tt.wantState == standard.State_SUCCESS {
				if gotResp.TotalPage != tt.wantTotalPage {
					t.Errorf("Service.QueryUsersByInviter() = %v, want %d", gotResp, tt.wantTotalPage)
					return
				}
				if int64(len(gotResp.Data)) != tt.wantDataSize {
					t.Errorf("Service.QueryUsersByInviter() = %v, want %d", gotResp, tt.wantDataSize)
					return
				}
				if gotResp.CurrentPage != tt.wantCurrentPage {
					t.Errorf("Service.QueryUsersByInviter() = %v, want %d", gotResp, tt.wantCurrentPage)
					return
				}
			}
		})
	}
}

func TestService_QueryUserByUsername(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name         string
		args         *standard.QueryUserByUsernameRequest
		wantUsername string
		wantState    standard.State
		wantErr      bool
	}{
		{"正常查询", &standard.QueryUserByUsernameRequest{Username: "Username"},
			"Username", standard.State_SUCCESS, false},
		{"正常查询", &standard.QueryUserByUsernameRequest{Username: "Username1"},
			"Username1", standard.State_SUCCESS, false},
		{"正常查询", &standard.QueryUserByUsernameRequest{Username: "Username2"},
			"Username2", standard.State_SUCCESS, false},
		{"空的 Username", &standard.QueryUserByUsernameRequest{Username: ""},
			"", standard.State_PARAMS_INVALID, false},
		{"不存在的 Username", &standard.QueryUserByUsernameRequest{Username: "TESTNAME"},
			"", standard.State_USER_NOT_EXIST, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.QueryUserByUsername(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.QueryUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.QueryUserByUsername() = %v, want %v", gotResp, tt.wantState)
				return
			}

			if tt.wantState == standard.State_SUCCESS {
				if gotResp.Data.Username != tt.wantUsername {
					t.Errorf("Service.QueryUserByUsername() = %v, want %v", gotResp, tt.wantUsername)
					return
				}
			}
		})
	}
}

func TestService_DeleteUserByID(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.DeleteUserByIDRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常删除", &standard.DeleteUserByIDRequest{ID: 2},
			standard.State_SUCCESS, false},
		{"正常删除", &standard.DeleteUserByIDRequest{ID: 3},
			standard.State_SUCCESS, false},
		{"正常删除", &standard.DeleteUserByIDRequest{ID: 4},
			standard.State_SUCCESS, false},
		{"不存在的 ID", &standard.DeleteUserByIDRequest{ID: 9999},
			standard.State_USER_NOT_EXIST, false},
		{"已删除的 ID", &standard.DeleteUserByIDRequest{ID: 2},
			standard.State_SUCCESS, false},
		{"空的 ID", &standard.DeleteUserByIDRequest{ID: 0},
			standard.State_PARAMS_INVALID, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.DeleteUserByID(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.DeleteUserByID() = %v, want %v", gotResp, tt.wantState)
				return
			}
		})
	}
}

func TestService_UpdateUserPasswordByID(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.UpdateUserPasswordByIDRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常更新", &standard.UpdateUserPasswordByIDRequest{ID: 2, Password: "password1"},
			standard.State_SUCCESS, false},
		{"正常更新", &standard.UpdateUserPasswordByIDRequest{ID: 3, Password: "password2"},
			standard.State_SUCCESS, false},
		{"正常更新", &standard.UpdateUserPasswordByIDRequest{ID: 4, Password: "password3"},
			standard.State_SUCCESS, false},
		{"对不存在的用户 ID 更新", &standard.UpdateUserPasswordByIDRequest{ID: 99999, Password: "password4"},
			standard.State_USER_NOT_EXIST, false},
		{"对存在的用户使用空密码更新", &standard.UpdateUserPasswordByIDRequest{ID: 2, Password: ""},
			standard.State_PARAMS_INVALID, false},
		// TODO: 测试密码正则

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.UpdateUserPasswordByID(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateUserPasswordByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.UpdateUserPasswordByID() = %v, want %v", gotResp, tt.wantState)
				return
			}
		})
	}
}

func TestService_VerifyUserPasswordByID(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.VerifyUserPasswordByIDRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常验证密码1", &standard.VerifyUserPasswordByIDRequest{ID: 2, Password: "password1"},
			standard.State_SUCCESS, false},
		{"正常验证密码2", &standard.VerifyUserPasswordByIDRequest{ID: 3, Password: "password2"},
			standard.State_SUCCESS, false},
		{"正常验证密码3", &standard.VerifyUserPasswordByIDRequest{ID: 4, Password: "password3"},
			standard.State_SUCCESS, false},
		{"正常验证错误密码", &standard.VerifyUserPasswordByIDRequest{ID: 2, Password: "password-error"},
			standard.State_FAILURE, false},
		{"对不存在的用户验证密码", &standard.VerifyUserPasswordByIDRequest{ID: 99999, Password: "password4"},
			standard.State_USER_NOT_EXIST, false},
		{"对存在的用户使用空密码更新", &standard.VerifyUserPasswordByIDRequest{ID: 2, Password: ""},
			standard.State_PARAMS_INVALID, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.VerifyUserPasswordByID(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.VerifyUserPasswordByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.VerifyUserPasswordByID() = %v, want %v", gotResp, tt.wantState)
				return
			}
		})
	}
}

func TestService_VerifyUserPasswordByUsername(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.VerifyUserPasswordByUsernameRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常验证密码1", &standard.VerifyUserPasswordByUsernameRequest{Username: "Username", Password: "password1"},
			standard.State_SUCCESS, false},
		{"正常验证密码2", &standard.VerifyUserPasswordByUsernameRequest{Username: "Username1", Password: "password2"},
			standard.State_SUCCESS, false},
		{"正常验证密码3", &standard.VerifyUserPasswordByUsernameRequest{Username: "Username2", Password: "password3"},
			standard.State_SUCCESS, false},
		{"正常验证错误密码", &standard.VerifyUserPasswordByUsernameRequest{Username: "Username", Password: "password-error"},
			standard.State_FAILURE, false},
		{"对不存在的用户验证密码", &standard.VerifyUserPasswordByUsernameRequest{Username: "NullUsername", Password: "password4"},
			standard.State_USER_NOT_EXIST, false},
		{"对存在的用户使用空密码验证", &standard.VerifyUserPasswordByUsernameRequest{Username: "Username", Password: ""},
			standard.State_PARAMS_INVALID, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.VerifyUserPasswordByUsername(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.VerifyUserPasswordByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.VerifyUserPasswordByUsername() = %v, want %v", gotResp, tt.wantState)
				return
			}
		})
	}
}
