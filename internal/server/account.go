package server

import (
	"context"
	"strconv"
	"time"

	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/code"
	"github.com/atom-service/account/package/proto"
	"github.com/atom-service/common/logger"
)

var AccountServer = &accountServer{}

type accountServer struct {
	proto.UnimplementedAccountServiceServer
}

func (s *accountServer) SignIn(ctx context.Context, request *proto.SignInRequest) (response *proto.SignInResponse, err error) {
	response = &proto.SignInResponse{}

	var userSelector model.UserSelector
	userSelector.LoadProto(request.Selector)
	passwordHash := model.Password.Hash(request.Password)
	queryResult, err := model.UserTable.QueryUsers(ctx, userSelector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if len(queryResult) <= 0 {
		response.State = proto.State_PARAMS_INVALID
		return
	}

	// 用户已被禁用
	if queryResult[0].DisabledTime != nil && queryResult[0].DisabledTime.Before(time.Now()) {
		response.State = proto.State_ILLEGAL_REQUEST
		response.Code = code.USER_ALREADY_DISABLED
		return
	}

	// 验证密码是否正确
	if *queryResult[0].Password != passwordHash {
		response.State = proto.State_PARAMS_INVALID
		return
	}

	querySecretResult, err := model.SecretTable.QuerySecrets(ctx, model.SecretSelector{
		UserID: queryResult[0].ID,
		Type:   &model.SystemSecretType,
	}, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if len(querySecretResult) <= 0 {
		response.Code = code.USER_SECRET_NOT_EXIST
		response.State = proto.State_FAILURE
		logger.Errorf("No Secret available")
		return
	}

	firstSecret := querySecretResult[0]
	ExpiredTime := time.Now().Add(7 * 24 * time.Hour)
	token := auth.SignToken(*firstSecret.Key, *firstSecret.Value, auth.SignData{
		ExpiresAt: ExpiredTime,
	})

	go func() {
		// 异步更新一下 label 上的状态
		currentTime := time.Now().String()
		if err = model.LabelTable.UpsertLabel(ctx, model.Label{
			UserID: *queryResult[0].ID,
			Key:    model.LabelLastSignInTime,
			Value:  &currentTime,
		}); err != nil {
			logger.Errorf("Update last sign in time failed: %s", err)
		}

		if err = model.LabelTable.UpsertLabel(ctx, model.Label{
			UserID: *queryResult[0].ID,
			Key:    model.LabelLastVerifyTime,
			Value:  &currentTime,
		}); err != nil {
			logger.Errorf("Update last sign in time failed: %s", err)
		}
	}()

	response.Token = &proto.SignedInToken{
		ExpiredTime: ExpiredTime.String(),
		Token:       token,
	}
	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) SignUp(ctx context.Context, request *proto.SignUpRequest) (response *proto.SignUpResponse, err error) {
	response = &proto.SignUpResponse{}

	countResult, err := model.UserTable.CountUsers(ctx, model.UserSelector{
		Username: &request.Username,
	})

	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if countResult > 0 {
		response.State = proto.State_FAILURE
		response.Code = code.USER_ALREADY_EXISTS
		return
	}

	passwordHash := model.Password.Hash(request.Password)
	err = model.UserTable.CreateUser(ctx, model.User{
		Username: &request.Username,
		Password: &passwordHash,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	users, err := model.UserTable.QueryUsers(ctx, model.UserSelector{Username: &request.Username}, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if len(users) <= 0 {
		response.State = proto.State_FAILURE
		return
	}

	// 创建一组 system AK/SK
	err = model.SecretTable.CreateSecret(ctx, model.CreateSecretParams{
		UserID: *users[0].ID,
		Type:   model.SystemSecretType,
	})

	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	go func() {
		// 异步更新一下 label 上的状态
		currentTime := time.Now().String()
		if err = model.LabelTable.UpsertLabel(ctx, model.Label{
			UserID: *users[0].ID,
			Key:    model.LabelLastVerifyTime,
			Value:  &currentTime,
		}); err != nil {
			logger.Errorf("Update last sign in time failed: %s", err)
		}
	}()

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) SignOut(ctx context.Context, request *proto.SignOutRequest) (response *proto.SignOutResponse, err error) {
	response = &proto.SignOutResponse{}

	authData := auth.ResolveAuth(ctx)
	if authData == nil || authData.User == nil {
		response.State = proto.State_NO_PERMISSION
		return
	}

	return nil, nil
}

func (s *accountServer) QueryUsers(ctx context.Context, request *proto.QueryUsersRequest) (response *proto.QueryUsersResponse, err error) {
	response = &proto.QueryUsersResponse{}
	response.Data = &proto.QueryUsersResponse_DataType{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		if request.Selector != nil && request.Selector.ID != nil {
			if permission.HasOwner() && *request.Selector.ID == *user.ID {
				return true
			}

			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "id",
				Value: strconv.FormatInt(*request.Selector.ID, 10),
			})
		}

		if request.Selector != nil && request.Selector.Username != nil {
			if permission.HasOwner() && *request.Selector.Username == *user.Username {
				return true
			}
		}
		return permission.MatchRules("user", model.ActionQuery, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	var sort model.Sort
	sort.LoadProto(request.Sort)
	var pagination model.Pagination
	pagination.LoadProto(request.Pagination)
	var userSelector model.UserSelector
	userSelector.LoadProto(request.Selector)

	query, err := model.UserTable.QueryUsers(ctx, userSelector, &pagination, &sort)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	count, err := model.UserTable.CountUsers(ctx, userSelector)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	for _, user := range query {
		response.Data.Users = append(
			response.Data.Users,
			user.ToProto(),
		)
	}

	response.Data.Total = count
	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) DeleteUser(ctx context.Context, request *proto.DeleteUserRequest) (response *proto.DeleteUserResponse, err error) {
	response = &proto.DeleteUserResponse{}
	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		if request.Selector != nil && request.Selector.ID != nil {
			if permission.HasOwner() && *request.Selector.ID == *user.ID {
				return true
			}

			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "id",
				Value: strconv.FormatInt(*request.Selector.ID, 10),
			})
		}

		if request.Selector != nil && request.Selector.Username != nil {
			if permission.HasOwner() && *request.Selector.Username == *user.Username {
				return true
			}
		}

		return permission.MatchRules("user", model.ActionDelete, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	var userSelector model.UserSelector
	userSelector.LoadProto(request.Selector)

	err = model.UserTable.DeleteUser(ctx, userSelector)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) CreateSecret(ctx context.Context, request *proto.CreateSecretRequest) (response *proto.CreateSecretResponse, err error) {
	response = &proto.CreateSecretResponse{}
	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request.Selector != nil && request.Selector.ID != nil {
			if permission.HasOwner() && *request.Selector.ID == *user.ID {
				return true
			}
			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "id",
				Value: strconv.FormatInt(*request.Selector.ID, 10),
			})
		}

		if request.Selector != nil && request.Selector.Username != nil {
			if permission.HasOwner() && *request.Selector.Username == *user.Username {
				return true
			}
		}
		return permission.MatchRules("secret", model.ActionInsert, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	count, err := model.SecretTable.CountSecrets(ctx, model.SecretSelector{
		UserID: request.Selector.ID,
		Type:   &model.UserSecretType,
	})

	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if count > 30 {
		response.State = proto.State_FAILURE
		response.Code = code.TOO_MANY_SECRETS
		logger.Error(err)
		return
	}

	err = model.SecretTable.CreateSecret(ctx, model.CreateSecretParams{
		UserID: *request.Selector.ID,
		Type:   model.UserSecretType,
	})

	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) DisableSecret(ctx context.Context, request *proto.DisableSecretRequest) (response *proto.DisableSecretResponse, err error) {
	response = &proto.DisableSecretResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request.Selector != nil && request.Selector.UserID != nil {
			if permission.HasOwner() && *request.Selector.UserID == *user.ID {
				return true
			}
			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "user_id",
				Value: strconv.FormatInt(*request.Selector.UserID, 10),
			})
		}
		return permission.MatchRules("secret", model.ActionUpdate, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SecretSelector{}
	selector.LoadProto(request.Selector)

	count, err := model.SecretTable.CountSecrets(ctx, selector)

	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if count == 0 {
		response.State = proto.State_FAILURE
		response.Code = code.USER_SECRET_NOT_EXIST
		logger.Error(err)
		return
	}

	disabledTime := time.Now()
	err = model.SecretTable.UpdateSecret(ctx, selector, &model.Secret{DeletedTime: &disabledTime})

	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) DeleteSecret(ctx context.Context, request *proto.DeleteSecreteRequest) (response *proto.DeleteSecreteResponse, err error) {
	response = &proto.DeleteSecreteResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request.Selector != nil && request.Selector.UserID != nil {
			if permission.HasOwner() && *request.Selector.UserID == *user.ID {
				return true
			}
			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "user_id",
				Value: strconv.FormatInt(*request.Selector.UserID, 10),
			})
		}
		return permission.MatchRules("secret", model.ActionDelete, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SecretSelector{}
	selector.LoadProto(request.Selector)

	queryResult, err := model.SecretTable.QuerySecrets(ctx, selector, nil, nil)

	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if len(queryResult) == 0 {
		response.State = proto.State_FAILURE
		response.Code = code.USER_SECRET_NOT_EXIST
		logger.Error(err)
		return
	}

	if queryResult[0].DisabledTime != nil {
		if queryResult[0].DisabledTime.Before(time.Now()) {
			response.State = proto.State_FAILURE
			response.Code = code.USER_SECRET_NOT_DISABLED
			logger.Error(err)
		}
	}

	err = model.SecretTable.DeleteSecret(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) QuerySecrets(ctx context.Context, request *proto.QuerySecretsRequest) (response *proto.QuerySecretsResponse, err error) {
	response = &proto.QuerySecretsResponse{}
	response.Data = &proto.QuerySecretsResponse_DataType{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request.Selector != nil && request.Selector.UserID != nil {
			if permission.HasOwner() && *request.Selector.UserID == *user.ID {
				return true
			}
			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "user_id",
				Value: strconv.FormatInt(*request.Selector.UserID, 10),
			})
		}
		return permission.MatchRules("secret", model.ActionQuery, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	var sort model.Sort
	var pagination model.Pagination
	var selector model.SecretSelector
	sort.LoadProto(request.Sort)
	pagination.LoadProto(request.Pagination)
	selector.LoadProto(request.Selector)

	query, err := model.SecretTable.QuerySecrets(ctx, selector, &pagination, &sort)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	count, err := model.SecretTable.CountSecrets(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	for _, secret := range query {
		response.Data.Secrets = append(
			response.Data.Secrets,
			secret.ToProto(),
		)
	}

	response.Data.Total = count
	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) CreateSetting(ctx context.Context, request *proto.CreateSettingRequest) (response *proto.CreateSettingResponse, err error) {
	response = &proto.CreateSettingResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request != nil {
			if permission.HasOwner() && request.UserID == *user.ID {
				return true
			}
			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "user_id",
				Value: strconv.FormatInt(request.UserID, 10),
			})
		}
		return permission.MatchRules("setting", model.ActionInsert, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SettingSelector{}
	selector.Key = &request.Key
	selector.UserID = &request.UserID

	count, err := model.SettingTable.CountSettings(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if count > 0 {
		response.State = proto.State_FAILURE
		response.Code = code.SETTING_ALREADY_EXISTS
		return
	}

	err = model.SettingTable.CreateSetting(ctx, model.Setting{
		Key:    request.Key,
		Value:  &request.Value,
		UserID: request.UserID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) UpdateSetting(ctx context.Context, request *proto.UpdateSettingRequest) (response *proto.UpdateSettingResponse, err error) {
	response = &proto.UpdateSettingResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request != nil && request.Selector.UserID != nil {
			if permission.HasOwner() && *request.Selector.UserID == *user.ID {
				return true
			}
			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "user_id",
				Value: strconv.FormatInt(*request.Selector.UserID, 10),
			})
		}
		return permission.MatchRules("setting", model.ActionInsert, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SettingSelector{}
	selector.Key = request.Selector.Key
	selector.UserID = request.Selector.UserID

	count, err := model.SettingTable.CountSettings(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	if count > 0 {
		response.State = proto.State_FAILURE
		response.Code = code.SETTING_NOT_EXIST
		return
	}

	err = model.SettingTable.CreateSetting(ctx, model.Setting{
		Value:  request.Data.Value,
		Key:    *request.Selector.Key,
		UserID: *request.Selector.UserID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) DeleteSetting(ctx context.Context, request *proto.DeleteSettingRequest) (response *proto.DeleteSettingResponse, err error) {
	response = &proto.DeleteSettingResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request != nil && request.Selector.UserID != nil {
			if permission.HasOwner() && *request.Selector.UserID == *user.ID {
				return true
			}
			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "user_id",
				Value: strconv.FormatInt(*request.Selector.UserID, 10),
			})
		}
		return permission.MatchRules("setting", model.ActionDelete, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SettingSelector{}
	selector.LoadProto(request.Selector)

	err = model.SettingTable.DeleteSetting(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) QuerySettings(ctx context.Context, request *proto.QuerySettingsRequest) (response *proto.QuerySettingsResponse, err error) {
	response = &proto.QuerySettingsResponse{}
	response.Data = &proto.QuerySettingsResponse_DataType{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request != nil && request.Selector.UserID != nil {
			if permission.HasOwner() && *request.Selector.UserID == *user.ID {
				return true
			}
			matchRules = append(matchRules, model.UserResourcePermissionRule{
				Key:   "user_id",
				Value: strconv.FormatInt(*request.Selector.UserID, 10),
			})
		}
		return permission.MatchRules("setting", model.ActionQuery, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	var sort model.Sort
	sort.LoadProto(request.Sort)
	var pagination model.Pagination
	pagination.LoadProto(request.Pagination)
	selector := model.SettingSelector{}
	selector.LoadProto(request.Selector)

	count, err := model.SettingTable.CountSettings(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	queryResult, err := model.SettingTable.QuerySettings(ctx, selector, &pagination, &sort)
	if err != nil {
		response.State = proto.State_FAILURE
		logger.Error(err)
		return
	}

	for _, setting := range queryResult {
		response.Data.Settings = append(
			response.Data.Settings,
			setting.ToProto(),
		)
	}

	response.Data.Total = count
	response.State = proto.State_SUCCESS
	return
}
