package model

import "github.com/grpcbrick/account/standard"

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
