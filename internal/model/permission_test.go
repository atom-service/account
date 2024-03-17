package model

import (
	"context"
	"math"
	"math/rand"
	"slices"
	"testing"
	"testing/quick"
)

func TestPermissionRoleTable(t *testing.T) {
	roleTable := &roleTable{}

	context := context.TODO()

	// 创建表
	if err := roleTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := roleTable.TruncateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	testRoles := []*Role{}

	// create test & check result
	if err := quick.Check(func() bool {
		name := generateRandomString(64)
		description := generateRandomString(128)

		testCreateParams := Role{
			Name:        &name,
			Description: &description,
		}

		roleSelector := RoleSelector{Name: testCreateParams.Name}

		if err := roleTable.CreateRole(context, testCreateParams); err != nil {
			t.Errorf("Create failed: %v", err)
			return false
		}

		countResult, err := roleTable.CountRoles(context, RoleSelector{})
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return false
		}

		if countResult != int64(len(testRoles)+1) {
			t.Errorf("Count result are incorrect: %v", err)
			return false
		}

		queryCreateResult, err := roleTable.QueryRoles(context, roleSelector, nil, nil)
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
		newName := generateRandomString(64)
		newDescription := generateRandomString(128)
		err = roleTable.UpdateRole(context, roleSelector, &Role{
			Name:        &newName,
			Description: &newDescription,
		})
		if err != nil {
			t.Errorf("Update failed: %v", err)
			return false
		}

		queryUpdatedResult, err := roleTable.QueryRoles(context, RoleSelector{ID: queryCreateResult[0].ID}, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryUpdatedResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryUpdatedResult)
			return false
		}

		if *queryUpdatedResult[0].Name != newName {
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

		queryPaginationResult, err := roleTable.QueryRoles(context, RoleSelector{}, &Pagination{
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
			if *queryPaginationResult[0].Name != *testRoles[offsetInt].Name {
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
		roleSelector := RoleSelector{}
		randUseSeed := rand.Intn(3)
		if randUseSeed == 0 {
			roleSelector.ID = testSecret.ID
		}

		if randUseSeed == 1 {
			roleSelector.Name = testSecret.Name
		}

		if randUseSeed == 2 {
			roleSelector.ID = testSecret.ID
			roleSelector.Name = testSecret.Name
		}

		err := roleTable.DeleteRole(context, roleSelector)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
			return
		}

		queryDeletedResult, err := roleTable.QueryRoles(context, roleSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return
		}

		if len(queryDeletedResult) != 0 {
			t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
			return
		}

		countUpdateResult, err := roleTable.CountRoles(context, roleSelector)
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

func TestPermissionResourceTable(t *testing.T) {
	resourceTable := &resourceTable{}

	context := context.TODO()

	// 创建表
	if err := resourceTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := resourceTable.TruncateTable(context); err != nil {
		t.Errorf("Truncate table failed: %v", err)
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	testResources := []*Resource{}

	// create test & check result
	if err := quick.Check(func() bool {
		name := generateRandomString(64)
		description := generateRandomString(128)

		testCreateParams := Resource{
			Name:        &name,
			Description: &description,
		}

		resourceSelector := ResourceSelector{Name: testCreateParams.Name}

		if err := resourceTable.CreateResource(context, testCreateParams); err != nil {
			t.Errorf("Create failed: %v", err)
			return false
		}

		countResult, err := resourceTable.CountResources(context, ResourceSelector{})
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return false
		}

		if countResult != int64(len(testResources)+1) {
			t.Errorf("Count result are incorrect: %v", err)
			return false
		}

		queryCreateResult, err := resourceTable.QueryResources(context, resourceSelector, nil, nil)
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
		newName := generateRandomString(64)
		newDescription := generateRandomString(128)
		err = resourceTable.UpdateResource(context, resourceSelector, &Resource{
			Name:        &newName,
			Description: &newDescription,
		})
		if err != nil {
			t.Errorf("Update failed: %v", err)
			return false
		}

		queryUpdatedResult, err := resourceTable.QueryResources(context, ResourceSelector{ID: queryCreateResult[0].ID}, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryUpdatedResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryUpdatedResult)
			return false
		}

		if *queryUpdatedResult[0].Name != newName {
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
		testResources = append(testResources, queryUpdatedResult...)
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

		queryPaginationResult, err := resourceTable.QueryResources(context, ResourceSelector{}, &Pagination{
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
			if *queryPaginationResult[0].Description != *testResources[offsetInt].Description {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// delete test & check result
	for _, testSecret := range testResources {
		resourceSelector := ResourceSelector{}
		randUseSeed := rand.Intn(3)
		if randUseSeed == 0 {
			resourceSelector.ID = testSecret.ID
		}

		if randUseSeed == 1 {
			resourceSelector.Name = testSecret.Name
		}

		if randUseSeed == 2 {
			resourceSelector.ID = testSecret.ID
			resourceSelector.Name = testSecret.Name
		}

		err := resourceTable.DeleteResource(context, resourceSelector)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
			return
		}

		queryDeletedResult, err := resourceTable.QueryResources(context, resourceSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return
		}

		if len(queryDeletedResult) != 0 {
			t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
			return
		}

		countUpdateResult, err := resourceTable.CountResources(context, resourceSelector)
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

func TestPermissionRoleResourceTable(t *testing.T) {
	roleResourceTable := &roleResourceTable{}

	context := context.TODO()

	// 创建表
	if err := roleResourceTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := roleResourceTable.TruncateTable(context); err != nil {
		t.Errorf("Truncate table failed: %v", err)
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	testRoleResources := []*RoleResource{}

	// create test & check result
	if err := quick.Check(func() bool {
		testCreateParams := RoleResource{
			ResourceID: rand.Int63n(math.MaxInt32),
		}

		randUseSeed := rand.Intn(4)
		if randUseSeed == 0 {
			testCreateParams.Action = RoleResourceInsertAction
		}
		if randUseSeed == 1 {
			testCreateParams.Action = RoleResourceQueryAction
		}
		if randUseSeed == 2 {
			testCreateParams.Action = RoleResourceUpdateAction
		}
		if randUseSeed == 3 {
			testCreateParams.Action = RoleResourceDeleteAction
		}

		roleResourceSelector := RoleResourceSelector{
			Action:     &testCreateParams.Action,
			ResourceID: &testCreateParams.ResourceID,
		}

		if err := roleResourceTable.CreateRoleResource(context, testCreateParams); err != nil {
			t.Errorf("Create failed: %v", err)
			return false
		}

		countResult, err := roleResourceTable.CountRoleResources(context, RoleResourceSelector{})
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return false
		}

		if countResult != int64(len(testRoleResources)+1) {
			t.Errorf("Count result are incorrect: %v", err)
			return false
		}

		queryCreateResult, err := roleResourceTable.QueryRoleResources(context, roleResourceSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryCreateResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryCreateResult)
			return false
		}

		if queryCreateResult[0].Action != testCreateParams.Action {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		if queryCreateResult[0].ResourceID != testCreateParams.ResourceID {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		// 最终的结果保存下来给后面的测试使用
		testRoleResources = append(testRoleResources, queryCreateResult...)
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

		queryPaginationResult, err := roleResourceTable.QueryRoleResources(context, RoleResourceSelector{}, &Pagination{
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
			if queryPaginationResult[0].Action != testRoleResources[offsetInt].Action {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
			if queryPaginationResult[0].ResourceID != testRoleResources[offsetInt].ResourceID {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// delete test & check result
	for _, testSecret := range testRoleResources {
		roleResourceSelector := RoleResourceSelector{}
		randUseSeed := rand.Intn(2)
		if randUseSeed == 0 {
			roleResourceSelector.ID = testSecret.ID
		}

		if randUseSeed == 1 {
			roleResourceSelector.Action = &testSecret.Action
			roleResourceSelector.ResourceID = &testSecret.ResourceID
		}

		err := roleResourceTable.DeleteRoleResource(context, roleResourceSelector)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
			return
		}

		queryDeletedResult, err := roleResourceTable.QueryRoleResources(context, roleResourceSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return
		}

		if len(queryDeletedResult) != 0 {
			t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
			return
		}

		countUpdateResult, err := roleResourceTable.CountRoleResources(context, roleResourceSelector)
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

func TestPermissionRoleResourceRuleTable(t *testing.T) {
	roleResourceRuleTable := &roleResourceRuleTable{}

	context := context.TODO()

	// 创建表
	if err := roleResourceRuleTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := roleResourceRuleTable.TruncateTable(context); err != nil {
		t.Errorf("Truncate table failed: %v", err)
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	testRoleResourceRules := []*RoleResourceRule{}

	// create test & check result
	if err := quick.Check(func() bool {
		testCreateParams := RoleResourceRule{
			Key:            generateRandomString(64),
			Value:          generateRandomString(128),
			RoleResourceID: rand.Int63n(math.MaxInt32),
		}

		roleResourceRuleSelector := RoleResourceRuleSelector{
			Key:            &testCreateParams.Key,
			RoleResourceID: &testCreateParams.RoleResourceID,
		}

		if err := roleResourceRuleTable.CreateRoleResourceRule(context, testCreateParams); err != nil {
			t.Errorf("Create failed: %v", err)
			return false
		}

		countResult, err := roleResourceRuleTable.CountRoleResourceRules(context, RoleResourceRuleSelector{})
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return false
		}

		if countResult != int64(len(testRoleResourceRules)+1) {
			t.Errorf("Count result are incorrect: %v", err)
			return false
		}

		queryCreateResult, err := roleResourceRuleTable.QueryRoleResourceRules(context, roleResourceRuleSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryCreateResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryCreateResult)
			return false
		}

		if queryCreateResult[0].Key != testCreateParams.Key {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		if queryCreateResult[0].Value != testCreateParams.Value {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		if queryCreateResult[0].RoleResourceID != testCreateParams.RoleResourceID {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		// 最终的结果保存下来给后面的测试使用
		testRoleResourceRules = append(testRoleResourceRules, queryCreateResult...)
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

		queryPaginationResult, err := roleResourceRuleTable.QueryRoleResourceRules(context, RoleResourceRuleSelector{}, &Pagination{
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
			if queryPaginationResult[0].Key != testRoleResourceRules[offsetInt].Key {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
			if queryPaginationResult[0].Value != testRoleResourceRules[offsetInt].Value {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
			if queryPaginationResult[0].RoleResourceID != testRoleResourceRules[offsetInt].RoleResourceID {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// delete test & check result
	for _, testSecret := range testRoleResourceRules {
		roleResourceRuleSelector := RoleResourceRuleSelector{}
		randUseSeed := rand.Intn(2)
		if randUseSeed == 0 {
			roleResourceRuleSelector.ID = testSecret.ID
		}

		if randUseSeed == 1 {
			roleResourceRuleSelector.Key = &testSecret.Key
			roleResourceRuleSelector.RoleResourceID = &testSecret.RoleResourceID
		}

		err := roleResourceRuleTable.DeleteRoleResourceRule(context, roleResourceRuleSelector)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
			return
		}

		queryDeletedResult, err := roleResourceRuleTable.QueryRoleResourceRules(context, roleResourceRuleSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return
		}

		if len(queryDeletedResult) != 0 {
			t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
			return
		}

		countUpdateResult, err := roleResourceRuleTable.CountRoleResourceRules(context, roleResourceRuleSelector)
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

func TestPermissionUserRoleTable(t *testing.T) {
	userRoleTable := &userRoleTable{}

	context := context.TODO()

	// 创建表
	if err := userRoleTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := userRoleTable.TruncateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	testUserRoles := []*UserRole{}

	// create test & check result
	if err := quick.Check(func() bool {
		testCreateParams := UserRole{
			UserID: rand.Int63n(math.MaxInt32),
			RoleID: rand.Int63n(math.MaxInt32),
		}

		userRoleSelector := UserRoleSelector{
			UserID: &testCreateParams.UserID,
			RoleID: &testCreateParams.RoleID,
		}

		if err := userRoleTable.CreateUserRole(context, testCreateParams); err != nil {
			t.Errorf("Create failed: %v", err)
			return false
		}

		countResult, err := userRoleTable.CountUserRole(context, UserRoleSelector{})
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return false
		}

		if countResult != int64(len(testUserRoles)+1) {
			t.Errorf("Count result are incorrect: %v", err)
			return false
		}

		queryCreateResult, err := userRoleTable.QueryUserRole(context, userRoleSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryCreateResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryCreateResult)
			return false
		}

		if queryCreateResult[0].RoleID != testCreateParams.RoleID {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		if queryCreateResult[0].UserID != testCreateParams.UserID {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		// 最终的结果保存下来给后面的测试使用
		testUserRoles = append(testUserRoles, queryCreateResult...)
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

		queryPaginationResult, err := userRoleTable.QueryUserRole(context, UserRoleSelector{}, &Pagination{
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
			if queryPaginationResult[0].UserID != testUserRoles[offsetInt].UserID {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
			if queryPaginationResult[0].RoleID != testUserRoles[offsetInt].RoleID {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// delete test & check result
	for _, testSecret := range testUserRoles {
		selector := UserRoleSelector{}
		randUseSeed := rand.Intn(2)
		if randUseSeed == 0 {
			selector.ID = testSecret.ID
		}

		if randUseSeed == 1 {
			selector.UserID = &testSecret.UserID
			selector.RoleID = &testSecret.RoleID
		}

		err := userRoleTable.DeleteUserRole(context, selector)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
			return
		}

		queryDeletedResult, err := userRoleTable.QueryUserRole(context, selector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return
		}

		if len(queryDeletedResult) != 0 {
			t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
			return
		}

		countUpdateResult, err := userRoleTable.CountUserRole(context, selector)
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
