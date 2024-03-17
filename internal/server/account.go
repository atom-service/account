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
		result.State = protos.State_PARAMS_INVALID
		return
	}

	// 验证密码是否正确
	if *queryResult[0].Password != passwordHash {
		result.State = protos.State_PARAMS_INVALID
		return
	}

	querySecretResult, err := model.SecretTable.QuerySecrets(ctx, model.SecretSelector{UserID: queryResult[0].ID}, nil, nil)
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
		result.State = protos.State_PARAMS_INVALID
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

	users, err := model.UserTable.QueryUsers(ctx, model.UserSelector{Username: &request.Username}, nil, nil)
	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if len(users) <= 0 {
		result.State = protos.State_FAILURE
		return
	}

	// 创建一组 system AK/SK
	err = model.SecretTable.CreateSecret(ctx, model.CreateSecretParams{
		UserID: *users[0].ID,
		Type:   model.SystemSecretType,
	})

	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

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

func (s *AccountServer) CreateSecret(ctx context.Context, request *protos.CreateSecretRequest) (result *protos.CreateSecretResponse, err error) {
	result = &protos.CreateSecretResponse{}

	user := auth.ResolveUserFromIncomingContext(ctx)
	if user == nil {
		result.State = protos.State_NO_PERMISSION
		result.Message = "Not logged in"
		return
	}

	result = &protos.CreateSecretResponse{}

	err = model.SecretTable.CreateSecret(ctx, model.CreateSecretParams{
		UserID: *user.ID,
		Type:   model.UserSecretType,
	})

	if err != nil {
		result.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	result.State = protos.State_SUCCESS
	return
}

func (s *AccountServer) DisableSecret(ctx context.Context, request *protos.DisableSecretRequest) (result *protos.DisableSecretResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableSecret not implemented")
}

func (s *AccountServer) DeleteSecret(ctx context.Context, request *protos.DeleteSecreteRequest) (result *protos.DeleteSecreteResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}

func (s *AccountServer) QuerySecrets(ctx context.Context, request *protos.QuerySecretsRequest) (result *protos.QuerySecretsResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySecrets not implemented")
}
