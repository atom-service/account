package model

import (
	"context"
	"testing"

	"github.com/atom-service/account/internal/database"
	"github.com/atom-service/common/logger"
)

func TestMain(m *testing.M) {
	logger.SetLevel(logger.InfoLevel)
	database.Init(context.TODO())
	m.Run()
}
