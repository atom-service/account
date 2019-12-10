package models

import (
	"database/sql"

	"github.com/grpcbrick/account/standard"
)

// User 用户
type User struct {
	ID          sql.NullInt64
	Class       sql.NullString
	Avatar      sql.NullString
	Inviter     sql.NullInt64
	Nickname    sql.NullString
	Username    sql.NullString
	Password    sql.NullString
	DeletedTime sql.NullTime
	CreatedTime sql.NullTime
	UpdatedTime sql.NullTime
}

// SetPassword SetPassword
func (srv *User) SetPassword(password string) {
	srv.Password.Scan(password)
}

// EqualPassword EqualPassword
func (srv *User) EqualPassword(target string) bool {
	return false
}

// LoadStringMap 从 string map 中加载数据
func (srv *User) LoadStringMap(data map[string]string) {
	srv.ID.Scan(data["ID"])
	srv.Class.Scan(data["Class"])
	srv.Avatar.Scan(data["Avatar"])
	srv.Inviter.Scan(data["Inviter"])
	srv.Username.Scan(data["Username"])
	srv.Nickname.Scan(data["Nickname"])
	srv.DeletedTime.Scan(data["DeletedTime"])
	srv.CreatedTime.Scan(data["CreatedTime"])
	srv.UpdatedTime.Scan(data["UpdatedTime"])
}

// LoadProtoStruct LoadProtoStruct
func (srv *User) LoadProtoStruct(user *standard.User) {
	srv.ID.Scan(user.ID)
	srv.Class.Scan(user.Class)
	srv.Avatar.Scan(user.Avatar)
	srv.Inviter.Scan(user.Inviter)
	srv.Nickname.Scan(user.Nickname)
	srv.Username.Scan(user.Username)
	srv.Password.Scan(user.Password)
	srv.DeletedTime.Scan(user.DeletedTime)
	srv.CreatedTime.Scan(user.CreatedTime)
	srv.UpdatedTime.Scan(user.UpdatedTime)
}

// OutProtoStruct OutProtoStruct
func (srv *User) OutProtoStruct() *standard.User {
	user := new(standard.User)
	user.ID = srv.ID.Int64
	user.Class = srv.Class.String
	user.Avatar = srv.Avatar.String
	user.Inviter = srv.Inviter.Int64
	user.Nickname = srv.Nickname.String
	user.Username = srv.Username.String
	user.Password = srv.Password.String
	user.DeletedTime = srv.DeletedTime.Time.String()
	user.CreatedTime = srv.CreatedTime.Time.String()
	user.UpdatedTime = srv.UpdatedTime.Time.String()

	return user
}
