package model

// Group 用户可以属于某个组
// 组管理员可以查看管理组内的成员
// type Group struct {
// 	ID          uint64 `db:"ID"`          // ID
// 	Type        string `db:"Type"`        // Type
// 	Name        string `db:"State"`       // Name
// 	CreateTime  string `db:"CreateTime"`  // 创建时间
// 	UpdateTime  string `db:"UpdateTime"`  // 更新时间
// 	Description string `db:"Description"` // 说明
// }

// // LoadProtoStruct LoadProtoStruct
// func (srv *Group) LoadProtoStruct(group *standard.Group) {
// 	srv.ID = group.ID
// 	srv.Type = group.Type
// 	srv.Name = group.Name
// 	srv.Description = group.Description
// }

// // OutProtoStruct OutProtoStruct
// func (srv *Group) OutProtoStruct() *standard.Group {
// 	lable := new(standard.Group)
// 	lable.ID = srv.ID
// 	lable.Name = srv.Name
// 	lable.Type = srv.Type
// 	lable.UpdateTime = srv.UpdateTime
// 	lable.Description = srv.Description
// 	lable.Description = srv.Description
// 	return lable
// }
