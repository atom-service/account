package server

import (
	"context"
	"log/slog"
	"sync"

	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/code"
	"github.com/atom-service/account/package/proto"
)

var PermissionServer = &permissionServer{}

type permissionServer struct {
	proto.UnimplementedPermissionServiceServer
}

func (s *permissionServer) CreateRole(ctx context.Context, request *proto.CreateRoleRequest) (response *proto.CreateRoleResponse, err error) {
	response = &proto.CreateRoleResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission.role" && permission.Action == model.ActionInsert
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.RoleSelector{}
	selector.Name = &request.Name
	countResult, err := model.RoleTable.CountRoles(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if countResult > 0 {
		response.State = proto.State_FAILURE
		response.Code = code.PERMISSION_ROLE_ALREADY_EXISTS
		return
	}

	err = model.RoleTable.CreateRole(ctx, model.Role{
		Name:        &request.Name,
		Description: &request.Description,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	// 简单插入 role 基本信息
	createdRoleResult, err := model.RoleTable.QueryRoles(ctx, selector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if len(createdRoleResult) <= 0 {
		response.State = proto.State_FAILURE
		return
	}

	response.Data = createdRoleResult[0].ToProto()

	if len(request.Resources) <= 0 {
		response.State = proto.State_SUCCESS
		return
	}

	updateRoleRequest := &proto.UpdateRoleRequest{}
	updateRoleRequest.Selector = &proto.RoleSelector{
		ID:   createdRoleResult[0].ID,
		Name: createdRoleResult[0].Name,
	}

	updateRoleRequest.Data = &proto.UpdateRoleRequest_UpdateData{
		Resources: request.Resources,
	}

	// 通过 UpdateRole 来完善所有信息
	updateResult, err := s.UpdateRole(ctx, updateRoleRequest)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if updateResult.State != proto.State_SUCCESS {
		response.State = updateResult.State
		response.Code = updateResult.Code
		return
	}

	// 通过 QueryRoles 再查询完整信息
	roleResponse, err := s.QueryRoles(ctx, &proto.QueryRolesRequest{
		Selector: &proto.RoleSelector{ID: createdRoleResult[0].ID},
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if roleResponse.State != proto.State_SUCCESS {
		response.State = roleResponse.State
		response.Code = roleResponse.Code
		return
	}

	response.State = proto.State_SUCCESS
	response.Data = roleResponse.Data.Roles[0]
	return
}

func (s *permissionServer) QueryRoles(ctx context.Context, request *proto.QueryRolesRequest) (response *proto.QueryRolesResponse, err error) {
	response = &proto.QueryRolesResponse{}
	response.Data = &proto.QueryRolesResponse_DataType{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission.role" && permission.Action == model.ActionQuery
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	var sort model.Sort
	sort.LoadProto(request.Sort)
	var pagination model.Pagination
	pagination.LoadProto(request.Pagination)
	var selector model.RoleSelector
	selector.LoadProto(request.Selector)

	roles, err := model.RoleTable.QueryRoles(ctx, selector, &pagination, &sort)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	count, err := model.RoleTable.CountRoles(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	// 循环查询 role 绑定的 resource
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, role := range roles {
		wg.Add(1)
		go func(role *model.Role) {
			defer wg.Done()
			roleResourceResponse, localErr := model.RoleResourceTable.QueryRoleResources(ctx, model.RoleResourceSelector{
				RoleID: role.ID,
			}, nil, nil)
			if localErr != nil {
				mu.Lock()
				err = localErr
				mu.Unlock()
				slog.ErrorContext(ctx, "", err)
				return
			}

			roleModel := role.ToProto()
			for _, item := range roleResourceResponse {
				resourceResponse, localErr := model.ResourceTable.QueryResources(ctx, model.ResourceSelector{
					ID: &item.ResourceID,
				}, nil, nil)
				if localErr != nil {
					mu.Lock()
					err = localErr
					mu.Unlock()
					slog.ErrorContext(ctx, "", err)
					return
				}

				roleResource := &proto.RoleResource{
					ResourceID:   item.ResourceID,
					ResourceName: *resourceResponse[0].Name,
					Action:       model.ActionToProto(item.Action),
					Rules:        []*proto.RoleResourceRule{},
				}

				for _, rule := range item.Rules {
					roleResource.Rules = append(roleResource.Rules, &proto.RoleResourceRule{
						Key:   rule.Key,
						Value: rule.Value,
					})
				}
				roleModel.Resources = append(roleModel.Resources, roleResource)
			}

			mu.Lock()
			response.Data.Roles = append(
				response.Data.Roles,
				roleModel,
			)
			mu.Unlock()
		}(role)
	}

	wg.Wait()
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.Data.Total = count
	response.State = proto.State_SUCCESS
	return
}

func (s *permissionServer) UpdateRole(ctx context.Context, request *proto.UpdateRoleRequest) (response *proto.UpdateRoleResponse, err error) {
	response = &proto.UpdateRoleResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission.role" && permission.Action == model.ActionUpdate
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.RoleSelector{}
	selector.LoadProto(request.Selector)

	queryRoleResult, err := model.RoleTable.QueryRoles(ctx, selector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if len(queryRoleResult) == 0 {
		response.State = proto.State_FAILURE
		response.Code = code.PERMISSION_ROLE_NOT_EXIST
		slog.ErrorContext(ctx, "", err)
		return
	}

	if request.Data.Name != nil || request.Data.Description != nil {
		err = model.RoleTable.UpdateRole(ctx, selector, &model.Role{
			Name:        request.Data.Name,
			Description: request.Data.Description,
		})

		if err != nil {
			response.State = proto.State_FAILURE
			slog.ErrorContext(ctx, "", err)
			return
		}
	}

	if len(request.Data.Resources) <= 0 {
		response.State = proto.State_SUCCESS
		return
	}

	// 查询数据库中关于当前 role 的全部 role_resources
	roleResourceSelector := model.RoleResourceSelector{}
	roleResourceSelector.RoleID = queryRoleResult[0].ID
	queryRoleResourceResult, err := model.RoleResourceTable.QueryRoleResources(ctx, roleResourceSelector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	// 先无脑删除
	for _, roleResource := range queryRoleResourceResult {
		err = model.RoleResourceTable.DeleteRoleResource(ctx, model.RoleResourceSelector{
			Action:     &roleResource.Action,
			RoleID:     &roleResource.RoleID,
			ResourceID: &roleResource.ResourceID,
		})
		if err != nil {
			response.State = proto.State_FAILURE
			slog.ErrorContext(ctx, "", err)
			return
		}
	}

	// 再无脑创建,这里数据正确比效率更重要
	for _, resource := range request.Data.Resources {
		selector := model.RoleResourceSelector{}
		selector.LoadProtoAction(resource.Action)
		selector.RoleID = queryRoleResult[0].ID
		selector.ResourceID = &resource.ResourceID

		rules := []*model.RoleResourceRule{}
		for _, ruleProto := range resource.Rules {
			rules = append(rules, &model.RoleResourceRule{
				Key:   ruleProto.Key,
				Value: ruleProto.Value,
			})
		}

		err = model.RoleResourceTable.CreateRoleResource(ctx, model.RoleResource{
			RoleID:     *queryRoleResult[0].ID,
			ResourceID: resource.ResourceID,
			Action:     *selector.Action,
			Rules:      rules,
		})
		if err != nil {
			response.State = proto.State_FAILURE
			slog.ErrorContext(ctx, "", err)
			return
		}
	}

	// 通过 QueryRoles 再查询完整信息
	roleResponse, err := s.QueryRoles(ctx, &proto.QueryRolesRequest{
		Selector: &proto.RoleSelector{ID: queryRoleResult[0].ID},
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if roleResponse.State != proto.State_SUCCESS {
		response.State = roleResponse.State
		response.Code = roleResponse.Code
		return
	}

	response.State = proto.State_SUCCESS
	response.Data = roleResponse.Data.Roles[0]
	return
}

func (s *permissionServer) DeleteRole(ctx context.Context, request *proto.DeleteRoleRequest) (response *proto.DeleteRoleResponse, err error) {
	response = &proto.DeleteRoleResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission.role" && permission.Action == model.ActionDelete
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	roleSelector := model.RoleSelector{}
	roleSelector.LoadProto(request.Selector)
	queryRoleResult, err := model.RoleTable.QueryRoles(ctx, roleSelector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	// 删除与之相关的所有 RoleResource 以及 Rules
	roleResourceSelector := model.RoleResourceSelector{}
	roleResourceSelector.RoleID = queryRoleResult[0].ID
	err = model.RoleResourceTable.DeleteRoleResource(ctx, roleResourceSelector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	err = model.RoleTable.DeleteRole(ctx, roleSelector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *permissionServer) CreateResource(ctx context.Context, request *proto.CreateResourceRequest) (response *proto.CreateResourceResponse, err error) {
	response = &proto.CreateResourceResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission.resource" && permission.Action == model.ActionInsert
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.ResourceSelector{}
	selector.Name = &request.Name
	countResult, err := model.ResourceTable.CountResources(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if countResult > 0 {
		response.State = proto.State_FAILURE
		response.Code = code.PERMISSION_RESOURCE_ALREADY_EXISTS
		return
	}

	err = model.ResourceTable.CreateResource(ctx, model.Resource{
		Name:        &request.Name,
		Description: &request.Description,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	queryResult, err := model.ResourceTable.QueryResources(ctx, selector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if len(queryResult) <= 0 {
		response.State = proto.State_FAILURE
		return
	}

	response.Data = queryResult[0].ToProto()
	response.State = proto.State_SUCCESS
	return
}

func (s *permissionServer) QueryResources(ctx context.Context, request *proto.QueryResourcesRequest) (response *proto.QueryResourcesResponse, err error) {
	response = &proto.QueryResourcesResponse{}
	response.Data = &proto.QueryResourcesResponse_DataType{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission.resource" && permission.Action == model.ActionQuery
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	var sort model.Sort
	sort.LoadProto(request.Sort)
	var pagination model.Pagination
	pagination.LoadProto(request.Pagination)
	var selector model.ResourceSelector
	selector.LoadProto(request.Selector)

	query, err := model.ResourceTable.QueryResources(ctx, selector, &pagination, &sort)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	count, err := model.ResourceTable.CountResources(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	for _, resources := range query {
		response.Data.Resources = append(
			response.Data.Resources,
			resources.ToProto(),
		)
	}

	response.Data.Total = count
	response.State = proto.State_SUCCESS
	return
}

func (s *permissionServer) DeleteResources(ctx context.Context, request *proto.DeleteResourcesRequest) (response *proto.DeleteResourcesResponse, err error) {
	response = &proto.DeleteResourcesResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission.resource" && permission.Action == model.ActionDelete
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.ResourceSelector{}
	selector.LoadProto(request.Selector)

	queryResult, err := model.ResourceTable.QueryResources(ctx, selector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	resourceRule, err := model.RoleResourceTable.QueryRoleResources(ctx, model.RoleResourceSelector{ResourceID: queryResult[0].ID}, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if len(resourceRule) > 0 {
		response.State = proto.State_FAILURE
		response.Code = code.PERMISSION_RESOURCE_IN_USE
		return
	}

	err = model.ResourceTable.DeleteResource(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *permissionServer) UpdateResource(ctx context.Context, request *proto.UpdateResourceRequest) (response *proto.UpdateResourceResponse, err error) {
	response = &proto.UpdateResourceResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission.resource" && permission.Action == model.ActionInsert
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	selector := model.ResourceSelector{}
	selector.LoadProto(request.Selector)

	countResult, err := model.ResourceTable.CountResources(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if countResult <= 0 {
		response.State = proto.State_FAILURE
		response.Code = code.PERMISSION_RESOURCE_NOT_EXIST
		return
	}

	err = model.ResourceTable.UpdateResource(ctx, selector, &model.Resource{
		Name:        request.Data.Name,
		Description: request.Data.Description,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *permissionServer) SummaryForUser(ctx context.Context, request *proto.SummaryForUserRequest) (response *proto.SummaryForUserResponse, err error) {
	response = &proto.SummaryForUserResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission" && permission.Action == model.ActionInsert
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	// 这个方法调用率极高, 所以应该添加缓存
	selector := model.UserResourceSummarySelector{UserID: request.UserSelector.ID}
	queryResult, err := model.Permission.QueryUserResourceSummaries(ctx, selector)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	for _, resources := range queryResult {
		response.Data = append(
			response.Data,
			resources.ToProto(),
		)
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *permissionServer) ApplyRoleForUser(ctx context.Context, request *proto.ApplyRoleForUserRequest) (response *proto.ApplyRoleForUserResponse, err error) {
	response = &proto.ApplyRoleForUserResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission" && permission.Action == model.ActionUpdate
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	userSelector := model.UserSelector{}
	userSelector.LoadProto(request.User)

	userQueryResult, err := model.UserTable.QueryUsers(ctx, userSelector, nil, nil)
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if len(userQueryResult) <= 0 {
		response.State = proto.State_FAILURE
		response.Code = code.USER_NOT_EXIST
		return
	}

	roleSelector := model.RoleSelector{}
	roleSelector.LoadProto(request.Role)

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

	userRoleCount, err := model.UserRoleTable.CountUserRoles(ctx, model.UserRoleSelector{
		UserID: userQueryResult[0].ID,
		RoleID: roleQueryResult[0].ID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	// 不管什么原因,既然存在了,就视为应用成功
	if userRoleCount > 0 {
		response.State = proto.State_SUCCESS
		return
	}

	err = model.UserRoleTable.CreateUserRole(ctx, model.UserRole{
		UserID: *userQueryResult[0].ID,
		RoleID: *request.Role.ID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}

func (s *permissionServer) RemoveRoleForUser(ctx context.Context, request *proto.RemoveRoleForUserRequest) (response *proto.RemoveRoleForUserResponse, err error) {
	response = &proto.RemoveRoleForUserResponse{}

	if pass := auth.ResolvePermission(ctx, func(user *model.User, permission *model.UserResourcePermissionSummary) bool {
		return permission.ResourceName == "permission" && permission.Action == model.ActionUpdate
	}); !pass {
		response.State = proto.State_NO_PERMISSION
		return
	}

	countResult, err := model.UserRoleTable.CountUserRoles(ctx, model.UserRoleSelector{
		UserID: request.User.ID,
		RoleID: request.Role.ID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	if countResult <= 0 {
		response.State = proto.State_SUCCESS
		return
	}

	err = model.UserRoleTable.DeleteUserRole(ctx, model.UserRoleSelector{
		UserID: request.User.ID,
		RoleID: request.Role.ID,
	})
	if err != nil {
		response.State = proto.State_FAILURE
		slog.ErrorContext(ctx, "", err)
		return
	}

	response.State = proto.State_SUCCESS
	return
}
