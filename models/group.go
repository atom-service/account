package models

// Group 用户可以属于某个组
// 组管理员可以查看管理组内的成员
type Group struct {
	ID          uint64 `db:"ID"`          // ID
	Type        string `db:"Type"`        // Type
	Name        string `db:"State"`       // Name
	CreateTime  string `db:"CreateTime"`  // 创建时间
	UpdateTime  string `db:"UpdateTime"`  // 更新时间
	Description string `db:"Description"` // 说明
}
