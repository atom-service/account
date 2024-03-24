package server

import (
	"context"
	"testing"
	"testing/quick"

	"github.com/atom-service/account/internal/database"
	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/proto"
	"github.com/atom-service/common/logger"
)

func TestPermissionUserRoleTable(t *testing.T) {
	context := context.TODO()
	logger.SetLevel(logger.InfoLevel)
	accountServer := &accountServer{}

	// 初始化数据库
	if err := database.Init(context); err != nil {
		panic(err)
	}

	// 初始化 model
	if err := model.Init(context); err != nil {
		panic(err)
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	// 已登录的用户
	signedInTokenUsers := []*proto.SignedInToken{}

	// sign up & sign in
	if err := quick.Check(func() bool {
		username := helper.GenerateRandomString(64)
		password := helper.GenerateRandomString(128)

		signUpResponse, err := accountServer.SignUp(context, &proto.SignUpRequest{
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

		signInResponse, err := accountServer.SignIn(context, &proto.SignInRequest{
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

	// secret create & query
	if err := quick.Check(func() bool {
		createResponse, err := accountServer.CreateSecret(context, &proto.CreateSecretRequest{})
		if err != nil {
			t.Errorf("SignUp failed: %v", err)
			return false
		}
		if createResponse.State != proto.State_NO_PERMISSION {
			t.Errorf("SignUp failed: %v", err)
			return false
		}

		queryResponse, err := accountServer.QuerySecrets(context, &proto.QuerySecretsRequest{})
		if err != nil {
			t.Errorf("SignUp failed: %v", err)
			return false
		}
		if queryResponse.State != proto.State_NO_PERMISSION {
			t.Errorf("SignUp failed: %v", err)
			return false
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}
}
