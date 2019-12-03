package provider

import (
	"context"
	"testing"

	"github.com/grpcbrick/account/standard"
)

func TestService_CreateLabel(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		wantErr   bool
		wantState standard.State
		args      *standard.CreateLabelRequest
	}{
		{"测试正常创建", false, standard.State_SUCCESS, &standard.CreateLabelRequest{Name: "TEST", Class: "Class", State: "Nickname", Value: "Username"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.CreateLabel(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.CreateLabel() = %s, want %v", gotResp.Message, tt.wantState)
				return
			}
		})
	}
}
