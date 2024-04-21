package server

import (
	"context"
	"slices"
	"testing"
	"testing/quick"
	"time"

	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/proto"
)

func TestPermissionServer(t *testing.T) {
	context := context.TODO()
	testServer := createTestServer()

	rootAdminUserID := int64(1) // admin user 是 model 在 init 初始化好的
	adminSecretSelector := model.SecretSelector{UserID: &rootAdminUserID}
	adminSecret, err := model.SecretTable.QuerySecrets(context, adminSecretSelector, nil, nil)
	if err != nil {
		t.Errorf("Query admin secret failed: %v", err)
		return
	}
	if len(adminSecret) == 0 {
		t.Errorf("Query admin secret failed: %v", err)
		return
	}

	token := auth.SignToken(*adminSecret[0].Key, *adminSecret[0].Value, auth.SignData{
		ExpiresAt: time.Now().Add(time.Hour),
	})

	permissionClient := testServer.CreatePermissionClientWithToken(token)

	config := &quick.Config{
		MaxCount: 100,
	}

	resourceList := []*proto.Resource{}

	// create resource
	if err := quick.Check(func() bool {
		name := helper.GenerateRandomString(64, nil)
		description := helper.GenerateRandomString(64, nil)
		createResponse, err := permissionClient.CreateResource(context, &proto.CreateResourceRequest{
			Name:        name,
			Description: description,
		})
		if err != nil {
			t.Errorf("CreateResource failed: %v", err)
			return false
		}
		if createResponse.State != proto.State_SUCCESS {
			t.Errorf("CreateResource failed: %v", err)
			return false
		}

		selector := proto.ResourceSelector{Name: &name}
		queryResponse, err := permissionClient.QueryResources(context, &proto.QueryResourcesRequest{
			Selector: &selector,
		})
		if err != nil {
			t.Errorf("Unexpected results after created: %v", err)
			return false
		}
		if queryResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after created")
			return false
		}
		if queryResponse.Data.Total != 1 {
			t.Errorf("Unexpected results after created")
			return false
		}

		if queryResponse.Data.Resources[0].Name != name {
			t.Errorf("Unexpected results after created")
			return false
		}

		newName := helper.GenerateRandomString(64, nil)
		newDescription := helper.GenerateRandomString(64, nil)
		updateResponse, err := permissionClient.UpdateResource(context, &proto.UpdateResourceRequest{
			Selector: &selector,
			Data: &proto.UpdateResourceRequest_UpdateData{
				Name:        &newName,
				Description: &newDescription,
			},
		})
		if err != nil {
			t.Errorf("Unexpected results on update: %v", err)
			return false
		}
		if updateResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results on update")
			return false
		}

		queryUpdatedResponse, err := permissionClient.QueryResources(context, &proto.QueryResourcesRequest{
			Selector: &proto.ResourceSelector{Name: &newName},
		})
		if err != nil {
			t.Errorf("Unexpected results after updated: %v", err)
			return false
		}
		if queryUpdatedResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after updated")
			return false
		}
		if queryUpdatedResponse.Data.Total != 1 {
			t.Errorf("Unexpected results after updated")
			return false
		}

		if queryUpdatedResponse.Data.Resources[0].Name != newName {
			t.Errorf("Unexpected results after updated")
			return false
		}

		resourceList = append(resourceList, queryResponse.Data.Resources[0])
		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}

	roleList := []*proto.Role{}
	// create role
	if err := quick.Check(func() bool {
		name := helper.GenerateRandomString(64, nil)
		description := helper.GenerateRandomString(64, nil)
		createResponse, err := permissionClient.CreateRole(context, &proto.CreateRoleRequest{
			Name:        name,
			Description: description,
			Resources: []*proto.RoleResource{
				{
					ResourceID: resourceList[0].ID,
					ResourceName: resourceList[0].Name,
					Action:     proto.ResourceAction_Insert,
					Rules: []*proto.RoleResourceRule{
						{
							Key:   "TEST_CREATE",
							Value: "TEST_CREATE",
						},
					},
				},
			},
		})
		if err != nil {
			t.Errorf("CreateRole failed: %v", err)
			return false
		}
		if createResponse.State != proto.State_SUCCESS {
			t.Errorf("CreateRole failed: %v", err)
			return false
		}

		selector := proto.RoleSelector{Name: &name}
		queryResponse, err := permissionClient.QueryRoles(context, &proto.QueryRolesRequest{
			Selector: &selector,
		})
		if err != nil {
			t.Errorf("Unexpected results after created: %v", err)
			return false
		}
		if queryResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after created")
			return false
		}
		if queryResponse.Data.Total != 1 {
			t.Errorf("Unexpected results after created")
			return false
		}

		if queryResponse.Data.Roles[0].Name != name {
			t.Errorf("Unexpected results after created")
			return false
		}

		if (queryResponse.Data.Roles[0].Resources == nil) {
			t.Errorf("Unexpected results after created")
			return false
		}

		if (queryResponse.Data.Roles[0].Resources[0] == nil) {
			t.Errorf("Unexpected results after created")
			return false
		}

		if (len(queryResponse.Data.Roles[0].Resources[0].Rules) == 0) {
			t.Errorf("Unexpected results after created")
			return false
		}

		if (queryResponse.Data.Roles[0].Resources[0].Rules[0].Key == "TEST_CREATE") {
			t.Errorf("Unexpected results after created")
			return false
		}

		if (queryResponse.Data.Roles[0].Resources[0].Rules[0].Value == "TEST_CREATE") {
			t.Errorf("Unexpected results after created")
			return false
		}

		newName := helper.GenerateRandomString(64, nil)
		newDescription := helper.GenerateRandomString(64, nil)
		updateResponse, err := permissionClient.UpdateRole(context, &proto.UpdateRoleRequest{
			Selector: &selector,
			Data: &proto.UpdateRoleRequest_UpdateData{
				Name:        &newName,
				Description: &newDescription,
				Resources: []*proto.RoleResource{
					{
						ResourceID: resourceList[0].ID,
						ResourceName: resourceList[0].Name,
						Action:     proto.ResourceAction_Insert,
						Rules: []*proto.RoleResourceRule{
							{
								Key:   "TEST_UPDATE",
								Value: "TEST_UPDATE",
							},
						},
					},
				},
			},
		})
		if err != nil {
			t.Errorf("Unexpected results on update: %v", err)
			return false
		}
		if updateResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results on update")
			return false
		}

		queryUpdatedResponse, err := permissionClient.QueryRoles(context, &proto.QueryRolesRequest{
			Selector: &proto.RoleSelector{Name: &newName},
		})
		if err != nil {
			t.Errorf("Unexpected results after updated: %v", err)
			return false
		}
		if queryUpdatedResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after updated")
			return false
		}
		if queryUpdatedResponse.Data.Total != 1 {
			t.Errorf("Unexpected results after updated")
			return false
		}

		if queryUpdatedResponse.Data.Roles[0].Name != newName {
			t.Errorf("Unexpected results after updated")
			return false
		}

		roleList = append(roleList, queryResponse.Data.Roles[0])
		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// test bound role into user
	for _, role := range roleList {
		userID := int64(1) // 直接用  admin 账号测试
		applyResponse, err := permissionClient.ApplyRoleForUser(context, &proto.ApplyRoleForUserRequest{
			Role: &proto.RoleSelector{ID: &role.ID},
			User: &proto.UserSelector{ID: &userID},
		})
		if err != nil {
			t.Errorf("ApplyRoleForUser failed: %v", err)
			return
		}
		if applyResponse.State != proto.State_SUCCESS {
			t.Errorf("ApplyRoleForUser failed: %v", err)
			return
		}

		summaryResponse, err := permissionClient.SummaryForUser(context, &proto.SummaryForUserRequest{
			UserSelector: &proto.UserSelector{ID: &userID},
		})
		if err != nil {
			t.Errorf("SummaryForUser failed: %v", err)
			return
		}
		if summaryResponse.State != proto.State_SUCCESS {
			t.Errorf("SummaryForUser failed: %v", err)
			return
		}

		if len(summaryResponse.Data) <= 0 {
			t.Errorf("SummaryForUser failed: %v", err)
			return
		}

		// 检查查询到的权限是否与 role 一致
		for _, resources := range role.Resources {
			// 检查能否再 summaryResponse 找到当前这条 resource
			fined := slices.ContainsFunc[[]*proto.RoleResource](summaryResponse.Data, func(summary *proto.RoleResource) bool {
				// 如果 resource 和 action 对的上就进一步比对 rules
				if summary.ResourceID == resources.ResourceID && summary.Action == resources.Action {
					same := slices.EqualFunc[[]*proto.RoleResourceRule, []*proto.RoleResourceRule](summary.Rules, resources.Rules, func(a, b *proto.RoleResourceRule) bool {
						return a.Key == b.Key && a.Value == b.Value
					})

					if !same {
						// 不匹配有问题！
						t.Errorf("SummaryForUser failed: %v", err)
						return false
					}
				}
				return true
			})

			if !fined {
				// 找不到说明有问题
				t.Errorf("SummaryForUser failed: %v", err)
				return
			}
		}

		removeResponse, err := permissionClient.RemoveRoleForUser(context, &proto.RemoveRoleForUserRequest{
			Role: &proto.RoleSelector{ID: &role.ID},
			User: &proto.UserSelector{ID: &userID},
		})
		if err != nil {
			t.Errorf("ApplyRoleForUser failed: %v", err)
			return
		}
		if removeResponse.State != proto.State_SUCCESS {
			t.Errorf("ApplyRoleForUser failed: %v", err)
			return
		}

		summary2Response, err := permissionClient.SummaryForUser(context, &proto.SummaryForUserRequest{
			UserSelector: &proto.UserSelector{ID: &userID},
		})
		if err != nil {
			t.Errorf("SummaryForUser failed: %v", err)
			return
		}
		if summary2Response.State != proto.State_SUCCESS {
			t.Errorf("SummaryForUser failed: %v", err)
			return
		}

		// 检查查询到的权限是否与 role 一致
		for _, resources := range role.Resources {
			// 检查能否再 summaryResponse 找到当前这条 resource
			fined := slices.ContainsFunc[[]*proto.RoleResource](summary2Response.Data, func(summary *proto.RoleResource) bool {
				return summary.ResourceID == resources.ResourceID && summary.Action == resources.Action
			})

			if fined {
				// 找到说明有问题
				t.Errorf("SummaryForUser failed: %v", err)
				return
			}
		}
	}

	// test remove role
	for _, role := range roleList {
		removeResponse, err := permissionClient.DeleteRole(context, &proto.DeleteRoleRequest{
			Selector: &proto.RoleSelector{ID: &role.ID},
		})
		if err != nil {
			t.Errorf("RemoveRole failed: %v", err)
			return
		}
		if removeResponse.State != proto.State_SUCCESS {
			t.Errorf("RemoveRole failed: %v", err)
			return
		}
	}

	// test remove resource
	for _, resource := range resourceList {
		selector := proto.ResourceSelector{ID: &resource.ID}
		removeResponse, err := permissionClient.DeleteResource(context, &proto.DeleteResourceRequest{
			Selector: &selector,
		})
		if err != nil {
			t.Errorf("RemoveResource failed: %v", err)
			return
		}
		if removeResponse.State != proto.State_SUCCESS {
			t.Errorf("RemoveResource failed: %v", err)
			return
		}

		queryUpdatedResponse, err := permissionClient.QueryResources(context, &proto.QueryResourcesRequest{
			Selector: &selector,
		})
		if err != nil {
			t.Errorf("Unexpected results after deleted: %v", err)
			return
		}
		if queryUpdatedResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after deleted")
			return
		}
		if queryUpdatedResponse.Data.Total != 0 {
			t.Errorf("Unexpected results after deleted")
			return
		}
	}
}
