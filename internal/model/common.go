package model

import "github.com/atom-service/account/package/protos"

type Pagination struct {
	Limit  *int64
	Offset *int64
}

func (srv *Pagination) LoadProtoStruct(data *protos.PaginationOption) {
	srv.Limit = data.Limit
	srv.Offset = data.Offset
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
	srv.Key = data.Key
	srv.Type = int32(data.Type.Number())
}

func (srv *Sort) OutProtoStruct() *protos.SortOption {
	result := new(protos.SortOption)
	result.Key = srv.Key
	result.Type = protos.SortOption_TypeOption(srv.Type)
	return result
}
