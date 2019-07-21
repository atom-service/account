package models

import "github.com/yinxulai/grpc-module-account/standard"

// User 用户
type User struct {
	ID         uint64 `db:"ID"`
	Type       string `db:"Type"`
	Avatar     string `db:"Avatar"`
	Inviter    uint64 `db:"Inviter"`
	Nickname   string `db:"Nickname"`
	Username   string `db:"Username"`
	Password   string `db:"Password"`
	CreateTime string `db:"CreateTime"`
	UpdateTime string `db:"UpdateTime"`
}

// SetPassword SetPassword
func (srv *User) SetPassword(password string) {
	srv.Password = password
}

// EqualPassword EqualPassword
func (srv *User) EqualPassword(target string) bool {
	return false
}

// LoadProtoStruct LoadProtoStruct
func (srv *User) LoadProtoStruct(user *standard.User) {
	srv.ID = user.ID
	srv.Type = user.Type
	srv.Avatar = user.Avatar
	srv.Inviter = user.Inviter
	srv.Nickname = user.Nickname
	srv.Username = user.Username
	srv.Password = user.Password
	srv.CreateTime = user.CreateTime
}

// OutProtoStruct OutProtoStruct
func (srv *User) OutProtoStruct() *standard.User {
	user := new(standard.User)

	user.ID = srv.ID
	user.Type = srv.Type
	user.Avatar = srv.Avatar
	user.Inviter = srv.Inviter
	user.Nickname = srv.Nickname
	user.Username = srv.Username
	user.Password = srv.Password
	user.CreateTime = srv.CreateTime

	return user
}
