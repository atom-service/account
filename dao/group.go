package dao

import (
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
	stmp := sqldb.CreateStmt("truncate table `?`")
	_, err := stmp.Exec(groupTableName)
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

	result := []*models.Group{}
	namedData := map[string]interface{}{
		"Limit":  limit,
		"Offset": offSet,
	}

	err = stmp.Select(&result, namedData)
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
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	return updataGroupFieldByID(id, map[string]string{"DeletedTime": nowTime})
}

// UpdateGroupNameByID 更新标签类型
func UpdateGroupNameByID(id int64, name string) error {
	return updataGroupFieldByID(id, map[string]string{"Name": name})
}

// UpdateGroupClassByID 更新标签状态
func UpdateGroupClassByID(id int64, class string) error {
	return updataGroupFieldByID(id, map[string]string{"Class": class})
}

// UpdateGroupStateByID 更新标签值
func UpdateGroupStateByID(id int64, class string) error {
	return updataGroupFieldByID(id, map[string]string{"State": class})
}

// UpdateGroupDescriptionByID 更新标签值
func UpdateGroupDescriptionByID(id int64, description string) error {
	return updataGroupFieldByID(id, map[string]string{"Description": description})
}

// 根据 ID 更新标签
func updataGroupFieldByID(id int64, field map[string]string) error {
	conn := easysql.GetConn()

	cond := map[string]string{"ID": strconv.FormatUint(uint64(id), 10)}
	_, err := conn.Where(cond).Update(groupTableName, field)
	return err
}

func truncateGroupMappingTable() error {
	conn := easysql.GetConn()

	_, err := conn.ExecSQL("truncate table `" + groupMappingUserTableName + "`")
	return err
}

func createGroupMappingTable() error {
	conn := easysql.GetConn()

	_, err := conn.ExecSQL(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + groupMappingUserTableName + "` (",
		"`Group` int(11) NOT NULL COMMENT '组 ID',",
		"`Owner` int(11) NOT NULL COMMENT '所属者 ID',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`Group`,`Owner`)",
		")ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "),
	)
	return err
}

// RemoveUserFromGroupByID 从组里移除用户
func RemoveUserFromGroupByID(group, user int64) error {
	conn := easysql.GetConn()

	cond := map[string]string{
		"Owner": strconv.FormatUint(uint64(user), 10),
		"Group": strconv.FormatUint(uint64(group), 10),
	}

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := conn.Where(cond).Update(groupMappingUserTableName, map[string]string{"DeletedTime": nowTime})
	return err
}

// AddUserToGroupByID 添加用户进组
func AddUserToGroupByID(group, user int64) error {
	conn := easysql.GetConn()

	data := map[string]string{
		"Owner": strconv.FormatUint(uint64(user), 10),
		"Group": strconv.FormatUint(uint64(group), 10),
	}

	_, err := conn.Insert(groupMappingUserTableName, data)
	return err
}

// IsAlreadyInGroup 是否已存在关联
func IsAlreadyInGroup(group, user int64) (bool, error) {
	conn := easysql.GetConn()

	cond := map[string]string{
		"Owner": strconv.FormatUint(uint64(user), 10),
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
