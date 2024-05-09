package model

import (
	"context"
	"fmt"
	"testing"

	"github.com/atom-service/account/internal/config"
	"github.com/yinxulai/sqls"
)

func dropTables(ctx context.Context) (err error) {
	dropTables := []string{
		roleTableName,
		userRoleTableName,
		resourceTableName,
		roleResourceTableName,
		userLabelTableName,
		userSettingTableName,
		secretTableName,
		userTableName,
	}

	for _, tableName := range dropTables {
		if _, err := Database.ExecContext(ctx, sqls.DROP_TABLE(tableName).IF_EXISTS().String()); err != nil {
			return fmt.Errorf("drop table %s failed: %v", tableName, err)
		}
	}

	return
}

func TestMain(m *testing.M) {
	config.MustInit("../../config.yaml")
	InitDB(context.TODO())
	dropTables(context.TODO())
	m.Run()
}
