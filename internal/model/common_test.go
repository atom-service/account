package model

import (
	"testing"

	"github.com/atom-service/account/internal/db"
)

func TestMain(m *testing.M) {
	db.Init()
	m.Run()
}
