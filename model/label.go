package model

import (
	"strconv"

	"github.com/grpcbrick/account/standard"
)

// Label 标签
type Label struct {
	ID         uint64 `db:"ID"`
	Class      string `db:"Class"`
	State      string `db:"State"`
	Value      string `db:"Value"`
	Owner      uint64 `db:"Owner"`
	CreateTime string `db:"CreateTime"`
	UpdateTime string `db:"UpdateTime"`
}

// LoadProtoStruct LoadProtoStruct
func (srv *Label) LoadProtoStruct(label *standard.Label) {
	srv.ID = label.ID
	srv.Class = label.Class
	srv.State = label.State
	srv.Value = label.Value
	srv.CreateTime = label.CreateTime
}

// LoadStringMap 从 string map 中加载数据
func (srv *Label) LoadStringMap(data map[string]string) {
	srv.Class = data["Class"]
	srv.State = data["State"]
	srv.Value = data["Value"]
	srv.UpdateTime = data["UpdateTime"]
	srv.CreateTime = data["CreateTime"]
	srv.ID, _ = strconv.ParseUint(data["ID"], 10, 64)
}

// OutProtoStruct OutProtoStruct
func (srv *Label) OutProtoStruct() *standard.Label {
	lable := new(standard.Label)
	lable.ID = srv.ID
	lable.Class = srv.Class
	lable.State = srv.State
	lable.Value = srv.Value
	lable.CreateTime = srv.CreateTime
	return lable
}
