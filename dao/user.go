package dao

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/grpcbrick/account/models"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/crypto"
	"github.com/yinxulai/goutils/sqldb"
)

const userTableName = "users"
const userHistoryTableName = "users-history"

func createUserTable() error {
	// 主表
	matserStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + userTableName + "`",
		"(",
		"`ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
		"`Class` varchar(128) NOT NULL COMMENT '账户类型',",
		"`Avatar` varchar(512) DEFAULT '' COMMENT '头像', ",
		"`Inviter` int(11) DEFAULT 0 COMMENT '邀请人',",
		"`Nickname` varchar(128) NOT NULL COMMENT '昵称',",
		"`Username` varchar(128) NOT NULL COMMENT '用户名',",
		"`Password` varchar(512) NOT NULL COMMENT '密码',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`ID`,`Nickname`,`Class`,`Username`),",
		"UNIQUE KEY `Username` (`Username`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err := matserStmp.Exec()
	if err != nil {
		return err
	}
	// 历史表
	historyStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + userHistoryTableName + "`",
		"(",
		"`Index` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Index',",
		"`ID` int(11)  COMMENT 'ID',",
		"`Class` varchar(128) COMMENT '账户类型',",
		"`Avatar` varchar(512) COMMENT '头像', ",
		"`Inviter` int(11) COMMENT '邀请人',",
		"`Nickname` varchar(128) COMMENT '昵称',",
		"`Username` varchar(128) COMMENT '用户名',",
		"`Password` varchar(512) COMMENT '密码',",
		"`DeletedTime` datetime COMMENT '删除时间',",
		"`CreatedTime` datetime COMMENT '创建时间',",
		"`UpdatedTime` datetime COMMENT '更新时间',",
		"PRIMARY KEY (`Index`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}

	return nil
}

func truncateUserTable() error {
	var err error
	masterStmp := sqldb.CreateStmt("truncate table `" + userTableName + "`")
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}
	historyStmp := sqldb.CreateStmt("truncate table `" + userHistoryTableName + "`")
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}
	return nil
}

// CreateUserHistory 对指定数据创建一条历史快照
func CreateUserHistory(id int64) error {
	var err error
	namedData := map[string]interface{}{
		"ID": id,
	}

	// 插入一条更新历史
	historyStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + userHistoryTableName + "`",
		"(ID,Class,Avatar,Inviter,Nickname,Username,Password,DeletedTime,CreatedTime,UpdatedTime)",
		"SELECT",
		"ID,Class,Avatar,Inviter,Nickname,Username,Password,DeletedTime,CreatedTime,UpdatedTime",
		"FROM",
		"`" + userTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))
	_, err = historyStmp.Exec(namedData)
	if err != nil {
		return err
	}
	return nil
}

// CountUserByID 根据 id 统计
func CountUserByID(id int64) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + userTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))

	result := struct{ Count int64 }{}
	namedData := map[string]interface{}{
		"ID": id,
	}
	err := stmp.Get(&result, namedData)
	if err != nil {
		return 0, err
	}

	return result.Count, nil
}

// QueryUsers 查询用户
func QueryUsers(page, limit int64) (totalPage, currentPage int64, users []*models.User, err error) {
	currentPage = page // 固定当前页

	// 查询数据长度
	countStmp := sqldb.CreateStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + userTableName + "`",
	}, " "))

	countResult := struct{ Count int64 }{}
	err = countStmp.Get(&countResult)
	if err != nil {
		return totalPage, currentPage, users, err
	}

	count := countResult.Count
	// 计算总页码数
	totalPage = int64(math.Ceil(float64(count) / float64(limit)))

	// 超出数据总页数
	if page > totalPage {
		// 返回当前页、空数据（当前页确实不存在数据）
		return totalPage, page, users, err
	}

	// 计算偏移
	offSet := (page - 1) * limit
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + userTableName + "`",
		"LIMIT :Limit",
		"OFFSET :Offset",
	}, " "))

	users = []*models.User{}
	namedData := map[string]interface{}{
		"Limit":  limit,
		"Offset": offSet,
	}

	err = stmp.Select(&users, namedData)
	if err != nil {
		return totalPage, currentPage, users, err
	}

	return totalPage, currentPage, users, err
}

// QueryUsersByInviter 查询用户
func QueryUsersByInviter(inviter int64, page, limit int64) (totalPage, currentPage int64, users []*models.User, err error) {
	currentPage = page // 固定当前页

	countNamedData := map[string]interface{}{
		"Inviter": inviter,
	}

	// 查询数据长度
	countStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + userTableName + "`",
		"WHERE",
		"`Inviter`=:Inviter",
	}, " "))

	countResult := struct{ Count int64 }{}
	err = countStmp.Get(&countResult, countNamedData)
	if err != nil {
		return totalPage, currentPage, users, err
	}

	count := countResult.Count
	// 计算总页码数
	totalPage = int64(math.Ceil(float64(count) / float64(limit)))

	// 超出数据总页数
	if currentPage > totalPage {
		// 返回当前页、空数据（当前页确实不存在数据）
		return totalPage, currentPage, users, err
	}

	// 计算偏移
	offSet := (page - 1) * limit
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + userTableName + "`",
		"WHERE",
		"`Inviter`=:Inviter",
		"LIMIT :Limit",
		"OFFSET :Offset",
	}, " "))

	users = []*models.User{}
	namedData := map[string]interface{}{
		"Limit":   limit,
		"Offset":  offSet,
		"Inviter": inviter,
	}

	err = stmp.Select(&users, namedData)
	if err != nil {
		return totalPage, currentPage, users, err
	}

	return totalPage, currentPage, users, err
}

// QueryUserByID 根据 id 查询
func QueryUserByID(id int64) (*models.User, error) {
	var err error
	namedData := map[string]interface{}{"ID": id}
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + userTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))

	user := new(models.User)
	err = stmp.Get(user, namedData)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// QueryUserByUsername 根据 username 查询
func QueryUserByUsername(username string) (*models.User, error) {
	var err error
	namedData := map[string]interface{}{"Username": username}
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + userTableName + "`",
		"WHERE",
		"`Username`=:Username",
	}, " "))

	user := new(models.User)
	err = stmp.Get(user, namedData)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CountUserByUsername 根据用户名统计
func CountUserByUsername(username string) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + userTableName + "`",
		"WHERE",
		"`Username`=:Username",
	}, " "))

	result := struct{ Count int64 }{}
	namedData := map[string]interface{}{
		"Username": username,
	}
	err := stmp.Get(&result, namedData)
	if err != nil {
		return 0, err
	}

	return result.Count, nil
}

// CreateUser 创建用户
func CreateUser(class, nickname, username, password string, inviter int64) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + userTableName + "`",
		"(`Class`, `Nickname`, `Username`, `Inviter`, `Password`)",
		"VALUES",
		"(:Class, :Nickname, :Username, :Inviter, :Password)",
	}, " "))

	data := map[string]interface{}{
		"Class":    class,
		"Nickname": nickname,
		"Username": username,
		"Inviter":  inviter,
		"Password": crypto.MD5Encrypt(password, config.MustGet("encrypt-password")),
	}

	result, err := stmp.Exec(data)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err

}

// UpdataUserFieldByID 根据 ID 更新用户指定字段
func UpdataUserFieldByID(id int64, field map[string]interface{}) error {
	var err error

	fieldSQL := []string{}
	for name := range field {
		fieldSQL = append(fieldSQL, fmt.Sprintf("`%s`=:%s", name, name))
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"UPDATE",
		"`" + userTableName + "`",
		"SET",
		strings.Join(fieldSQL, ","),
		"WHERE",
		"`ID`=:ID",
	}, " "))

	// 修改前创建历史
	err = CreateUserHistory(id)
	if err != nil {
		return err
	}

	// 更新
	field["ID"] = id
	_, err = stmp.Exec(field)
	return err
}

// DeleteUserByID 删除用户
func DeleteUserByID(id int64) error {
	return UpdataUserFieldByID(id, map[string]interface{}{"DeletedTime": time.Now()})
}

// UpdateUserClassByID 更新用户类型
func UpdateUserClassByID(id int64, class string) error {
	return UpdataUserFieldByID(id, map[string]interface{}{"Class": class})
}

// UpdateUserAvatarByID 更新用户头像
func UpdateUserAvatarByID(id int64, avatar string) error {
	return UpdataUserFieldByID(id, map[string]interface{}{"Avatar": avatar})
}

// UpdateUserNicknameByID 更新用户昵称
func UpdateUserNicknameByID(id int64, nickname string) error {
	return UpdataUserFieldByID(id, map[string]interface{}{"Nickname": nickname})
}

// UpdateUserInviterByID 更新用户邀请码
func UpdateUserInviterByID(id, inviter int64) error {
	return UpdataUserFieldByID(id, map[string]interface{}{"Inviter": inviter})
}

// UpdateUserPasswordByID 更新用户密码
func UpdateUserPasswordByID(id int64, password string) error {
	// 加密
	encryptPassword := crypto.MD5Encrypt(password, config.MustGet("encrypt-password"))
	return UpdataUserFieldByID(id, map[string]interface{}{"Password": encryptPassword})
}

// VerifyUserPasswordByID 验证用户密码
func VerifyUserPasswordByID(id int64, password string) (bool, error) {
	user, err := QueryUserByID(id)
	if err != nil {
		return false, err
	}
	// 加密
	encryptPassword := crypto.MD5Encrypt(password, config.MustGet("encrypt-password"))
	if !user.Password.Valid { // NULL 密码禁止使用
		return false, nil
	}

	if user.Password.String != encryptPassword {
		return false, nil
	}

	return true, nil
}
