package server

import (
	"context"
	"testing"
	"testing/quick"
	"time"

	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/code"
	"github.com/atom-service/account/package/proto"
)

func TestAccountServer(t *testing.T) {
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

	adminAccountClient := testServer.CreateAccountClientWithToken(token)

	config := &quick.Config{
		MaxCount: 100,
	}

	// 已登录的用户
	signedInTokenUsers := []*proto.SignedInToken{}

	// sign up & sign in
	if err := quick.Check(func() bool {
		username := helper.GenerateRandomString(64, nil)
		password := helper.GenerateRandomString(128, nil)

		signUpResponse, err := adminAccountClient.SignUp(context, &proto.SignUpRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			t.Errorf("SignUp failed: %v", err)
			return false
		}
		if signUpResponse.State != proto.State_SUCCESS {
			t.Errorf("SignUp failed: %v", err)
			return false
		}

		signInResponse, err := adminAccountClient.SignIn(context, &proto.SignInRequest{
			Selector: &proto.UserSelector{Username: &username},
			Password: password,
		})
		if err != nil {
			t.Errorf("SignUp failed: %v", err)
			return false
		}
		if signInResponse.State != proto.State_SUCCESS {
			t.Errorf("SignUp failed: %v", err)
			return false
		}

		signedInTokenUsers = append(signedInTokenUsers, signInResponse.Token)
		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// pagination test & check result
	for _, user := range signedInTokenUsers {
		var offsetUint64 = int64(0)
		var limitUint64 = int64(1)

		selector := &proto.UserSelector{ID: &user.UserID}
		pagination := &proto.PaginationOption{Offset: &offsetUint64, Limit: &limitUint64}
		accountClientWithUserAuth := testServer.CreateAccountClientWithToken(user.Token)
		request := &proto.QueryUsersRequest{Selector: selector, Pagination: pagination}
		queryPaginationResult, err := accountClientWithUserAuth.QueryUsers(context, request)
		if err != nil || queryPaginationResult.State != proto.State_SUCCESS {
			t.Errorf("Query failed: %v", err)
			return
		}

		if queryPaginationResult.Data.Total != int64(1) {
			t.Errorf("Query result length are incorrect: %v", queryPaginationResult)
			return
		}

		if queryPaginationResult.Data.Users[0].ID != user.UserID {
			t.Errorf("Query result length are incorrect: %v", queryPaginationResult)
			return
		}
	}

	// test setting create & query & delete
	for _, user := range signedInTokenUsers {
		testKey := helper.GenerateRandomString(64, nil)
		testValue := helper.GenerateRandomString(512, nil)

		accountClientWithUserAuth := testServer.CreateAccountClientWithToken(user.Token)
		createResponse, err := accountClientWithUserAuth.CreateSetting(context, &proto.CreateSettingRequest{
			Key:   testKey,
			Value: testValue,
		})
		if err != nil {
			t.Errorf("CreateSetting failed: %v", err)
			return
		}
		if createResponse.State != proto.State_SUCCESS {
			t.Errorf("CreateSetting failed: %v", err)
			return
		}

		queryResponse, err := accountClientWithUserAuth.QuerySettings(context, &proto.QuerySettingsRequest{})
		if err != nil {
			t.Errorf("Unexpected results after query: %v", err)
			return
		}
		if queryResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after query: %s", queryResponse.State.String())
			return
		}

		if queryResponse.Data.Total == 0 {
			t.Errorf("Unexpected results after query")
			return
		}

		newTestKey := helper.GenerateRandomString(64, nil)
		newTestValue := helper.GenerateRandomString(128, nil)
		updateResponse, err := accountClientWithUserAuth.UpdateSetting(context, &proto.UpdateSettingRequest{
			Selector: &proto.SettingSelector{
				Key: &queryResponse.Data.Settings[0].Key,
			},
			Data: &proto.UpdateSettingRequest_UpdateData{
				Key:   &newTestKey,
				Value: &newTestValue,
			},
		})
		if err != nil {
			t.Errorf("Unexpected results on update: %v", err)
			return
		}
		if updateResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results on update: %v", err)
			return
		}

		queryResponse2, err := accountClientWithUserAuth.QuerySettings(context, &proto.QuerySettingsRequest{
			Selector: &proto.SettingSelector{
				Key: &newTestKey,
			},
		})
		if err != nil {
			t.Errorf("Unexpected results after updated: %v", err)
			return
		}
		if queryResponse2.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after updated: %s", queryResponse.State.String())
			return
		}

		if queryResponse2.Data.Total == 0 {
			t.Errorf("Unexpected results after updated")
			return
		}

		if queryResponse2.Data.Settings[0].Key != newTestKey {
			t.Errorf("Unexpected results after updated")
			return
		}

		if queryResponse2.Data.Settings[0].Value != newTestValue {
			t.Errorf("Unexpected results after updated")
			return
		}

		deleteResponse, err := accountClientWithUserAuth.DeleteSettings(context, &proto.DeleteSettingsRequest{
			Selector: &proto.SettingSelector{
				Key: &queryResponse.Data.Settings[0].Key,
			},
		})
		if err != nil {
			t.Errorf("Unexpected results after deleted: %v", err)
			return
		}
		if deleteResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after deleted: %v", err)
			return
		}
	}

	// test label create & query & delete
	for _, user := range signedInTokenUsers {
		testKey := helper.GenerateRandomString(64, nil)
		testValue := helper.GenerateRandomString(128, nil)

		accountClientWithUserAuth := testServer.CreateAccountClientWithToken(user.Token)
		createResponse, err := accountClientWithUserAuth.UpsertLabel(context, &proto.UpsertLabelRequest{
			Key:   &testKey,
			Value: &testValue,
		})
		if err != nil {
			t.Errorf("UpsertLabel failed: %v", err)
			return
		}
		if createResponse.State != proto.State_SUCCESS {
			t.Errorf("UpsertLabel failed: %v", err)
			return
		}

		queryResponse, err := accountClientWithUserAuth.QueryLabels(context, &proto.QueryLabelsRequest{})
		if err != nil {
			t.Errorf("Unexpected results after query: %v", err)
			return
		}
		if queryResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after query: %s", queryResponse.State.String())
			return
		}

		if queryResponse.Data.Total == 0 {
			t.Errorf("Unexpected results after query")
			return
		}

		newTestKey := helper.GenerateRandomString(64, nil)
		newTestValue := helper.GenerateRandomString(128, nil)
		updateResponse, err := accountClientWithUserAuth.UpsertLabel(context, &proto.UpsertLabelRequest{
			Key:   &newTestKey,
			Value: &newTestValue,
		})
		if err != nil {
			t.Errorf("Unexpected results on update: %v", err)
			return
		}
		if updateResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results on update: %v", err)
			return
		}

		queryResponse2, err := accountClientWithUserAuth.QueryLabels(context, &proto.QueryLabelsRequest{
			Selector: &proto.LabelSelector{
				Key: &newTestKey,
			},
		})
		if err != nil {
			t.Errorf("Unexpected results after updated: %v", err)
			return
		}
		if queryResponse2.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after updated: %s", queryResponse.State.String())
			return
		}

		if queryResponse2.Data.Total == 0 {
			t.Errorf("Unexpected results after updated")
			return
		}

		if queryResponse2.Data.Labels[0].Key != newTestKey {
			t.Errorf("Unexpected results after updated")
			return
		}

		if queryResponse2.Data.Labels[0].Value != newTestValue {
			t.Errorf("Unexpected results after updated")
			return
		}

		deleteResponse, err := accountClientWithUserAuth.DeleteLabels(context, &proto.DeleteLabelsRequest{
			Selector: &proto.LabelSelector{
				Key: &newTestKey,
			},
		})
		if err != nil {
			t.Errorf("Unexpected results after deleted: %v", err)
			return
		}
		if deleteResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after deleted: %v", err)
			return
		}
	}

	// test secret create & query & delete
	for _, user := range signedInTokenUsers {
		accountClientWithUserAuth := testServer.CreateAccountClientWithToken(user.Token)
		createResponse, err := accountClientWithUserAuth.CreateSecret(context, &proto.CreateSecretRequest{})
		if err != nil {
			t.Errorf("CreateSecret failed: %v", err)
			return
		}
		if createResponse.State != proto.State_SUCCESS {
			t.Errorf("CreateSecret failed: %v", err)
			return
		}

		queryResponse, err := accountClientWithUserAuth.QuerySecrets(context, &proto.QuerySecretsRequest{
			Selector: &proto.SecretSelector{},
		})
		if err != nil {
			t.Errorf("QuerySecrets failed: %v", err)
			return
		}
		if queryResponse.State != proto.State_SUCCESS {
			t.Errorf("QuerySecrets failed: %s", queryResponse.State.String())
			return
		}

		if queryResponse.Data.Total == 0 {
			t.Errorf("QuerySecrets failed")
			return
		}

		deleteResponse, err := accountClientWithUserAuth.DeleteSecrets(context, &proto.DeleteSecretsRequest{
			Selector: &proto.SecretSelector{
				UserID: &user.UserID,
				Key:    &queryResponse.Data.Secrets[0].Key,
			},
		})
		if err != nil {
			t.Errorf("DeleteSecret failed: %v", err)
			return
		}
		if deleteResponse.State != proto.State_FAILURE {
			t.Errorf("DeleteSecret failed: %v", err)
			return
		}
		if deleteResponse.Code != code.USER_SECRET_NOT_DISABLED {
			t.Errorf("DeleteSecret failed: %v", err)
			return
		}

		disableResponse, err := accountClientWithUserAuth.DisableSecret(context, &proto.DisableSecretRequest{
			Selector: &proto.SecretSelector{
				UserID: &user.UserID,
				Key:    &queryResponse.Data.Secrets[0].Key,
			},
		})
		if err != nil {
			t.Errorf("DisableSecret failed: %v", err)
			return
		}
		if disableResponse.State != proto.State_SUCCESS {
			t.Errorf("DisableSecret failed: %v", err)
			return
		}

		deleteSecretResponse2, err := accountClientWithUserAuth.DeleteSecrets(context, &proto.DeleteSecretsRequest{
			Selector: &proto.SecretSelector{
				UserID: &user.UserID,
				Key:    &queryResponse.Data.Secrets[0].Key,
			},
		})
		if err != nil {
			t.Errorf("DeleteSecret failed: %v", err)
			return
		}
		if deleteSecretResponse2.State != proto.State_SUCCESS {
			t.Errorf("DeleteSecret failed: %v", err)
			return
		}
	}

	// test delete user
	for _, user := range signedInTokenUsers {
		selector := proto.UserSelector{ID: &user.UserID}
		deleteResponse, err := adminAccountClient.DeleteUser(context, &proto.DeleteUserRequest{
			Selector: &selector,
		})
		if err != nil {
			t.Errorf("DeleteUser failed: %v", err)
			return
		}
		if deleteResponse.State != proto.State_SUCCESS {
			t.Errorf("DeleteUser failed: %v", err)
			return
		}

		queryResponse, err := adminAccountClient.QueryUsers(context, &proto.QueryUsersRequest{
			Selector: &selector,
		})
		if err != nil {
			t.Errorf("Unexpected results after deleted: %v", err)
			return
		}
		if queryResponse.State != proto.State_SUCCESS {
			t.Errorf("Unexpected results after deleted")
			return
		}

	}
}
