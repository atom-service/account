package models

import (
	"database/sql"

	"github.com/grpcbrick/account/standard"
)

// Group 用户可以属于某个组
// 组管理员可以查看管理组内的成员
type Group struct {
	ID          sql.NullInt64  // ID
	Name        sql.NullString // Name
	Category       sql.NullString // 分类
	State       sql.NullString
	Description sql.NullString
	DeletedTime sql.NullTime
	CreatedTime sql.NullTime
	UpdatedTime sql.NullTime
}

// LoadProtoStruct LoadProtoStruct
func (srv *Group) LoadProtoStruct(group *standard.Group) {
	srv.ID.Scan(group.ID)
	srv.Name.Scan(group.Name)
	srv.Category.Scan(group.Category)
	srv.State.Scan(group.State)
	srv.Description.Scan(group.Description)
	srv.DeletedTime.Scan(group.DeletedTime)
	srv.CreatedTime.Scan(group.CreatedTime)
	srv.UpdatedTime.Scan(group.UpdatedTime)
}

// LoadStringMap 从 string map 中加载数据
func (srv *Group) LoadStringMap(data map[string]string) {
	srv.Name.Scan(data["ID"])
	srv.Name.Scan(data["Name"])
	srv.Category.Scan(data["Category"])
	srv.State.Scan(data["State"])
	srv.Description.Scan(data["Description"])
	srv.DeletedTime.Scan(data["DeletedTime"])
	srv.CreatedTime.Scan(data["CreatedTime"])
	srv.UpdatedTime.Scan(data["UpdatedTime"])
}

// OutProtoStruct OutProtoStruct
func (srv *Group) OutProtoStruct() *standard.Group {
	lable := new(standard.Group)
	lable.ID = srv.ID.Int64
	lable.Name = srv.Name.String
	lable.Category = srv.Category.String
	lable.State = srv.State.String
	lable.Description = srv.Description.String
	lable.DeletedTime = srv.DeletedTime.Time.String()
	lable.CreatedTime = srv.CreatedTime.Time.String()
	lable.UpdatedTime = srv.UpdatedTime.Time.String()
	return lable
}
