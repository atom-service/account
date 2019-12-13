package dao

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/grpcbrick/account/models"
	"github.com/yinxulai/goutils/sqldb"
)

const groupTableName = "group"
const groupHistoryTableName = "group-history"
const groupMappingUserTableName = "group-mapping"
const groupMappingUserHistoryTableName = "group-mapping-history"

func createGroupTable() error {
	var err error
	// 主表
	masterStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS",
		"`" + groupTableName + "`",
		"(",
		"`ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
		"`Name` varchar(128) NOT NULL COMMENT '名称',",
		"`Class` varchar(128) NOT NULL COMMENT '类型',",
		"`State` varchar(128) DEFAULT '' COMMENT '状态',",
		"`Description` varchar(512) DEFAULT '' COMMENT '简介',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`ID`,`Name`,`Class`,`State`),",
		"UNIQUE KEY `Name` (`Name`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}

	// 历史表
	historyStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS",
		"`" + groupHistoryTableName + "`",
		"(",
		"`Index` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Index',",
		"`ID` int(11) COMMENT 'ID',",
		"`Name` varchar(128) COMMENT '名称',",
		"`Class` varchar(128) COMMENT '类型',",
		"`State` varchar(128) COMMENT '状态',",
		"`Description` varchar(512) COMMENT '简介',",
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

func truncateGroupTable() error {
	var err error
	masterStmp := sqldb.CreateStmt("truncate table `" + groupTableName + "`")
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}
	historyStmp := sqldb.CreateStmt("truncate table `" + groupHistoryTableName + "`")
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}
	return nil
}

// CreateGroupHistory 对指定数据创建一条历史快照
func CreateGroupHistory(id int64) error {
	var err error
	namedData := map[string]interface{}{
		"ID": id,
	}

	// 插入一条更新历史
	historyStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + groupHistoryTableName + "`",
		"(`ID`,`Name`,`Class`,`State`,`Description`,`DeletedTime`,`CreatedTime`,`UpdatedTime`)",
		"SELECT",
		"`ID`,`Name`,`Class`,`State`,`Description`,`DeletedTime`,`CreatedTime`,`UpdatedTime`",
		"FROM",
		"`" + groupTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))
	_, err = historyStmp.Exec(namedData)
	if err != nil {
		return err
	}
	return nil
}

// CreateGroup 创建组
func CreateGroup(name, class, state, description string) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + groupTableName + "`",
		"(`Name`, `Class`, `State`, `Description`)",
		"VALUES",
		"(:Name, :Class, :State, :Description)",
	}, " "))

	data := map[string]interface{}{
		"Name":        name,
		"Class":       class,
		"State":       state,
		"Description": description,
	}
	result, err := stmp.Exec(data)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

// CountGroupByName 根据 name 统计组的数量
func CountGroupByName(name string) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + groupTableName + "`",
		"WHERE",
		"`Name`=:Name",
	}, " "))

	result := struct{ Count int64 }{}
	namedData := map[string]interface{}{
		"Name": name,
	}
	err := stmp.Get(&result, namedData)
	if err != nil {
		return 0, err
	}

	return result.Count, nil
}

// CountGroupByID 根据 id 统计
func CountGroupByID(id int64) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + groupTableName + "`",
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

// QueryGroups 查询组
func QueryGroups(page, limit int64) (totalPage, currentPage int64, groups []*models.Group, err error) {
	currentPage = page // 固定当前页

	// 查询数据长度
	countStmp := sqldb.CreateStmt(strings.Join([]string{
		"SELECT COUNT(*) as `Count` FROM",
		"`" + groupTableName + "`",
	}, " "))

	countResult := struct{ Count int64 }{}
	err = countStmp.Get(&countResult)
	if err != nil {
		return totalPage, currentPage, groups, err
	}

	count := countResult.Count
	// 计算总页码数
	totalPage = int64(math.Ceil(float64(count) / float64(limit)))

	// 超出数据总页数
	if page > totalPage {
		// 返回当前页、空数据（当前页确实不存在数据）
		return totalPage, page, groups, err
	}

	// 计算偏移
	offSet := (page - 1) * limit
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + groupTableName + "`",
		"LIMIT :Limit",
		"OFFSET :Offset",
	}, " "))

	groups = []*models.Group{}
	namedData := map[string]interface{}{
		"Limit":  limit,
		"Offset": offSet,
	}

	err = stmp.Select(&groups, namedData)
	if err != nil {
		return totalPage, currentPage, groups, err
	}

	return totalPage, currentPage, groups, err
}

// QueryGroupByID 根据 id 查询
func QueryGroupByID(id int64) (*models.Group, error) {
	var err error
	namedData := map[string]interface{}{"ID": id}
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + groupTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))

	group := new(models.Group)
	err = stmp.Get(group, namedData)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// DeleteGroupByID 删除标签
func DeleteGroupByID(id int64) error {
	return updataGroupFieldByID(id, map[string]interface{}{"DeletedTime": time.Now()})
}

// UpdateGroupNameByID 更新标签类型
func UpdateGroupNameByID(id int64, name string) error {
	return updataGroupFieldByID(id, map[string]interface{}{"Name": name})
}

// UpdateGroupClassByID 更新标签状态
func UpdateGroupClassByID(id int64, class string) error {
	return updataGroupFieldByID(id, map[string]interface{}{"Class": class})
}

// UpdateGroupStateByID 更新标签值
func UpdateGroupStateByID(id int64, class string) error {
	return updataGroupFieldByID(id, map[string]interface{}{"State": class})
}

// UpdateGroupDescriptionByID 更新标签值
func UpdateGroupDescriptionByID(id int64, description string) error {
	return updataGroupFieldByID(id, map[string]interface{}{"Description": description})
}

// 根据 ID 更新标签
func updataGroupFieldByID(id int64, field map[string]interface{}) error {
	var err error

	fieldSQL := []string{}
	for name := range field {
		fieldSQL = append(fieldSQL, fmt.Sprintf("`%s`=:%s", name, name))
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"UPDATE",
		"`" + groupTableName + "`",
		"SET",
		strings.Join(fieldSQL, ","),
		"WHERE",
		"`ID`=:ID",
	}, " "))

	// 修改前创建历史
	err = CreateGroupHistory(id)
	if err != nil {
		return err
	}

	// 更新
	field["ID"] = id
	_, err = stmp.Exec(field)
	return err
}

// 关系相关表

func createGroupMappingTable() error {
	// 主表
	matserStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + groupMappingUserTableName + "`",
		"(",
		"`ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
		"`Group` int(11) NOT NULL COMMENT '组 ID',",
		"`User` int(11) NOT NULL COMMENT '用户 ID',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`ID`,`Group`,`User`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err := matserStmp.Exec()
	if err != nil {
		return err
	}
	// 历史表
	historyStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + groupMappingUserHistoryTableName + "`",
		"(",
		"`Index` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Index',",
		"`ID` int(11) COMMENT 'ID',",
		"`Group` int(11) COMMENT '组 ID',",
		"`User` int(11) COMMENT '用户 ID',",
		"`UpdatedTime` datetime COMMENT '更新时间',",
		"`CreatedTime` datetime COMMENT '创建时间',",
		"`DeletedTime` datetime COMMENT '删除时间',",
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

func truncateGroupMappingTable() error {
	var err error
	masterStmp := sqldb.CreateStmt("truncate table `" + groupMappingUserTableName + "`")
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}
	historyStmp := sqldb.CreateStmt("truncate table `" + groupMappingUserHistoryTableName + "`")
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}
	return nil
}

// RemoveUserFromGroupByID 从组里移除用户
func RemoveUserFromGroupByID(group, user int64) error {
	var err error

	namedData := map[string]interface{}{
		"User":        user,
		"Group":       group,
		"DeletedTime": time.Now(),
	}

	updateStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"UPDATE",
		"`" + groupMappingUserTableName + "`",
		"SET",
		"`DeletedTime`=:DeletedTime", // 更新删除时间
		"WHERE",
		"`User`=:User",
		"AND",
		"`Group`=:Group",
	}, " "))

	// 先保存一条历史快照
	err = CreateGroupMappingHistory(group, user)
	if err != nil {
		return err
	}

	// 更新主标数据
	_, err = updateStmp.Exec(namedData)
	return err
}

// AddUserToGroupByID 添加用户进组
func AddUserToGroupByID(group, user int64) error {
	var err error

	namedData := map[string]interface{}{
		"User":  user,
		"Group": group,
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + groupMappingUserTableName + "`",
		"(`User`, `Group`)",
		"VALUES",
		"(:User, :Group)",
	}, " "))

	_, err = stmp.Exec(namedData)
	if err != nil {
		return err
	}

	return nil
}

// IsAlreadyInGroup 是否已存在关联
func IsAlreadyInGroup(group, user int64) (bool, error) {
	namedData := map[string]interface{}{
		"User":  user,
		"Group": user,
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + groupMappingUserTableName + "`",
		"WHERE",
		"`User`=:User",
		"AND",
		"`Group`=:Group",
	}, " "))

	result := struct{ Count int64 }{}
	err := stmp.Get(&result, namedData)
	if err != nil {
		return false, err
	}

	if result.Count <= 0 {
		return false, nil
	}

	return true, nil
}

// CreateGroupMappingHistory 对指定数据创建一条历史快照
func CreateGroupMappingHistory(group, user int64) error {

	namedData := map[string]interface{}{
		"User":  user,
		"Group": group,
	}

	// 插入一条更新历史
	historyStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + groupMappingUserHistoryTableName + "`",
		"(`ID`,`Group`,`User`,`UpdatedTime`,`CreatedTime`,`DeletedTime`)",
		"SELECT",
		"`ID`,Group,`User`,`UpdatedTime`,`CreatedTime`,`DeletedTime`",
		"FROM",
		"`" + groupMappingUserTableName + "`",
		"WHERE",
		"`User`=:User",
		"AND",
		"`Group`=:Group",
	}, " "))
	_, err := historyStmp.Exec(namedData)
	if err != nil {
		return err
	}
	return nil
}
