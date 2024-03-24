package model

import (
	"context"
	"math/rand"
	"slices"
	"testing"
	"testing/quick"

	"github.com/atom-service/account/internal/helper"
)

func TestUserTable(t *testing.T) {
	// 创建一个用户表实例
	userTable := &userTable{}

	context := context.TODO()

	// 创建表
	if err := userTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := userTable.TruncateTable(context); err != nil {
		t.Errorf("Truncate table failed: %v", err)
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	testUsers := []*User{}

	// create test & check result
	if err := quick.Check(func() bool {
		username := helper.GenerateRandomString(64)
		password := helper.GenerateRandomString(128)

		testUser := User{
			Username: &username,
			Password: &password,
		}

		if err := userTable.CreateUser(context, testUser); err != nil {
			t.Errorf("Create failed: %v", err)
			return false
		}

		countResult, err := userTable.CountUsers(context, UserSelector{})
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return false
		}

		if countResult != int64(len(testUsers)+1) {
			t.Errorf("Count result are incorrect: %v", err)
			return false
		}

		queryCreateResult, err := userTable.QueryUsers(context, UserSelector{Username: &username}, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryCreateResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryCreateResult)
			return false
		}

		if *queryCreateResult[0].Username != username {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		// update test & check result
		newPassword := helper.GenerateRandomString(64)
		err = userTable.UpdateUser(context, UserSelector{Username: &username}, &User{
			Password: &newPassword,
		})
		if err != nil {
			t.Errorf("Update failed: %v", err)
			return false
		}

		queryUpdatedResult, err := userTable.QueryUsers(context, UserSelector{Username: &username}, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryUpdatedResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryUpdatedResult)
			return false
		}

		if *queryUpdatedResult[0].Password != newPassword {
			t.Errorf("Unexpected results after updated: %v", queryCreateResult)
			return false
		}

		if *queryUpdatedResult[0].CreatedTime != *queryCreateResult[0].CreatedTime {
			t.Errorf("Unexpected results after updated: %v", queryCreateResult)
			return false
		}

		// 最终的结果保存下来给后面的测试使用
		testUsers = append(testUsers, queryUpdatedResult...)
		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// pagination test & check result
	if err := quick.Check(func() bool {
		var offsetInt = rand.Intn(config.MaxCount)
		var limitInt = rand.Intn(config.MaxCount)
		var offsetUint64 = int64(offsetInt)
		var limitUint64 = int64(limitInt)

		queryPaginationResult, err := userTable.QueryUsers(context, UserSelector{}, &Pagination{
			Offset: &offsetUint64,
			Limit:  &limitUint64,
		}, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		expectedSize := slices.Min[[]int]([]int{config.MaxCount - offsetInt, limitInt})
		if len(queryPaginationResult) != expectedSize {
			t.Errorf("Query result length are incorrect: %v", queryPaginationResult)
			return false
		}

		if expectedSize > 0 { // 最后一条不需要检查了，删光了
			if *queryPaginationResult[0].Username != *testUsers[offsetInt].Username {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// delete test & check result
	for _, testUser := range testUsers {
		userSelector := UserSelector{}
		randUseSeed := rand.Intn(3)
		if randUseSeed == 0 { // 0 use id
			userSelector.ID = testUser.ID
		}
		if randUseSeed == 1 { // 1 use username
			userSelector.Username = testUser.Username
		}

		if randUseSeed == 2 { // 2 use id with username
			userSelector.ID = testUser.ID
			userSelector.Username = testUser.Username
		}

		err := userTable.DeleteUser(context, userSelector)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
			return
		}

		queryDeletedResult, err := userTable.QueryUsers(context, userSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return
		}

		if len(queryDeletedResult) != 0 {
			t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
			return
		}

		countUpdateResult, err := userTable.CountUsers(context, userSelector)
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return
		}

		if countUpdateResult != 0 {
			t.Errorf("Unexpected results after deleted: %v", err)
			return
		}
	}
}

func TestAdminUserInit(t *testing.T) {
	// 创建一个用户表实例
	userTable := &userTable{}

	context := context.TODO()

	if err := userTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := userTable.InitAdminUser(context); err != nil {
		t.Errorf("Init admin user failed: %v", err)
		return
	}

	selectID := int64(0)
	selector := UserSelector{ID: &selectID}
	queryDeletedResult, err := userTable.QueryUsers(context, selector, nil, nil)
	if err != nil {
		t.Errorf("Query failed: %v", err)
		return
	}

	if len(queryDeletedResult) != 0 {
		t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
		return
	}
}
