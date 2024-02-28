package server

import (
	"context"

	"github.com/atom-service/account/package/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountServer struct {
	protos.UnimplementedAccountServer
}

func NewAccountServer() *AccountServer {
	return &AccountServer{}
}

func (s *AccountServer) SignIn(ctx context.Context, request *protos.SignInRequest) (*protos.SignInResponse, error) {
	

	return nil, nil
}

func (s *AccountServer) SignUp(ctx context.Context, request *protos.SignUpRequest) (*protos.SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func (s *AccountServer) SignOut(ctx context.Context, request *protos.SignOutRequest) (*protos.SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func (s *AccountServer) QueryUsers(ctx context.Context, request *protos.QueryUsersRequest) (*protos.QueryUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUsers not implemented")
}

func (s *AccountServer) DeleteUser(ctx context.Context, request *protos.DeleteUserRequest) (*protos.DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
