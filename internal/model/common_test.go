package model

import (
	"context"
	"testing"

	"github.com/atom-service/account/internal/database"
)

func TestMain(m *testing.M) {
	database.Init(context.TODO())
	m.Run()
}
