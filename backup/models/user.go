package models

import (
	"database/sql"

	"github.com/grpcbrick/account/standard"
)

// User 用户
type User struct {
	ID          sql.NullInt64
	Inviter     sql.NullInt64
	Avatar      sql.NullString
	Category    sql.NullString
	Nickname    sql.NullString
	Username    sql.NullString
	Password    sql.NullString
	DeletedTime sql.NullTime
	CreatedTime sql.NullTime
	UpdatedTime sql.NullTime
}

// LoadStringMap 从 string map 中加载数据
func (srv *User) LoadStringMap(data map[string]string) {
	srv.ID.Scan(data["ID"])
	srv.Category.Scan(data["Category"])
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
	srv.Avatar.Scan(user.Avatar)
	srv.Inviter.Scan(user.Inviter)
	srv.Nickname.Scan(user.Nickname)
	srv.Username.Scan(user.Username)
	srv.Password.Scan(user.Password)
	srv.Category.Scan(user.Category)
	srv.DeletedTime.Scan(user.DeletedTime)
	srv.CreatedTime.Scan(user.CreatedTime)
	srv.UpdatedTime.Scan(user.UpdatedTime)
}

// OutProtoStruct OutProtoStruct
func (srv *User) OutProtoStruct() *standard.User {
	user := new(standard.User)
	user.ID = srv.ID.Int64
	user.Password = "Privacyfield"
	user.Avatar = srv.Avatar.String
	user.Inviter = srv.Inviter.Int64
	user.Nickname = srv.Nickname.String
	user.Username = srv.Username.String
	user.Category = srv.Category.String
	user.DeletedTime = srv.DeletedTime.Time.String()
	user.CreatedTime = srv.CreatedTime.Time.String()
	user.UpdatedTime = srv.UpdatedTime.Time.String()

	return user
}
