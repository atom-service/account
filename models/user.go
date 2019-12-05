package models

import (
	"strconv"

	"github.com/grpcbrick/account/standard"
)

// User 用户
type User struct {
	ID          uint32 `db:"ID"`
	Class       string `db:"Class"`
	Avatar      string `db:"Avatar"`
	Inviter     uint32 `db:"Inviter"`
	Nickname    string `db:"Nickname"`
	Username    string `db:"Username"`
	Password    string `db:"Password"`
	DeletedTime string `db:"DeletedTime"`
	CreatedTime string `db:"CreatedTime"`
	UpdatedTime string `db:"UpdatedTime"`
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
	srv.DeletedTime = data["DeletedTime"]
	srv.CreatedTime = data["CreatedTime"]
	srv.UpdatedTime = data["UpdatedTime"]
	id, _ := strconv.ParseUint(data["ID"], 10, 64)
	inviter, _ := strconv.ParseUint(data["Inviter"], 10, 64)
	srv.ID = uint32(id)
	srv.Inviter = uint32(inviter)
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
	srv.DeletedTime = user.DeletedTime
	srv.CreatedTime = user.CreatedTime
	srv.UpdatedTime = user.UpdatedTime
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
	user.DeletedTime = srv.DeletedTime
	user.CreatedTime = srv.CreatedTime
	user.UpdatedTime = srv.UpdatedTime

	return user
}
