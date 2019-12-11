package dao

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/grpcbrick/account/models"
	"github.com/yinxulai/goutils/easysql"
	"github.com/yinxulai/goutils/sqldb"
)

const groupTableName = "group"
const groupMappingUserTableName = "group-mapping"

func truncateGroupTable() error {
	stmp := sqldb.CreateStmt("truncate table `" + groupTableName + "`")
	_, err := stmp.Exec()
	return err
}

func createGroupTable() error {
	stmp := sqldb.CreateStmt(strings.Join([]string{
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
		")ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " ",
	))
	_, err := stmp.Exec()
	return err
}

// CreateGroup 创建组
func CreateGroup(name, class, state, description string) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + groupTableName + "`",
		"(Name, Class, State, Description)",
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
		"Name=:Name",
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
		"ID=:ID",
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
		"SELECT COUNT(*) as Count FROM",
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
		" SELECT * FROM",
		"`" + groupTableName + "`",
		" LIMIT :Limit",
		" OFFSET :Offset",
	}, ""))

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
	conn := easysql.GetConn()

	idstr := strconv.FormatUint(uint64(id), 10)
	cond := map[string]string{"ID": idstr}
	result, err := conn.Select(groupTableName, nil).Where(cond).QueryRow()
	if err != nil {
		return nil, err
	}

	group := new(models.Group)
	group.LoadStringMap(result)
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
	// TODO: 考虑如何记录修改记录

	fieldSQL := []string{}
	for name := range field {
		fieldSQL = append(fieldSQL, fmt.Sprintf("`%s`=:%s", name, name))
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"UPDATE",
		"`" + groupTableName + "`",
		"SET",
		strings.Join(fieldSQL, ","),
		fmt.Sprintf("WHERE ID = %d", id),
	}, " "))

	_, err := stmp.Exec(field)
	return err
}

func truncateGroupMappingTable() error {
	stmp := sqldb.CreateStmt("truncate table `" + groupMappingUserTableName + "`")
	_, err := stmp.Exec()
	return err
}

func createGroupMappingTable() error {
	stmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + groupMappingUserTableName + "` (",
		"`Group` int(11) NOT NULL COMMENT '组 ID',",
		"`User` int(11) NOT NULL COMMENT '用户 ID',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`Group`,`Owner`)",
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err := stmp.Exec()
	return err
}

// 关系历史表
// 关系操作先复制当前记录到历史表、再更新主表
func createGroupMappingHistoreTable() error {
	//
	stmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + groupMappingUserTableName + "History` (",
		"`Group` int(11) NOT NULL COMMENT '组 ID',",
		"`User` int(11) NOT NULL COMMENT '用户 ID',",
		"`UpdatedTime` datetime COMMENT '更新时间'",
		"`CreatedTime` datetime COMMENT '创建时间',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err := stmp.Exec()
	return err
}

// RemoveUserFromGroupByID 从组里移除用户
func RemoveUserFromGroupByID(group, user int64) error {
	// TODO: 重新设计：删除数据时、更新删除时间
	// 同时将该条数据移动到对应的 history 表中

	tx, err := sqldb.GetDB().Beginx()
	if err != nil {
		return err
	}

	// 复制
	stmp, err := tx.PrepareNamed(strings.Join([]string{
		"UPDATE",
		"`" + groupMappingUserTableName + "`",
		"SET",
		"`DeletedTime`=:DeletedTime",
		"WHERE",
		"`User`=:User",
		"AND",
		"`Group`=:Group",
		"AND",
		"DeletedTime IS NOT NULL", // 删除时间已存在说明是旧数据
	}, " "))

	if err != nil {
		return err
	}

	namedData := map[string]interface{}{
		"User":        user,
		"Group":       group,
		"DeletedTime": time.Now(),
	}

	// 更新历史
	stmp, err := tx.PrepareNamed(strings.Join([]string{
		"UPDATE",
		"`" + groupMappingUserTableName + "`",
		"SET",
		"`DeletedTime`=:DeletedTime",
		"WHERE",
		"`User`=:User",
		"AND",
		"`Group`=:Group",
		"AND",
		"DeletedTime IS NOT NULL", // 删除时间已存在说明是旧数据
	}, " "))
	if err != nil {
		return err
	}

	_, err := stmp.Exec(namedData)
	return err
}

// AddUserToGroupByID 添加用户进组
func AddUserToGroupByID(group, user int64) error {
	conn := easysql.GetConn()

	data := map[string]string{
		"User":  strconv.FormatUint(uint64(user), 10),
		"Group": strconv.FormatUint(uint64(group), 10),
	}

	_, err := conn.Insert(groupMappingUserTableName, data)
	return err
}

// IsAlreadyInGroup 是否已存在关联
func IsAlreadyInGroup(group, user int64) (bool, error) {
	conn := easysql.GetConn()

	cond := map[string]string{
		"User":  strconv.FormatUint(uint64(user), 10),
		"Group": strconv.FormatUint(uint64(group), 10),
	}
	queryField := []string{"count(*) as count"}
	result, err := conn.Select(groupMappingUserTableName, queryField).Where(cond).QueryRow()
	if err != nil {
		return false, err
	}
	count, err := strconv.Atoi(result["count"])
	if err != nil {
		return false, err
	}

	if count <= 0 {
		return false, nil
	}

	return true, nil
}
