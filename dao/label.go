package dao

import (
	"strconv"
	"strings"
	"time"

	"github.com/grpcbrick/account/model"
	"github.com/yinxulai/goutils/easysql"
)

const labelTableName = "labels"
const labelMappingUserTableName = "label-mapping"

func createLabelTable() error {
	conn := easysql.GetConn()
	defer conn.Close()

	_, err := conn.ExecSQL(
		strings.Join([]string{
			" CREATE TABLE IF NOT EXISTS `" + labelTableName + "` (",
			" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
			" `Name` varchar(128) NOT NULL COMMENT '名称',",
			" `Class` varchar(128) NOT NULL COMMENT '类型',",
			" `State` varchar(128) DEFAULT '' COMMENT '状态',",
			" `Value` varchar(512) DEFAULT '' COMMENT '值',",
			" `DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
			" `CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
			" `UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			" PRIMARY KEY (`ID`,`Class`,`DeletedTime`)",
			" )ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
		}, "",
		),
	)
	return err
}

// CreateLabel 创建标签
func createLabel(name, class, state, value string) error {
	conn := easysql.GetConn()
	defer conn.Close()

	data := map[string]string{
		"Name":  name,
		"Class": class,
		"State": state,
		"Value": value,
	}

	_, err := conn.Insert(labelTableName, data)
	return err
}

// CountLabelByID 根据 id 统计
func CountLabelByID(id uint64) (int, error) {
	conn := easysql.GetConn()
	defer conn.Close()

	idstr := strconv.FormatUint(id, 10)
	cond := map[string]string{"ID": idstr}
	queryField := []string{"count(*) as count"}
	result, err := conn.Select(labelTableName, queryField).Where(cond).QueryRow()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(result["count"])
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryLabelByID 根据 id 查询
func QueryLabelByID(id uint64) (*model.Label, error) {
	conn := easysql.GetConn()
	defer conn.Close()

	idstr := strconv.FormatUint(id, 10)
	cond := map[string]string{"ID": idstr}
	result, err := conn.Select(labelTableName, nil).Where(cond).QueryRow()
	if err != nil {
		return nil, err
	}

	lable := new(model.Label)
	lable.LoadStringMap(result)
	return lable, nil
}

// DeleteLabelByID 删除标签
func DeleteLabelByID(id uint64, class string) error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	return updataLabelFieldByID(id, map[string]string{"DeletedTime": nowTime})
}

// UpdateLabelClassByID 更新标签类型
func UpdateLabelClassByID(id uint64, class string) error {
	return updataLabelFieldByID(id, map[string]string{"Class": class})
}

// UpdateLabelStateByID 更新标签状态
func UpdateLabelStateByID(id uint64, class string) error {
	return updataLabelFieldByID(id, map[string]string{"State": class})
}

// UpdateLabelValueByID 更新标签值
func UpdateLabelValueByID(id uint64, class string) error {
	return updataLabelFieldByID(id, map[string]string{"Value": class})
}

// 根据 ID 更新标签
func updataLabelFieldByID(id uint64, field map[string]string) error {
	conn := easysql.GetConn()
	defer conn.Close()

	cond := map[string]string{"ID": strconv.FormatUint(id, 10)}
	_, err := conn.Where(cond).Update(labelTableName, field)
	return err
}

// 映射关系
func createLabelMappingTable() error {
	conn := easysql.GetConn()
	defer conn.Close()

	_, err := conn.ExecSQL(
		strings.Join([]string{
			" CREATE TABLE IF NOT EXISTS `" + labelMappingUserTableName + "` (",
			" `label` int(11) NOT NULL COMMENT '标签 ID',",
			" `Owner` int(11) NOT NULL COMMENT '所属者 ID',",
			" `DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
			" `CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
			" `UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			" PRIMARY KEY (`Source`,`Owner`,`DeletedTime`)",
			" )ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
		}, "",
		),
	)
	return err
}

// RemoveLabelFromUserByID 从用户身上移除标签
func RemoveLabelFromUserByID(label, user uint64) error {
	conn := easysql.GetConn()
	defer conn.Close()

	cond := map[string]string{
		"Owner": strconv.FormatUint(user, 10),
		"label": strconv.FormatUint(label, 10),
	}

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := conn.Where(cond).Update(labelMappingUserTableName, map[string]string{"DeletedTime": nowTime})
	return err
}

// AddLabelToUserByID 添加标签给用户
func AddLabelToUserByID(label, user uint64) error {
	conn := easysql.GetConn()
	defer conn.Close()

	data := map[string]string{
		"Owner": strconv.FormatUint(user, 10),
		"label": strconv.FormatUint(label, 10),
	}

	_, err := conn.Insert(labelMappingUserTableName, data)
	return err
}
