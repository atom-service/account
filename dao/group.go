package dao

import (
	"strings"

	"github.com/yinxulai/goutils/easysql"
)

const groupTableName = "group"
const groupMappingUserTableName = "group-mapping"

func createGroupTable() error {
	conn := easysql.GetConn()
	defer conn.Close()

	_, err := conn.ExecSQL(
		strings.Join([]string{
			" CREATE TABLE IF NOT EXISTS `" + groupTableName + "` (",
			" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
			" `Class` varchar(128) NOT NULL COMMENT '类型',",
			" `State` varchar(128) DEFAULT '' COMMENT '状态',",
			" `Value` varchar(512) DEFAULT '' COMMENT '值',",
			" `Owner` int(11) NOT NULL COMMENT '所属',",
			" `DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
			" `CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
			" `UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			" PRIMARY KEY (`ID`,`Class`)",
			" )ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
		}, "",
		),
	)
	return err
}

func createGroupMappingTable() error {
	conn := easysql.GetConn()
	defer conn.Close()

	_, err := conn.ExecSQL(
		strings.Join([]string{
			" CREATE TABLE IF NOT EXISTS `" + groupTableName + "` (",
			" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
			" `Class` varchar(128) NOT NULL COMMENT '类型',",
			" `State` varchar(128) DEFAULT '' COMMENT '状态',",
			" `Value` varchar(512) DEFAULT '' COMMENT '值',",
			" `Owner` int(11) NOT NULL COMMENT '所属',",
			" `DeletedTime` datetime DEFAULT NULL COMMENT '删除时间',",
			" `CreatedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
			" `UpdatedTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
			" PRIMARY KEY (`ID`,`Class`)",
			" )ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
		}, "",
		),
	)
	return err
}
