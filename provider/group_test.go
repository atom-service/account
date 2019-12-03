package provider

import (
	"context"
	"testing"

	"github.com/grpcbrick/account/standard"
)

func TestService_CreateGroup(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.CreateGroupRequest
		wantState standard.State
		wantErr   bool
	}{
		{"创建分组测试:正常创建", &standard.CreateGroupRequest{Name: "TEST", Class: "Class", State: "State", Description: "Description"},
			standard.State_SUCCESS, false},

		{"创建分组测试:重复的 Name", &standard.CreateGroupRequest{Name: "TEST", Class: "Class", State: "State", Description: "Description"},
			standard.State_GROUP_ALREADY_EXISTS, false},

		{"创建分组测试:空的 Name", &standard.CreateGroupRequest{Name: "", Class: "Class", State: "Nickname", Description: "Username"},
			standard.State_PARAMS_INVALID, false},

		{"创建分组测试:空的 Class", &standard.CreateGroupRequest{Name: "TEST", Class: "", State: "Nickname", Description: "Username"},
			standard.State_PARAMS_INVALID, false},

		{"创建分组测试:空的 State", &standard.CreateGroupRequest{Name: "TEST", Class: "Class", State: "", Description: "Username"},
			standard.State_PARAMS_INVALID, false},

		{"创建分组测试:空的 Description", &standard.CreateGroupRequest{Name: "TEST", Class: "Class", State: "Nickname", Description: ""},
			standard.State_PARAMS_INVALID, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.CreateGroup(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.CreateGroup() = %v, want %v", gotResp, tt.wantState)
				return
			}
		})
	}
}
