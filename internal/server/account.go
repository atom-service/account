package server

import (
	"context"
	"time"

	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountServer struct {
	protos.UnimplementedAccountServiceServer
}

func NewAccountServer() *AccountServer {
	return &AccountServer{}
}

func (s *AccountServer) SignIn(ctx context.Context, request *protos.SignInRequest) (result *protos.SignInResponse, err error) {
	result = &protos.SignInResponse{}

	var userSelector model.UserSelector
	userSelector.LoadProtoStruct(request.Selector)
	passwordHash := model.Password.Hash(request.Password)
	queryResult, err := model.UserTable.QueryUsers(ctx, userSelector, nil, nil)
	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if len(queryResult) <= 0 {
		result.State = protos.State_USER_NOT_EXIST
		return
	}

	// 验证密码是否正确
	if *queryResult[0].Password != passwordHash {
		result.State = protos.State_PARAMS_INVALID
		return
	}

	querySecretResult, err := model.SecretTable.QuerySecrets(ctx, model.SecretSelector{OwnerID: queryResult[0].ID}, nil, nil)
	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if len(querySecretResult) <= 0 {
		result.Message = "No Secret available"
		result.State = protos.State_FAILURE
		logger.Errorf(result.Message)
		return
	}

	firstSecret := querySecretResult[0]
	token := auth.SignToken(*firstSecret.Key, *firstSecret.Value, auth.SignData{
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})

	result.AuthenticationToken = token
	result.State = protos.State_SUCCESS
	return
}

func (s *AccountServer) SignUp(ctx context.Context, request *protos.SignUpRequest) (result *protos.SignUpResponse, err error) {
	result = &protos.SignUpResponse{}

	countResult, err := model.UserTable.CountUsers(ctx, model.UserSelector{
		Username: &request.Username,
	})

	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if countResult > 0 {
		result.State = protos.State_USER_ALREADY_EXISTS
		return
	}

	passwordHash := model.Password.Hash(request.Password)
	err = model.UserTable.CreateUser(ctx, model.User{
		Username: &request.Username,
		Password: &passwordHash,
	})
	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	// TODO 创建一组 AK/SK

	return
}

func (s *AccountServer) SignOut(ctx context.Context, request *protos.SignOutRequest) (result *protos.SignOutResponse, err error) {
	result = &protos.SignOutResponse{}

	user := auth.ResolveUserFromIncomingContext(ctx)
	if user == nil {
		result.State = protos.State_NO_PERMISSION
		result.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func (s *AccountServer) QueryUsers(ctx context.Context, request *protos.QueryUsersRequest) (result *protos.QueryUsersResponse, err error) {
	result = &protos.QueryUsersResponse{}

	user := auth.ResolveUserFromIncomingContext(ctx)
	if user == nil {
		result.State = protos.State_NO_PERMISSION
		result.Message = "Not logged in"
		return
	}

	var sort model.Sort
	var pagination model.Pagination
	var userSelector model.UserSelector
	sort.LoadProtoStruct(request.Sort)
	pagination.LoadProtoStruct(request.Pagination)
	userSelector.LoadProtoStruct(request.Selector)
	query, err := model.UserTable.QueryUsers(ctx, userSelector, &pagination, &sort)
	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	count, err := model.UserTable.CountUsers(ctx, userSelector)
	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	for _, user := range query {
		result.Data.Users = append(
			result.Data.Users,
			user.OutProtoStruct(),
		)
	}

	result.Data.Total = count
	result.State = protos.State_SUCCESS
	return
}

func (s *AccountServer) DeleteUser(ctx context.Context, request *protos.DeleteUserRequest) (result *protos.DeleteUserResponse, err error) {
	result = &protos.DeleteUserResponse{}

	user := auth.ResolveUserFromIncomingContext(ctx)
	if user == nil {
		result.State = protos.State_NO_PERMISSION
		result.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
