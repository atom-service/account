package model

import (
	"context"
	"math"
	"math/rand"
	"slices"
	"testing"
	"testing/quick"
)

func TestSecretTable(t *testing.T) {
	secretTable := &secretTable{}

	context := context.TODO()

	// 创建表
	if err := secretTable.CreateTable(context); err != nil {
		t.Errorf("Create table failed: %v", err)
		return
	}

	if err := secretTable.TruncateTable(context); err != nil {
		t.Errorf("Truncate table failed: %v", err)
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	testSecrets := []*Secret{}

	// create test & check result
	if err := quick.Check(func() bool {
		description := generateRandomString(128)

		testCreateParams := CreateSecretParams{
			UserID:      rand.Int63n(math.MaxInt32),
			Description: &description,
		}

		randTypeSeed := rand.Intn(2)
		if randTypeSeed == 0 { // 0 use system
			testCreateParams.Type = SystemSecretType
		}

		if randTypeSeed == 1 { // 1 use user
			testCreateParams.Type = UserSecretType
		}

		if err := secretTable.CreateSecret(context, testCreateParams); err != nil {
			t.Errorf("Create failed: %v", err)
			return false
		}

		countResult, err := secretTable.CountSecrets(context, SecretSelector{})
		if err != nil {
			t.Errorf("Count failed: %v", err)
			return false
		}

		if countResult != int64(len(testSecrets)+1) {
			t.Errorf("Count result are incorrect: %v", err)
			return false
		}

		queryCreateSelector := SecretSelector{UserID: &testCreateParams.UserID, Type: &testCreateParams.Type}
		queryCreateResult, err := secretTable.QuerySecrets(context, queryCreateSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return false
		}

		if len(queryCreateResult) != 1 {
			t.Errorf("Query result length are incorrect: %v", queryCreateResult)
			return false
		}

		if *queryCreateResult[0].UserID != testCreateParams.UserID {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		if *queryCreateResult[0].Type != testCreateParams.Type {
			t.Errorf("Query result are incorrect: %v", queryCreateResult)
			return false
		}

		// 最终的结果保存下来给后面的测试使用
		testSecrets = append(testSecrets, queryCreateResult...)
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

		queryPaginationResult, err := secretTable.QuerySecrets(context, SecretSelector{}, &Pagination{
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
			if *queryPaginationResult[0].Key != *testSecrets[offsetInt].Key {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}

			if *queryPaginationResult[0].Value != *testSecrets[offsetInt].Value {
				t.Errorf("Query result are incorrect: %v", queryPaginationResult)
				return false
			}
		}

		return true
	}, config); err != nil {
		t.Errorf("Test failed: %v", err)
	}

	// delete test & check result
	for _, testSecret := range testSecrets {
		secretSelector := SecretSelector{}
		randUseSeed := rand.Intn(3)
		if randUseSeed == 0 {
			secretSelector.Key = testSecret.Key
		}

		if randUseSeed == 1 {
			secretSelector.Key = testSecret.Key
			secretSelector.Type = testSecret.Type
		}

		if randUseSeed == 2 {
			secretSelector.Key = testSecret.Key
			secretSelector.Type = testSecret.Type
			secretSelector.UserID = testSecret.UserID
		}

		err := secretTable.DeleteSecret(context, secretSelector)
		if err != nil {
			t.Errorf("Delete failed: %v", err)
			return
		}

		queryDeletedResult, err := secretTable.QuerySecrets(context, secretSelector, nil, nil)
		if err != nil {
			t.Errorf("Query failed: %v", err)
			return
		}

		if len(queryDeletedResult) != 0 {
			t.Errorf("Unexpected results after deleted: %v", queryDeletedResult)
			return
		}

		countUpdateResult, err := secretTable.CountSecrets(context, secretSelector)
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
