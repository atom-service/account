package model

import (
	"context"
	"testing"
)

func TestCreateUserTable(t *testing.T) {
	ut := &userTable{}
	err := ut.CreateTable(context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestTruncateUserTable(t *testing.T) {
	ut := &userTable{}
	err := ut.TruncateTable(context.TODO())
	if err != nil {
		t.Error(err)
	}
}