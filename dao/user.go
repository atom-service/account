package dao

import (
	"strconv"
	"strings"
	"time"

	"github.com/grpcbrick/account/model"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/crypto"
	"github.com/yinxulai/goutils/easysql"
)

const userTableName = "users"

func createUserTable() error {
	conn := easysql.GetConn()
	defer conn.Close()

	_, err := conn.ExecSQL(
		strings.Join([]string{
			" CREATE TABLE IF NOT EXISTS `" + userTableName + "`(",
			" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
			" `Class` varchar(128) NOT NULL COMMENT '账户类型',",
			" `Avatar` varchar(512) DEFAULT '' COMMENT '头像', ",
			" `Inviter` int(11) DEFAULT 0 COMMENT '邀请人',",
			" `Nickname` varchar(128) NOT NULL COMMENT '昵称',",
			" `Username` varchar(128) NOT NULL COMMENT '用户名',",
			" `Password` varchar(512) NOT NULL COMMENT '密码',",
			" `DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
			" `CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
			" `UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			" PRIMARY KEY (`ID`,`Nickname`,`Username`),",
			" UNIQUE KEY `Username` (`Username`)",
			" ) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
		}, "",
		),
	)
	return err
}

// CountUserByID 根据 id 统计
func CountUserByID(id uint64) (int, error) {
	conn := easysql.GetConn()
	defer conn.Close()

	idstr := strconv.FormatUint(id, 10)
	cond := map[string]string{"ID": idstr}
	queryField := []string{"count(*) as count"}
	result, err := conn.Select(userTableName, queryField).Where(cond).QueryRow()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(result["count"])
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryUserByID 根据 id 查询
func QueryUserByID(id uint64) (*model.User, error) {
	conn := easysql.GetConn()
	defer conn.Close()

	idstr := strconv.FormatUint(id, 10)
	cond := map[string]string{"ID": idstr}
	result, err := conn.Select(userTableName, nil).Where(cond).QueryRow()
	if err != nil {
		return nil, err
	}

	user := new(model.User)
	user.LoadStringMap(result)
	return user, nil
}

// QueryUserByUsername 根据 id 查询
func QueryUserByUsername(username string) (*model.User, error) {
	conn := easysql.GetConn()
	defer conn.Close()

	cond := map[string]string{"Username": username}
	result, err := conn.Select(userTableName, nil).Where(cond).QueryRow()
	if err != nil {
		return nil, err
	}

	user := new(model.User)
	user.LoadStringMap(result)
	return user, nil
}

// CountUserByUsername 根据用户名统计
func CountUserByUsername(username string) (int, error) {
	conn := easysql.GetConn()
	defer conn.Close()

	queryField := []string{"count(*) as count"}
	cond := map[string]string{"Username": username}
	result, err := conn.Select(userTableName, queryField).Where(cond).QueryRow()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(result["count"])
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CreateUser 创建用户
func CreateUser(class, nickname, username, password string) error {
	conn := easysql.GetConn()
	defer conn.Close()

	cond := map[string]string{
		"Class":    class,
		"Nickname": nickname,
		"Username": username,
		"Password": crypto.MD5Encrypt(password, config.MustGet("encrypt-password")),
	}

	_, err := conn.Insert("users", cond)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserByID 删除用户
func DeleteUserByID(id uint64, class string) error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	return updataUserFieldByID(id, map[string]string{"DeletedTime": nowTime})
}

// UpdateUserClassByID 更新用户类型
func UpdateUserClassByID(id uint64, class string) error {
	return updataUserFieldByID(id, map[string]string{"Class": class})
}

// UpdateUserAvatarByID 更新用户头像
func UpdateUserAvatarByID(id uint64, avatar string) error {
	return updataUserFieldByID(id, map[string]string{"Avatar": avatar})
}

// UpdateUserNicknameByID 更新用户昵称
func UpdateUserNicknameByID(id uint64, nickname string) error {
	return updataUserFieldByID(id, map[string]string{"Nickname": nickname})
}

// UpdateUserInviterByID 更新用户邀请码
func UpdateUserInviterByID(id uint64, inviter string) error {
	return updataUserFieldByID(id, map[string]string{"Inviter": inviter})
}

// UpdateUserPasswordByID 更新用户密码
func UpdateUserPasswordByID(id uint64, password string) error {
	// 加密
	encryptPassword := crypto.MD5Encrypt(password, config.MustGet("encrypt-password"))
	return updataUserFieldByID(id, map[string]string{"Password": encryptPassword})
}

// 根据 ID 更新用户指定字段
func updataUserFieldByID(id uint64, field map[string]string) error {
	conn := easysql.GetConn()
	defer conn.Close()

	cond := map[string]string{"ID": strconv.FormatUint(id, 10)}
	_, err := conn.Where(cond).Update(userTableName, field)
	return err
}
