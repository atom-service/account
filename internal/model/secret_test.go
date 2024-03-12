package model

import (
	"context"
	"testing"
)

func TestCreateSecretTable(t *testing.T) {
	ut := &secretTable{}
	err := ut.CreateTable(context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestTruncateSecretTable(t *testing.T) {
	ut := &secretTable{}
	err := ut.TruncateTable(context.TODO())
	if err != nil {
		t.Error(err)
	}
}
