package models

import (
	"database/sql"

	"github.com/grpcbrick/account/standard"
)

// Label 标签
type Label struct {
	ID          sql.NullInt64
	Name        sql.NullString
	Category       sql.NullString
	State       sql.NullString
	Value       sql.NullString
	DeletedTime sql.NullTime
	CreatedTime sql.NullTime
	UpdatedTime sql.NullTime
}

// LoadProtoStruct LoadProtoStruct
func (srv *Label) LoadProtoStruct(label *standard.Label) {
	srv.ID.Scan(label.ID)
	srv.Name.Scan(label.Name)
	srv.Category.Scan(label.Category)
	srv.State.Scan(label.State)
	srv.Value.Scan(label.Value)
	srv.DeletedTime.Scan(label.DeletedTime)
	srv.CreatedTime.Scan(label.CreatedTime)
	srv.UpdatedTime.Scan(label.UpdatedTime)
}

// LoadStringMap 从 string map 中加载数据
func (srv *Label) LoadStringMap(data map[string]string) {
	srv.ID.Scan(data["ID"])
	srv.Name.Scan(data["Name"])
	srv.Category.Scan(data["Category"])
	srv.State.Scan(data["State"])
	srv.Value.Scan(data["Value"])
	srv.DeletedTime.Scan(data["DeletedTime"])
	srv.CreatedTime.Scan(data["CreatedTime"])
	srv.UpdatedTime.Scan(data["UpdatedTime"])
}

// OutProtoStruct OutProtoStruct
func (srv *Label) OutProtoStruct() *standard.Label {
	lable := new(standard.Label)
	lable.ID = srv.ID.Int64
	lable.Name = srv.Name.String
	lable.Category = srv.Category.String
	lable.State = srv.State.String
	lable.Value = srv.Value.String
	lable.DeletedTime = srv.DeletedTime.Time.String()
	lable.CreatedTime = srv.CreatedTime.Time.String()
	lable.UpdatedTime = srv.UpdatedTime.Time.String()
	return lable
}
