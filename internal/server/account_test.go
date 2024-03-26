package server

import (
	"context"
	"testing"
	"testing/quick"

	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/package/code"
	"github.com/atom-service/account/package/proto"
)

func TestPermissionUserRoleTable(t *testing.T) {
	context := context.TODO()
	testServer := createTestServer()
	accountClient := testServer.CreateAccountClientWithToken("")

	config := &quick.Config{
		MaxCount: 100,
	}

	// 已登录的用户
	signedInTokenUsers := []*proto.SignedInToken{}

	// sign up & sign in
	if err := quick.Check(func() bool {
		username := helper.GenerateRandomString(64, nil)
		password := helper.GenerateRandomString(128, nil)

		signUpResponse, err := accountClient.SignUp(context, &proto.SignUpRequest{
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

		signInResponse, err := accountClient.SignIn(context, &proto.SignInRequest{
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

		deleteResponse, err := accountClientWithUserAuth.DeleteSetting(context, &proto.DeleteSettingRequest{
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

		queryResponse, err := accountClientWithUserAuth.QuerySecrets(context, &proto.QuerySecretsRequest{})
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

		deleteResponse, err := accountClientWithUserAuth.DeleteSecret(context, &proto.DeleteSecreteRequest{
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

		deleteSecretResponse2, err := accountClientWithUserAuth.DeleteSecret(context, &proto.DeleteSecreteRequest{
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

}
