package model

import (
	"context"
	"testing"

	"github.com/atom-service/account/internal/db"
)

func TestMain(m *testing.M) {
	db.Init()
	m.Run()
}

func TestCreateTable(t *testing.T) {
	ut := &userTable{}
	err := ut.CreateTable(context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestTruncateTable(t *testing.T) {
	ut := &userTable{}
	err := ut.TruncateTable(context.TODO())
	if err != nil {
		t.Error(err)
	}
}


func FuzzTestUsers(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, i int64) {

	})
}
