package provider

import (
	"context"
	"reflect"
	"testing"

	"github.com/grpcbrick/account/standard"
)

func TestService_CreateLabel(t *testing.T) {
	type args struct {
		ctx context.Context
		req *standard.CreateLabelRequest
	}
	tests := []struct {
		name     string
		srv      *Service
		args     args
		wantResp *standard.CreateLabelResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Service{}
			gotResp, err := srv.CreateLabel(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Service.CreateLabel() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
