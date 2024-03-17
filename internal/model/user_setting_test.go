package model

import (
	"context"
	"math"
	"math/rand"
	"slices"
	"testing"
	"testing/quick"
)

func TestUserSettingTable(t *testing.T) {
	settingTable := &settingTable{}

	context := context.TODO()

	// 创建表
	if err := settingTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := settingTable.TruncateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	testRoles := []*Setting{}

	// create test & check result
	if err := quick.Check(func() bool {
		name := generateRandomString(64)
		value := generateRandomString(128)
		description := generateRandomString(128)

		testCreateParams := Setting{
			Key:         name,
			Value:       &value,
			Description: &description,
			UserID:      rand.Int63n(math.MaxInt32),
		}

		roleSelector := SettingSelector{
			Key:    &testCreateParams.Key,
			UserID: &testCreateParams.UserID,
		}

		if err := settingTable.CreateSetting(context, testCreateParams); err != nil {
			t.Errorf("Create failed: %v", err)
			return false
		}

		countResult, err := settingTable.CountSettings(context, SettingSelector{})
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return false
		}

		if countResult != int64(len(testRoles)+1) {
			t.Errorf("Count result are incorrect: %v", err)
			return false
		}

		queryCreateResult, err := settingTable.QuerySettings(context, roleSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryCreateResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryCreateResult)
			return false
		}

		if *queryCreateResult[0].Description != *testCreateParams.Description {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		// update test & check result
		newValue := generateRandomString(128)
		newDescription := generateRandomString(128)
		err = settingTable.UpdateSetting(context, roleSelector, &Setting{
			Value:       &newValue,
			Description: &newDescription,
		})
		if err != nil {
			t.Errorf("Update failed: %v", err)
			return false
		}

		queryUpdatedResult, err := settingTable.QuerySettings(context, SettingSelector{ID: queryCreateResult[0].ID}, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryUpdatedResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryUpdatedResult)
			return false
		}

		if *queryUpdatedResult[0].Value != newValue {
			t.Errorf("Unexpected results after updated: %v", queryCreateResult)
			return false
		}

		if *queryUpdatedResult[0].Description != newDescription {
			t.Errorf("Unexpected results after updated: %v", queryCreateResult)
			return false
		}

		if *queryUpdatedResult[0].CreatedTime != *queryCreateResult[0].CreatedTime {
			t.Errorf("Unexpected results after updated: %v", queryCreateResult)
			return false
		}

		// 最终的结果保存下来给后面的测试使用
		testRoles = append(testRoles, queryUpdatedResult...)
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

		queryPaginationResult, err := settingTable.QuerySettings(context, SettingSelector{}, &Pagination{
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
			if *queryPaginationResult[0].Value != *testRoles[offsetInt].Value {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
			if *queryPaginationResult[0].Description != *testRoles[offsetInt].Description {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// delete test & check result
	for _, testSecret := range testRoles {
		selector := SettingSelector{}
		randUseSeed := rand.Intn(2)
		if randUseSeed == 0 {
			selector.ID = testSecret.ID
		}

		if randUseSeed == 1 {
			selector.Key = &testSecret.Key
			selector.UserID = &testSecret.UserID
		}

		err := settingTable.DeleteSetting(context, selector)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
			return
		}

		queryDeletedResult, err := settingTable.QuerySettings(context, selector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return
		}

		if len(queryDeletedResult) != 0 {
			t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
			return
		}

		countUpdateResult, err := settingTable.CountSettings(context, selector)
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
