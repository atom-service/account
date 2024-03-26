package model

import (
	"context"
	"testing"
)

func TestMain(m *testing.M) {
	InitDB(context.TODO())
	m.Run()
}
