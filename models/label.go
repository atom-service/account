package model

import (
	"strconv"

	"github.com/grpcbrick/account/standard"
)

// Label 标签
type Label struct {
	ID          uint64 `db:"ID"`
	Name        string `db:"Name"`
	Class       string `db:"Class"`
	State       string `db:"State"`
	Value       string `db:"Value"`
	DeletedTime string `db:"DeletedTime"`
	CreatedTime string `db:"CreatedTime"`
	UpdatedTime string `db:"UpdatedTime"`
}

// LoadProtoStruct LoadProtoStruct
func (srv *Label) LoadProtoStruct(label *standard.Label) {
	srv.ID = label.ID
	srv.Name = label.Name
	srv.Class = label.Class
	srv.State = label.State
	srv.Value = label.Value
	srv.DeletedTime = label.DeletedTime
	srv.CreatedTime = label.CreatedTime
	srv.UpdatedTime = label.UpdatedTime
}

// LoadStringMap 从 string map 中加载数据
func (srv *Label) LoadStringMap(data map[string]string) {
	srv.Name = data["Name"]
	srv.Class = data["Class"]
	srv.State = data["State"]
	srv.Value = data["Value"]
	srv.DeletedTime = data["DeletedTime"]
	srv.CreatedTime = data["CreatedTime"]
	srv.UpdatedTime = data["UpdatedTime"]
	srv.ID, _ = strconv.ParseUint(data["ID"], 10, 64)
}

// OutProtoStruct OutProtoStruct
func (srv *Label) OutProtoStruct() *standard.Label {
	lable := new(standard.Label)
	lable.ID = srv.ID
	lable.Name = srv.Name
	lable.Class = srv.Class
	lable.State = srv.State
	lable.Value = srv.Value
	lable.DeletedTime = srv.DeletedTime
	lable.CreatedTime = srv.CreatedTime
	lable.UpdatedTime = srv.UpdatedTime
	return lable
}
