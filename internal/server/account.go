package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var AccountServer = &accountServer{}

type accountServer struct {
	protos.UnimplementedAccountServiceServer
}

func (s *accountServer) InitAdminUser(ctx context.Context) (err error) {
	adminUserID := int64(1)
	userSelector := model.UserSelector{ID: &adminUserID}
	queryResult, err := model.UserTable.QueryUsers(ctx, userSelector, nil, nil)
	if err != nil {
		return err
	}

	if len(queryResult) > 0 {
		return nil
	}

	adminUsername := helper.GenerateRandomString(64)
	adminPassword := helper.GenerateRandomString(128)
	signUpResponse, err := s.SignUp(ctx, &protos.SignUpRequest{Username: adminUsername, Password: adminPassword})
	if err != nil {
		return err
	}

	if signUpResponse.State != protos.State_SUCCESS {
		return fmt.Errorf("create admin user failed: %s", signUpResponse.Message)
	}

	logger.Info("admin are created:")
	logger.Infof("username: %s", adminUsername)
	logger.Infof("password: %s", adminPassword)
	return nil
}

func (s *accountServer) SignIn(ctx context.Context, request *protos.SignInRequest) (response *protos.SignInResponse, err error) {
	response = &protos.SignInResponse{}

	var userSelector model.UserSelector
	userSelector.LoadProtoStruct(request.Selector)
	passwordHash := model.Password.Hash(request.Password)
	queryResult, err := model.UserTable.QueryUsers(ctx, userSelector, nil, nil)
	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if len(queryResult) <= 0 {
		response.State = protos.State_PARAMS_INVALID
		return
	}

	// 验证密码是否正确
	if *queryResult[0].Password != passwordHash {
		response.State = protos.State_PARAMS_INVALID
		return
	}

	querySecretResult, err := model.SecretTable.QuerySecrets(ctx, model.SecretSelector{UserID: queryResult[0].ID}, nil, nil)
	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if len(querySecretResult) <= 0 {
		response.Message = "No Secret available"
		response.State = protos.State_FAILURE
		logger.Errorf(response.Message)
		return
	}

	firstSecret := querySecretResult[0]
	token := auth.SignToken(*firstSecret.Key, *firstSecret.Value, auth.SignData{
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})

	response.AuthenticationToken = token
	response.State = protos.State_SUCCESS
	return
}

func (s *accountServer) SignUp(ctx context.Context, request *protos.SignUpRequest) (response *protos.SignUpResponse, err error) {
	response = &protos.SignUpResponse{}

	countResult, err := model.UserTable.CountUsers(ctx, model.UserSelector{
		Username: &request.Username,
	})

	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if countResult > 0 {
		response.State = protos.State_PARAMS_INVALID
		return
	}

	passwordHash := model.Password.Hash(request.Password)
	err = model.UserTable.CreateUser(ctx, model.User{
		Username: &request.Username,
		Password: &passwordHash,
	})
	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	users, err := model.UserTable.QueryUsers(ctx, model.UserSelector{Username: &request.Username}, nil, nil)
	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	if len(users) <= 0 {
		response.State = protos.State_FAILURE
		return
	}

	// 创建一组 system AK/SK
	err = model.SecretTable.CreateSecret(ctx, model.CreateSecretParams{
		UserID: *users[0].ID,
		Type:   model.SystemSecretType,
	})

	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = protos.State_SUCCESS
	return
}

func (s *accountServer) SignOut(ctx context.Context, request *protos.SignOutRequest) (response *protos.SignOutResponse, err error) {
	response = &protos.SignOutResponse{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func (s *accountServer) QueryUsers(ctx context.Context, request *protos.QueryUsersRequest) (response *protos.QueryUsersResponse, err error) {
	response = &protos.QueryUsersResponse{}
	response.Data = &protos.QueryUsersResponse_DataType{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	var sort model.Sort
	var pagination model.Pagination
	var userSelector model.UserSelector
	sort.LoadProtoStruct(request.Sort)
	pagination.LoadProtoStruct(request.Pagination)
	userSelector.LoadProtoStruct(request.Selector)

	// 先将 ID 范围设置为当前用户
	userSelector.ID = &authData.User.ID
	// 如果是内部调用则设置 ID 范围为参数范围
	if helper.IsGodSecret(authData.Secret.Key, authData.Secret.Value) {
		userSelector.ID = request.Selector.ID
	}

	query, err := model.UserTable.QueryUsers(ctx, userSelector, &pagination, &sort)
	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	count, err := model.UserTable.CountUsers(ctx, userSelector)
	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	for _, user := range query {
		response.Data.Users = append(
			response.Data.Users,
			user.OutProtoStruct(),
		)
	}

	response.Data.Total = count
	response.State = protos.State_SUCCESS
	return
}

func (s *accountServer) DeleteUser(ctx context.Context, request *protos.DeleteUserRequest) (response *protos.DeleteUserResponse, err error) {
	response = &protos.DeleteUserResponse{}
	if pass := ResolvePermissionFormIncomeContext(ctx, "User", func(rule PermissionRule) bool {
		matchID := rule.ExactMatch(model.ActionDelete, "id", strconv.FormatInt(*request.Selectors.ID, 10))
		matchName := rule.ExactMatch(model.ActionDelete, "username", *request.Selectors.Username)
		return matchID || matchName
	}); !pass {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func (s *accountServer) CreateSecret(ctx context.Context, request *protos.CreateSecretRequest) (response *protos.CreateSecretResponse, err error) {
	response = &protos.CreateSecretResponse{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	err = model.SecretTable.CreateSecret(ctx, model.CreateSecretParams{
		UserID: authData.User.ID,
		Type:   model.UserSecretType,
	})

	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = protos.State_SUCCESS
	return
}

func (s *accountServer) DisableSecret(ctx context.Context, request *protos.DisableSecretRequest) (response *protos.DisableSecretResponse, err error) {
	response = &protos.DisableSecretResponse{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method DisableSecret not implemented")
}

func (s *accountServer) DeleteSecret(ctx context.Context, request *protos.DeleteSecreteRequest) (response *protos.DeleteSecreteResponse, err error) {
	response = &protos.DeleteSecreteResponse{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}

func (s *accountServer) QuerySecrets(ctx context.Context, request *protos.QuerySecretsRequest) (response *protos.QuerySecretsResponse, err error) {
	response = &protos.QuerySecretsResponse{}
	response.Data = &protos.QuerySecretsResponse_DataType{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	var sort model.Sort
	var pagination model.Pagination
	var selector model.SecretSelector
	sort.LoadProtoStruct(request.Sort)
	pagination.LoadProtoStruct(request.Pagination)
	selector.LoadProtoStruct(request.Selector)

	// 如果不是内部调用则将 UserID 范围设置为当前用户
	if !helper.IsGodSecret(authData.Secret.Key, authData.Secret.Value) {
		selector.UserID = &authData.User.ID
	}

	query, err := model.SecretTable.QuerySecrets(ctx, selector, &pagination, &sort)
	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	count, err := model.SecretTable.CountSecrets(ctx, selector)
	if err != nil {
		response.State = protos.State_FAILURE
		logger.Error(err)
		return
	}

	for _, secret := range query {
		response.Data.Secrets = append(
			response.Data.Secrets,
			secret.OutProtoStruct(),
		)
	}

	response.Data.Total = count
	response.State = protos.State_SUCCESS
	return
}

func (s *accountServer) CreateSetting(ctx context.Context, request *protos.CreateSettingRequest) (response *protos.CreateSettingResponse, err error) {
	response = &protos.CreateSettingResponse{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method CreateSetting not implemented")
}

func (s *accountServer) UpdateSetting(ctx context.Context, request *protos.UpdateSettingRequest) (response *protos.UpdateSettingResponse, err error) {
	response = &protos.UpdateSettingResponse{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method UpdateSetting not implemented")
}

func (s *accountServer) DeleteSetting(ctx context.Context, request *protos.DeleteSettingRequest) (response *protos.DeleteSettingResponse, err error) {
	response = &protos.DeleteSettingResponse{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}

	return nil, status.Errorf(codes.Unimplemented, "method DeleteSetting not implemented")
}

func (s *accountServer) QuerySettings(ctx context.Context, request *protos.QuerySettingsRequest) (response *protos.QuerySettingsResponse, err error) {
	response = &protos.QuerySettingsResponse{}

	authData := auth.ResolveAuthFromIncomingContext(ctx)
	if authData == nil || authData.User == nil {
		response.State = protos.State_NO_PERMISSION
		response.Message = "Not logged in"
		return
	}
	return nil, status.Errorf(codes.Unimplemented, "method QuerySettings not implemented")
}
