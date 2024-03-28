package server

import (
	"context"
	"testing"
	"testing/quick"

	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/package/proto"
)

func TestPermissionServer(t *testing.T) {
	context := context.TODO()
	testServer := createTestServer()

	// TODO: 得 admin 的 token
	permissionClient := testServer.CreatePermissionClientWithToken("")

	config := &quick.Config{
		MaxCount: 100,
	}

	roleList := []*proto.Role{}
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
			Selector: &selector,
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

		if queryUpdatedResponse.Data.Resources[0].Name != newDescription {
			t.Errorf("Unexpected results after updated")
			return false
		}

		resourceList = append(resourceList, queryResponse.Data.Resources[0])
		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

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
			Selector: &selector,
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

		if queryUpdatedResponse.Data.Roles[0].Name != newDescription {
			t.Errorf("Unexpected results after updated")
			return false
		}

		roleList = append(roleList, queryResponse.Data.Roles[0])
		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// test bound resource into role
	// for _, role := range roleList {

	// }
}
