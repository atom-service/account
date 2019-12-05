package dao

import (
	"strconv"
	"strings"
	"time"

	"github.com/grpcbrick/account/models"
	"github.com/yinxulai/goutils/easysql"
)

const groupTableName = "group"
const groupMappingUserTableName = "group-mapping"

func truncateGroupTable() error {
	conn := easysql.GetConn()

	_, err := conn.ExecSQL("truncate table `" + groupTableName + "`")
	return err
}

func createGroupTable() error {
	conn := easysql.GetConn()

	_, err := conn.ExecSQL(
		strings.Join([]string{
			" CREATE TABLE IF NOT EXISTS `" + groupTableName + "` (",
			" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
			" `Name` varchar(128) NOT NULL COMMENT '名称',",
			" `Class` varchar(128) NOT NULL COMMENT '类型',",
			" `State` varchar(128) DEFAULT '' COMMENT '状态',",
			" `Description` varchar(512) DEFAULT '' COMMENT '简介',",
			" `DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
			" `CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
			" `UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			" PRIMARY KEY (`ID`,`Name`,`Class`,`State`),",
			" UNIQUE KEY `Name` (`Name`)",
			" )ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
		}, "",
		),
	)
	return err
}

// CreateGroup 创建组
func CreateGroup(name, class, state, description string) (int64, error) {
	conn := easysql.GetConn()

	data := map[string]string{
		"Name":        name,
		"Class":       class,
		"State":       state,
		"Description": description,
	}

	id, err := conn.Insert(groupTableName, data)
	return id, err
}

// CountGroupByName 根据 name 统计
func CountGroupByName(name string) (int, error) {
	conn := easysql.GetConn()

	cond := map[string]string{"Name": "'" + name + "'"}
	queryField := []string{"count(*) as count"}
	result, err := conn.Select(groupTableName, queryField).Where(cond).QueryRow()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(result["count"])
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CountGroupByID 根据 id 统计
func CountGroupByID(id uint32) (int, error) {
	conn := easysql.GetConn()

	idstr := strconv.FormatUint(uint64(id), 10)
	cond := map[string]string{"ID": idstr}
	queryField := []string{"count(*) as count"}
	result, err := conn.Select(groupTableName, queryField).Where(cond).QueryRow()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(result["count"])
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryGroups 查询组
func QueryGroups(page, limit int) (totalPage, currentPage int, groups []*models.Group, err error) {
	conn := easysql.GetConn()
	result, err := conn.Select(groupTableName, nil).Pagination(page, limit)
	if err != nil {
		return 0, 0, nil, err
	}

	groups = []*models.Group{}
	totalPage = result["totalPage"].(int)
	currentPage = result["currentPage"].(int)
	for _, mapData := range result["rows"].([]interface{}) {
		group := new(models.Group)
		group.LoadStringMap(mapData.(map[string]string))
		groups = append(groups, group)
	}

	return totalPage, currentPage, groups, err
}

// QueryGroupByID 根据 id 查询
func QueryGroupByID(id uint32) (*models.Group, error) {
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
func DeleteGroupByID(id uint32) error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	return updataGroupFieldByID(id, map[string]string{"DeletedTime": nowTime})
}

// UpdateGroupNameByID 更新标签类型
func UpdateGroupNameByID(id uint32, name string) error {
	return updataGroupFieldByID(id, map[string]string{"Name": name})
}

// UpdateGroupClassByID 更新标签状态
func UpdateGroupClassByID(id uint32, class string) error {
	return updataGroupFieldByID(id, map[string]string{"Class": class})
}

// UpdateGroupStateByID 更新标签值
func UpdateGroupStateByID(id uint32, class string) error {
	return updataGroupFieldByID(id, map[string]string{"State": class})
}

// UpdateGroupDescriptionByID 更新标签值
func UpdateGroupDescriptionByID(id uint32, description string) error {
	return updataGroupFieldByID(id, map[string]string{"Description": description})
}

// 根据 ID 更新标签
func updataGroupFieldByID(id uint32, field map[string]string) error {
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

	_, err := conn.ExecSQL(
		strings.Join([]string{
			" CREATE TABLE IF NOT EXISTS `" + groupMappingUserTableName + "` (",
			" `Group` int(11) NOT NULL COMMENT '组 ID',",
			" `Owner` int(11) NOT NULL COMMENT '所属者 ID',",
			" `DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
			" `CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
			" `UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			" PRIMARY KEY (`Group`,`Owner`)",
			" )ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
		}, "",
		),
	)
	return err
}

// RemoveUserFromGroupByID 从组里移除用户
func RemoveUserFromGroupByID(group, user uint32) error {
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
func AddUserToGroupByID(group, user uint32) error {
	conn := easysql.GetConn()

	data := map[string]string{
		"Owner": strconv.FormatUint(uint64(user), 10),
		"Group": strconv.FormatUint(uint64(group), 10),
	}

	_, err := conn.Insert(groupMappingUserTableName, data)
	return err
}

// IsAlreadyInGroup 是否已存在关联
func IsAlreadyInGroup(group, user uint32) (bool, error) {
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