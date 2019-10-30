package model

import (
	"strconv"

	"github.com/grpcbrick/account/standard"
)

// User 用户
type User struct {
	ID         uint64 `db:"ID"`
	Class      string `db:"Class"`
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

// LoadStringMap 从 string map 中加载数据
func (srv *User) LoadStringMap(data map[string]string) {
	srv.Class = data["Class"]
	srv.Avatar = data["Avatar"]
	srv.Username = data["Username"]
	srv.Nickname = data["Nickname"]
	srv.UpdateTime = data["UpdateTime"]
	srv.CreateTime = data["CreateTime"]
	srv.ID, _ = strconv.ParseUint(data["ID"], 10, 64)
	srv.Inviter, _ = strconv.ParseUint(data["Inviter"], 10, 64)
}

// LoadProtoStruct LoadProtoStruct
func (srv *User) LoadProtoStruct(user *standard.User) {
	srv.ID = user.ID
	srv.Class = user.Class
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
	user.Class = srv.Class
	user.Avatar = srv.Avatar
	user.Inviter = srv.Inviter
	user.Nickname = srv.Nickname
	user.Username = srv.Username
	user.Password = srv.Password
	user.CreateTime = srv.CreateTime

	return user
}
