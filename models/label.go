package models

import (
	"strconv"

	"github.com/grpcbrick/account/standard"
)

// Label 标签
type Label struct {
	ID          uint32
	Name        string
	Class       string
	State       string
	Value       string
	DeletedTime string
	CreatedTime string
	UpdatedTime string
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
	id, _ := strconv.ParseUint(data["ID"], 10, 32)
	srv.ID = uint32(id)
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
