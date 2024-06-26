package server

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/code"
	"github.com/atom-service/account/package/proto"
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
		slog.ErrorContext(ctx, "QueryUsers failed", slog.Any("error", err))
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
		slog.ErrorContext(ctx, "QuerySecrets failed", err)
		return
	}

	if len(querySecretResult) <= 0 {
		response.Code = code.USER_SECRET_NOT_EXIST
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "No Secret available")
		return
	}

	firstSecret := querySecretResult[0]
	ExpiredTime := time.Now().Add(7 * 24 * time.Hour)
	token := auth.SignToken(*firstSecret.Key, *firstSecret.Value, auth.SignData{
		ExpiresAt: ExpiredTime,
	})

	go func() {
		var internalError error
		localContext, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// 异步更新一下 label 上的状态
		currentTime := time.Now().String()
		if internalError = model.LabelTable.UpsertLabel(localContext, model.Label{
			UserID: *queryResult[0].ID,
			Key:    model.LabelLastSignInTime,
			Value:  &currentTime,
		}); internalError != nil {
			slog.ErrorContext(localContext, "Update last sign in time failed", slog.Any("error", internalError))
		}

		if internalError = model.LabelTable.UpsertLabel(localContext, model.Label{
			UserID: *queryResult[0].ID,
			Key:    model.LabelLastVerifyTime,
			Value:  &currentTime,
		}); internalError != nil {
			slog.ErrorContext(localContext, "Update last sign in time failed", slog.Any("error", internalError))
		}
	}()

	response.Data = &proto.SignInResponse_DataType{
		ExpiredTime: ExpiredTime.UTC().String(),
		UserID:      *queryResult[0].ID,
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
		slog.ErrorContext(ctx, "", err)
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
		slog.ErrorContext(ctx, "", err)
		return
	}

	queryUserResult, err := model.UserTable.QueryUsers(ctx, model.UserSelector{Username: &request.Username}, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if len(queryUserResult) <= 0 {
		response.State = proto.State_FAILURE
		return
	}

	// 创建一组 system AK/SK
	err = model.SecretTable.CreateSecret(ctx, model.CreateSecretParams{
		UserID: *queryUserResult[0].ID,
		Type:   model.SystemSecretType,
	})

	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	// 赋予基本的 owner 权限
	roleSelector := model.RoleSelector{}
	roleSelector.Name = &model.OwnerRoleName
	roleQueryResult, err := model.RoleTable.QueryRoles(ctx, roleSelector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if len(roleQueryResult) <= 0 {
		response.State = proto.State_FAILURE
		response.Code = code.PERMISSION_ROLE_NOT_EXIST
		return
	}

	err = model.UserRoleTable.CreateUserRole(ctx, model.UserRole{
		UserID: *queryUserResult[0].ID,
		RoleID: *roleQueryResult[0].ID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	go func() {
		var internalError error
		localContext, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// 异步更新一下 label 上的状态
		currentTime := time.Now().String()
		if internalError = model.LabelTable.UpsertLabel(localContext, model.Label{
			UserID: *queryUserResult[0].ID,
			Value:  &currentTime,
			Key:    model.LabelLastVerifyTime,
		}); internalError != nil {
			slog.ErrorContext(localContext, "Update last sign in time failed", slog.Any("error", internalError))
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
		return permission.MatchRules(model.UserResourceName, model.ActionQuery, matchRules...)
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
		slog.ErrorContext(ctx, "", err)
		return
	}

	count, err := model.UserTable.CountUsers(ctx, userSelector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
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

		return permission.MatchRules(model.UserResourceName, model.ActionDelete, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	var userSelector model.UserSelector
	userSelector.LoadProto(request.Selector)

	err = model.UserTable.DeleteUser(ctx, userSelector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) CreateSecret(ctx context.Context, request *proto.CreateSecretRequest) (response *proto.CreateSecretResponse, err error) {
	response = &proto.CreateSecretResponse{}
	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}

		// 如果未指定，则默认为自己
		if request.UserID == nil {
			request.UserID = user.ID
		}

		// 检查用户是否操作的是自己
		if permission.HasOwner() && *request.UserID == *user.ID {
			return true
		}
		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "id",
			Value: strconv.FormatInt(*request.UserID, 10),
		})

		return permission.MatchRules(model.UserSecretResourceName, model.ActionInsert, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	count, err := model.SecretTable.CountSecrets(ctx, model.SecretSelector{
		UserID: request.UserID,
		Type:   &model.UserSecretType,
	})

	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if count > 30 {
		response.State = proto.State_FAILURE
		response.Code = code.TOO_MANY_SECRETS
		slog.ErrorContext(ctx, "", err)
		return
	}

	err = model.SecretTable.CreateSecret(ctx, model.CreateSecretParams{
		UserID: *request.UserID,
		Type:   model.UserSecretType,
	})

	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
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
		if request.Selector == nil {
			request.Selector = &proto.SecretSelector{
				UserID: user.ID,
			}
		}

		if request.Selector.UserID == nil {
			request.Selector.UserID = user.ID
		}

		if permission.HasOwner() && *request.Selector.UserID == *user.ID {
			return true
		}

		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.Selector.UserID, 10),
		})
		return permission.MatchRules(model.UserSecretResourceName, model.ActionUpdate, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SecretSelector{}
	selector.LoadProto(request.Selector)

	count, err := model.SecretTable.CountSecrets(ctx, selector)

	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if count == 0 {
		response.State = proto.State_FAILURE
		response.Code = code.USER_SECRET_NOT_EXIST
		slog.ErrorContext(ctx, "", err)
		return
	}

	disabledTime := time.Now()
	err = model.SecretTable.UpdateSecret(ctx, selector, &model.Secret{DisabledTime: &disabledTime})

	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) DeleteSecrets(ctx context.Context, request *proto.DeleteSecretsRequest) (response *proto.DeleteSecretsResponse, err error) {
	response = &proto.DeleteSecretsResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 检查用户是否操作的是自己
		if request.Selector == nil {
			request.Selector = &proto.SecretSelector{
				UserID: user.ID,
			}
		}

		if request.Selector.UserID == nil {
			request.Selector.UserID = user.ID
		}

		if permission.HasOwner() && *request.Selector.UserID == *user.ID {
			return true
		}

		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.Selector.UserID, 10),
		})
		return permission.MatchRules(model.UserSecretResourceName, model.ActionDelete, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SecretSelector{}
	selector.LoadProto(request.Selector)

	queryResult, err := model.SecretTable.QuerySecrets(ctx, selector, nil, nil)

	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if len(queryResult) == 0 {
		response.State = proto.State_FAILURE
		response.Code = code.USER_SECRET_NOT_EXIST
		slog.ErrorContext(ctx, "", err)
		return
	}

	if queryResult[0].DisabledTime == nil {
		response.State = proto.State_FAILURE
		response.Code = code.USER_SECRET_NOT_DISABLED
		return
	}

	if queryResult[0].DisabledTime != nil {
		if queryResult[0].DisabledTime.After(time.Now()) {
			response.State = proto.State_FAILURE
			response.Code = code.USER_SECRET_NOT_DISABLED
			return
		}
	}

	err = model.SecretTable.DeleteSecret(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
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
		// 不指定则操作当前账号
		if request.Selector == nil {
			request.Selector = &proto.SecretSelector{
				UserID: user.ID,
			}
		}

		if request.Selector.UserID == nil {
			request.Selector.UserID = user.ID
		}

		// 检查用户是否拥有 owner 权限并且正在操作自己的资源
		if permission.HasOwner() && *request.Selector.UserID == *user.ID {
			return true
		}

		// 检查是否有指定的其他 uid 的权限
		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.Selector.UserID, 10),
		})

		return permission.MatchRules(model.UserSecretResourceName, model.ActionQuery, matchRules...)
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
		slog.ErrorContext(ctx, "", err)
		return
	}

	count, err := model.SecretTable.CountSecrets(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
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

func (s *accountServer) UpsertLabel(ctx context.Context, request *proto.UpsertLabelRequest) (response *proto.UpsertLabelResponse, err error) {
	response = &proto.UpsertLabelResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 不指定则操作当前账号
		if request.UserID == nil {
			request.UserID = user.ID
		}

		if permission.HasOwner() && *request.UserID == *user.ID {
			return true
		}
		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.UserID, 10),
		})
		return permission.MatchRules(model.UserLabelResourceName, model.ActionInsert, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	err = model.LabelTable.UpsertLabel(ctx, model.Label{
		Key:    *request.Key,
		Value:  request.Value,
		UserID: *request.UserID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) QueryLabels(ctx context.Context, request *proto.QueryLabelsRequest) (response *proto.QueryLabelsResponse, err error) {
	response = &proto.QueryLabelsResponse{}
	response.Data = &proto.QueryLabelsResponse_DataType{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 不指定则操作当前账号
		if request.Selector == nil {
			request.Selector = &proto.LabelSelector{
				UserID: user.ID,
			}
		}

		if request.Selector.UserID == nil {
			request.Selector.UserID = user.ID
		}

		if permission.HasOwner() && *request.Selector.UserID == *user.ID {
			return true
		}

		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.Selector.UserID, 10),
		})
		return permission.MatchRules(model.UserSettingResourceName, model.ActionQuery, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	var sort model.Sort
	sort.LoadProto(request.Sort)
	var pagination model.Pagination
	pagination.LoadProto(request.Pagination)
	selector := model.LabelSelector{}
	selector.LoadProto(request.Selector)

	count, err := model.LabelTable.CountLabels(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	queryResult, err := model.LabelTable.QueryLabels(ctx, selector, &pagination, &sort)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	for _, label := range queryResult {
		response.Data.Labels = append(
			response.Data.Labels,
			label.ToProto(),
		)
	}

	response.Data.Total = count
	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) DeleteLabels(ctx context.Context, request *proto.DeleteLabelsRequest) (response *proto.DeleteLabelsResponse, err error) {
	response = &proto.DeleteLabelsResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 不指定则操作当前账号
		if request.Selector == nil {
			request.Selector = &proto.LabelSelector{
				UserID: user.ID,
			}
		}

		if request.Selector.UserID == nil {
			request.Selector.UserID = user.ID
		}

		if permission.HasOwner() && *request.Selector.UserID == *user.ID {
			return true
		}

		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.Selector.UserID, 10),
		})
		return permission.MatchRules(model.UserSettingResourceName, model.ActionQuery, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.LabelSelector{}
	selector.LoadProto(request.Selector)

	err = model.LabelTable.DeleteLabel(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) CreateSetting(ctx context.Context, request *proto.CreateSettingRequest) (response *proto.CreateSettingResponse, err error) {
	response = &proto.CreateSettingResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}

		if request.UserID == nil {
			request.UserID = user.ID
		}

		// 检查用户是否操作的是自己
		if permission.HasOwner() && *request.UserID == *user.ID {
			return true
		}

		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.UserID, 10),
		})
		return permission.MatchRules(model.UserSettingResourceName, model.ActionInsert, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SettingSelector{}
	selector.Key = &request.Key
	selector.UserID = request.UserID

	count, err := model.SettingTable.CountSettings(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
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
		UserID: *request.UserID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) UpdateSetting(ctx context.Context, request *proto.UpdateSettingRequest) (response *proto.UpdateSettingResponse, err error) {
	response = &proto.UpdateSettingResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 不指定则操作当前账号
		if request.Selector == nil {
			request.Selector = &proto.SettingSelector{
				UserID: user.ID,
			}
		}

		if request.Selector.UserID == nil {
			request.Selector.UserID = user.ID
		}

		if permission.HasOwner() && *request.Selector.UserID == *user.ID {
			return true
		}
		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.Selector.UserID, 10),
		})
		return permission.MatchRules(model.UserSettingResourceName, model.ActionInsert, matchRules...)
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
		slog.ErrorContext(ctx, "", err)
		return
	}

	if count <= 0 {
		response.State = proto.State_FAILURE
		response.Code = code.SETTING_NOT_EXIST
		return
	}

	err = model.SettingTable.UpdateSetting(ctx, selector, &model.Setting{
		Key:    *request.Data.Key,
		Value:  request.Data.Value,
		UserID: *request.Selector.UserID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *accountServer) DeleteSettings(ctx context.Context, request *proto.DeleteSettingsRequest) (response *proto.DeleteSettingsResponse, err error) {
	response = &proto.DeleteSettingsResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		matchRules := []model.UserResourcePermissionRule{}
		// 不指定则操作当前账号
		if request.Selector == nil {
			request.Selector = &proto.SettingSelector{
				UserID: user.ID,
			}
		}
		if request.Selector.UserID == nil {
			request.Selector.UserID = user.ID
		}

		if permission.HasOwner() && *request.Selector.UserID == *user.ID {
			return true
		}
		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.Selector.UserID, 10),
		})
		return permission.MatchRules(model.UserSettingResourceName, model.ActionDelete, matchRules...)
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.SettingSelector{}
	selector.LoadProto(request.Selector)

	err = model.SettingTable.DeleteSetting(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
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
		// 不指定则操作当前账号
		if request.Selector == nil {
			request.Selector = &proto.SettingSelector{
				UserID: user.ID,
			}
		}

		if request.Selector.UserID == nil {
			request.Selector.UserID = user.ID
		}

		if permission.HasOwner() && *request.Selector.UserID == *user.ID {
			return true
		}

		matchRules = append(matchRules, model.UserResourcePermissionRule{
			Key:   "user_id",
			Value: strconv.FormatInt(*request.Selector.UserID, 10),
		})
		return permission.MatchRules(model.UserSettingResourceName, model.ActionQuery, matchRules...)
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
		slog.ErrorContext(ctx, "", err)
		return
	}

	queryResult, err := model.SettingTable.QuerySettings(ctx, selector, &pagination, &sort)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
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
