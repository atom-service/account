package model

import (
	"context"

	"github.com/atom-service/account/package/protos"
)

func Init(ctx context.Context) error {
	if err := UserTable.CreateTable(ctx);err!= nil {
		return err
	}

	if err := SettingTable.CreateTable(ctx);err!= nil {
		return err
	}

	if err := SecretTable.CreateTable(ctx);err!= nil {
		return err
	}

	if err := LabelTable.CreateTable(ctx);err!= nil {
		return err
	}

	if err := RoleTable.CreateTable(ctx);err!= nil {
		return err
	}
	if err := UserRoleTable.CreateTable(ctx);err!= nil {
		return err
	}
	if err := ResourceTable.CreateTable(ctx);err!= nil {
		return err
	}
	if err := RoleResourceTable.CreateTable(ctx);err!= nil {
		return err
	}
	if err := RoleResourceRuleTable.CreateTable(ctx);err!= nil {
		return err
	}

	return nil
}

type Pagination struct {
	Limit  *int64
	Offset *int64
}

func (srv *Pagination) LoadProtoStruct(data *protos.PaginationOption) {
	if (data != nil) {
		srv.Limit = data.Limit
		srv.Offset = data.Offset
	}
}

func (srv *Pagination) OutProtoStruct() *protos.PaginationOption {
	result := new(protos.PaginationOption)
	result.Offset = srv.Offset
	result.Limit = srv.Limit
	return result
}

type Sort struct {
	Key  string
	Type SortType
}

type SortType = int32

const (
	SortAsc  = SortType(0)
	SortDesc = SortType(1)
)

func (srv *Sort) LoadProtoStruct(data *protos.SortOption) {
	if (data != nil) {
		srv.Key = data.Key
		srv.Type = int32(data.Type.Number())
	}
}

func (srv *Sort) OutProtoStruct() *protos.SortOption {
	result := new(protos.SortOption)
	result.Key = srv.Key
	result.Type = protos.SortOption_TypeOption(srv.Type)
	return result
}
