package dao

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/grpcbrick/account/models"
	"github.com/yinxulai/goutils/sqldb"
)

const labelTableName = "labels"
const labelHistoryTableName = "labels-history"
const labelMappingUserTableName = "label-mapping"
const labelMappingUserHistoryTableName = "label-mapping-history"

func truncateLabelTable() error {
	var err error
	masterStmp := sqldb.CreateStmt("truncate table `" + labelTableName + "`")
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}
	historyStmp := sqldb.CreateStmt("truncate table `" + labelHistoryTableName + "`")
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}
	return nil
}

func createLabelTable() error {
	var err error
	// 主表
	masterStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS",
		"`" + labelTableName + "`",
		"(",
		"`ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
		"`Name` varchar(128) NOT NULL COMMENT '名称',",
		"`Category` varchar(128) NOT NULL COMMENT '类型',",
		"`State` varchar(128) DEFAULT '' COMMENT '状态',",
		"`Value` varchar(512) DEFAULT '' COMMENT '值',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`ID`,`Name`,`Category`,`State`)",
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
		"`" + labelHistoryTableName + "`",
		"(",
		"`Index` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Index',",
		"`ID` int(11) COMMENT 'ID',",
		"`Name` varchar(128) COMMENT '名称',",
		"`Category` varchar(128) COMMENT '类型',",
		"`State` varchar(128) COMMENT '状态',",
		"`Value` varchar(512) COMMENT '简介',",
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

	return err
}

// CreateLabelHistory 对指定数据创建一条历史快照
func CreateLabelHistory(id int64) error {
	var err error
	namedData := map[string]interface{}{
		"ID": id,
	}

	// 插入一条更新历史
	historyStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + labelHistoryTableName + "`",
		"(`ID`,`Name`,`Category`,`State`,`Value`,`DeletedTime`,`CreatedTime`,`UpdatedTime`)",
		"SELECT",
		"`ID`,`Name`,`Category`,`State`,`Value`,`DeletedTime`,`CreatedTime`,`UpdatedTime`",
		"FROM",
		"`" + labelTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))
	_, err = historyStmp.Exec(namedData)
	if err != nil {
		return err
	}
	return nil
}

// CreateLabel 创建标签
func CreateLabel(name, category, state, value string) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + labelTableName + "`",
		"(`Name`, `Category`, `State`, `Value`)",
		"VALUES",
		"(:Name, :Category, :State, :Value)",
	}, " "))

	namedData := map[string]interface{}{
		"Name":     name,
		"Category": category,
		"State":    state,
		"Value":    value,
	}

	result, err := stmp.Exec(namedData)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

// CountLabelByID 根据 id 统计
func CountLabelByID(id int64) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + labelTableName + "`",
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

// QueryLabels 查询标签
func QueryLabels(page, limit int64) (totalPage, currentPage int64, labels []*models.Label, err error) {
	currentPage = page // 固定当前页

	// 查询数据长度
	countStmp := sqldb.CreateStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + labelTableName + "`",
	}, " "))

	countResult := struct{ Count int64 }{}
	err = countStmp.Get(&countResult)
	if err != nil {
		return totalPage, currentPage, labels, err
	}

	count := countResult.Count
	// 计算总页码数
	totalPage = int64(math.Ceil(float64(count) / float64(limit)))

	// 超出数据总页数
	if page > totalPage {
		// 返回当前页、空数据（当前页确实不存在数据）
		return totalPage, page, labels, err
	}

	// 计算偏移
	offSet := (page - 1) * limit
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + labelTableName + "`",
		"LIMIT :Limit",
		"OFFSET :Offset",
	}, " "))

	labels = []*models.Label{}
	namedData := map[string]interface{}{
		"Limit":  limit,
		"Offset": offSet,
	}

	err = stmp.Select(&labels, namedData)
	if err != nil {
		return totalPage, currentPage, labels, err
	}

	return totalPage, currentPage, labels, err
}

// QueryLabelByID 根据 id 查询
func QueryLabelByID(id int64) (*models.Label, error) {
	var err error
	namedData := map[string]interface{}{"ID": id}
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + labelTableName + "`",
		"WHERE",
		"`ID`=:ID",
	}, " "))

	label := new(models.Label)
	err = stmp.Get(label, namedData)
	if err != nil {
		return nil, err
	}

	return label, nil
}

// DeleteLabelByID 删除标签
func DeleteLabelByID(id int64) error {
	return updataLabelFieldByID(id, map[string]interface{}{"DeletedTime": time.Now()})
}

// UpdateLabelNameByID 更新标签类型
func UpdateLabelNameByID(id int64, name string) error {
	return updataLabelFieldByID(id, map[string]interface{}{"Name": name})
}

// UpdateLabelCategoryByID 更新标签类型
func UpdateLabelCategoryByID(id int64, category string) error {
	return updataLabelFieldByID(id, map[string]interface{}{"Category": category})
}

// UpdateLabelStateByID 更新标签状态
func UpdateLabelStateByID(id int64, category string) error {
	return updataLabelFieldByID(id, map[string]interface{}{"State": category})
}

// UpdateLabelValueByID 更新标签值
func UpdateLabelValueByID(id int64, category string) error {
	return updataLabelFieldByID(id, map[string]interface{}{"Value": category})
}

// 根据 ID 更新标签
func updataLabelFieldByID(id int64, field map[string]interface{}) error {
	var err error

	fieldSQL := []string{}
	for name := range field {
		fieldSQL = append(fieldSQL, fmt.Sprintf("`%s`=:%s", name, name))
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"UPDATE",
		"`" + labelTableName + "`",
		"SET",
		strings.Join(fieldSQL, ","),
		"WHERE",
		"`ID`=:ID",
	}, " "))

	// 修改前创建历史
	err = CreateLabelHistory(id)
	if err != nil {
		return err
	}

	// 更新
	field["ID"] = id
	_, err = stmp.Exec(field)
	return err
}

func truncateLabelMappingTable() error {
	var err error
	masterStmp := sqldb.CreateStmt("truncate table `" + labelMappingUserTableName + "`")
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}
	historyStmp := sqldb.CreateStmt("truncate table `" + labelMappingUserHistoryTableName + "`")
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}
	return nil
}

// 映射关系
func createLabelMappingTable() error {
	// 主表
	matserStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + labelMappingUserTableName + "`",
		"(",
		"`Label` int(11) NOT NULL COMMENT '标签 ID',",
		"`User` int(11) NOT NULL COMMENT '所属者 ID',",
		"`DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
		"`CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`Label`,`User`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err := matserStmp.Exec()
	if err != nil {
		return err
	}
	// 历史表
	historyStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + labelMappingUserHistoryTableName + "`",
		"(",
		"`Index` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Index',",
		"`Label` int(11)  COMMENT '标签 ID',",
		"`User` int(11)  COMMENT '所属者 ID',",
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

// RemoveLabelFromUserByID 从用户身上移除标签
func RemoveLabelFromUserByID(label, user int64) error {
	var err error

	namedData := map[string]interface{}{
		"User":        user,
		"Label":       label,
		"DeletedTime": time.Now(),
	}

	updateStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"UPDATE",
		"`" + labelMappingUserTableName + "`",
		"SET",
		"`DeletedTime`=:DeletedTime", // 更新删除时间
		"WHERE",
		"`User`=:User",
		"AND",
		"`Label`=:Label",
	}, " "))

	// 先保存一条历史快照
	err = CreateLabelMappingHistory(label, user)
	if err != nil {
		return err
	}

	// 更新主标数据
	_, err = updateStmp.Exec(namedData)
	return err
}

// AddLabelToUserByID 添加标签给用户
func AddLabelToUserByID(label, user int64) error {
	var err error

	namedData := map[string]interface{}{
		"User":  user,
		"Label": label,
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + labelMappingUserTableName + "`",
		"(`User`, `Label`)",
		"VALUES",
		"(:User, :Label)",
	}, " "))

	_, err = stmp.Exec(namedData)
	if err != nil {
		return err
	}

	return nil
}

// IsAlreadyOwnLabel 是否已存在关联
func IsAlreadyOwnLabel(label, user int64) (bool, error) {
	namedData := map[string]interface{}{
		"User":  user,
		"Label": label,
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + labelMappingUserTableName + "`",
		"WHERE",
		"`User`=:User",
		"AND",
		"`Label`=:Label",
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

// CreateLabelMappingHistory 对指定数据创建一条历史快照
func CreateLabelMappingHistory(label, user int64) error {

	namedData := map[string]interface{}{
		"User":  user,
		"Label": label,
	}

	// 插入一条更新历史
	historyStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + labelMappingUserHistoryTableName + "`",
		"(`Label`,`User`,`DeletedTime`,`CreatedTime`,`UpdatedTime`)",
		"SELECT",
		"`Label`,`User`,`DeletedTime`,`CreatedTime`,`UpdatedTime`",
		"FROM",
		"`" + labelMappingUserTableName + "`",
		"WHERE",
		"`User`=:User",
		"AND",
		"`Label`=:Label",
	}, " "))
	_, err := historyStmp.Exec(namedData)
	if err != nil {
		return err
	}
	return nil
}
