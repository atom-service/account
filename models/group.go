package models

import (
	"strconv"

	"github.com/grpcbrick/account/standard"
)

// Group 用户可以属于某个组
// 组管理员可以查看管理组内的成员
type Group struct {
	ID          uint32 // ID
	Name        string // Name
	Class       string // 分类
	State       string
	Description string
	DeletedTime string
	CreatedTime string
	UpdatedTime string
}

// LoadProtoStruct LoadProtoStruct
func (srv *Group) LoadProtoStruct(group *standard.Group) {
	srv.ID = group.ID
	srv.Name = group.Name
	srv.Class = group.Class
	srv.State = group.State
	srv.Description = group.Description
	srv.DeletedTime = group.DeletedTime
	srv.CreatedTime = group.CreatedTime
	srv.UpdatedTime = group.UpdatedTime
}

// LoadStringMap 从 string map 中加载数据
func (srv *Group) LoadStringMap(data map[string]string) {
	srv.Name = data["Name"]
	srv.Class = data["Class"]
	srv.State = data["State"]
	srv.Description = data["Description"]
	srv.DeletedTime = data["DeletedTime"]
	srv.CreatedTime = data["CreatedTime"]
	srv.UpdatedTime = data["UpdatedTime"]
	id, _ := strconv.ParseUint(data["ID"], 10, 32)
	srv.ID = uint32(id)
}

// OutProtoStruct OutProtoStruct
func (srv *Group) OutProtoStruct() *standard.Group {
	lable := new(standard.Group)
	lable.ID = srv.ID
	lable.Name = srv.Name
	lable.Class = srv.Class
	lable.State = srv.State
	lable.Description = srv.Description
	lable.DeletedTime = srv.DeletedTime
	lable.CreatedTime = srv.CreatedTime
	lable.UpdatedTime = srv.UpdatedTime
	return lable
}
