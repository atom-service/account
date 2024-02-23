package models

import (
	"database/sql"

	"github.com/grpcbrick/account/standard"
)

// Label 标签
type Label struct {
	ID          sql.NullInt64
	Name        sql.NullString
	State       sql.NullString
	Value       sql.NullString
	Category    sql.NullString
	DeletedTime sql.NullTime
	CreatedTime sql.NullTime
	UpdatedTime sql.NullTime
}

// LoadProtoStruct LoadProtoStruct
func (srv *Label) LoadProtoStruct(label *standard.Label) {
	srv.ID.Scan(label.ID)
	srv.Name.Scan(label.Name)
	srv.State.Scan(label.State)
	srv.Value.Scan(label.Value)
	srv.Category.Scan(label.Category)
	srv.DeletedTime.Scan(label.DeletedTime)
	srv.CreatedTime.Scan(label.CreatedTime)
	srv.UpdatedTime.Scan(label.UpdatedTime)
}

// LoadStringMap 从 string map 中加载数据
func (srv *Label) LoadStringMap(data map[string]string) {
	srv.ID.Scan(data["ID"])
	srv.Name.Scan(data["Name"])
	srv.State.Scan(data["State"])
	srv.Value.Scan(data["Value"])
	srv.Category.Scan(data["Category"])
	srv.DeletedTime.Scan(data["DeletedTime"])
	srv.CreatedTime.Scan(data["CreatedTime"])
	srv.UpdatedTime.Scan(data["UpdatedTime"])
}

// OutProtoStruct OutProtoStruct
func (srv *Label) OutProtoStruct() *standard.Label {
	lable := new(standard.Label)
	lable.ID = srv.ID.Int64
	lable.Name = srv.Name.String
	lable.State = srv.State.String
	lable.Value = srv.Value.String
	lable.Category = srv.Category.String
	lable.DeletedTime = srv.DeletedTime.Time.String()
	lable.CreatedTime = srv.CreatedTime.Time.String()
	lable.UpdatedTime = srv.UpdatedTime.Time.String()
	return lable
}
