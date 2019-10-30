package dao

import (
	"strconv"
	"strings"

	"github.com/yinxulai/goutils/easysql"
)

func createUserTable() {
	conn := easysql.GetConn()
	defer conn.Close()

	result, err := conn.ExecSQL(
		strings.Join([]string{
			" CREATE TABLE IF NOT EXISTS `users`(",
			" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
			" `Class` varchar(128) NOT NULL COMMENT '账户类型',",
			" `Avatar` varchar(512) DEFAULT '' COMMENT '头像', ",
			" `Inviter` int(11) DEFAULT 0 COMMENT '邀请人',",
			" `Nickname` varchar(128) NOT NULL COMMENT '昵称',",
			" `Username` varchar(128) NOT NULL COMMENT '用户名',",
			" `Password` varchar(512) NOT NULL COMMENT '密码',",
			" `CreateTime` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
			" `UpdateTime` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			" PRIMARY KEY (`ID`,`Nickname`,`Username`),",
			" UNIQUE KEY `Username` (`Username`)",
			" ) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
		}, "",
		),
	)
	print(result, err)
}

// 根据 id 统计
func countUserByID(id uint64) (int, error) {
	conn := easysql.GetConn()
	defer conn.Close()

	idstr := strconv.FormatUint(id, 10)
	cond := map[string]string{"ID": idstr}
	queryField := []string{"count(*) as count"}
	result, err := conn.Select("users", queryField).Where(cond).QueryRow()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(result["count"])
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 根据用户名统计
func countUserByUsername(username string) (int, error) {
	conn := easysql.GetConn()
	defer conn.Close()

	queryField := []string{"count(*) as count"}
	cond := map[string]string{"Username": username}
	result, err := conn.Select("users", queryField).Where(cond).QueryRow()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(result["count"])
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 创建用户
func createUser(class, nickname, username, password string) error {
	conn := easysql.GetConn()
	defer conn.Close()

	cond := map[string]string{
		"Class":    class,
		"Nickname": nickname,
		"Username": username,
		"Password": password,
	}

	_, err := conn.Insert("users", cond)
	if err != nil {
		return err
	}

	return nil
}

// 更新用户类型
func updateUserClassByID(id uint64, class string) error {
	return updataUserFieldByID(id, map[string]string{"Class": class})
}

// 更新用户头像
func updateUserAvatarByID(id uint64, avatar string) error {
	return updataUserFieldByID(id, map[string]string{"Avatar": avatar})
}

// 更新用户昵称
func updateUserNicknameByID(id uint64, nickname string) error {
	return updataUserFieldByID(id, map[string]string{"Nickname": nickname})
}

// 更新用户邀请码
func updateUserInviterByID(id uint64, inviter string) error {
	return updataUserFieldByID(id, map[string]string{"Inviter": inviter})
}

// 更新用户密码
func updateUserPasswordByID(id uint64, password string) error {
	return updataUserFieldByID(id, map[string]string{"Password": password})
}

// 根据 ID 更新用户指定字段
func updataUserFieldByID(id uint64, field map[string]string) error {
	conn := easysql.GetConn()
	defer conn.Close()

	cond := map[string]string{"ID": strconv.FormatUint(id, 10)}
	_, err := conn.Where(cond).Update("users", field)
	return err
}
