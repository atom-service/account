package provider

import (
	"bytes"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql 驱动
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	createLabelTableStmt       *sqlx.Stmt
	createUserTableStmt        *sqlx.Stmt
	countLabelByIDStmt         *sqlx.NamedStmt
	countUserByUsernameStmt    *sqlx.NamedStmt
	countLabelByOwnerStmt      *sqlx.NamedStmt
	queryLabelByOwnerStmt      *sqlx.NamedStmt
	deleteLabelByIDStmt        *sqlx.NamedStmt
	updateLabelByIDStmt        *sqlx.NamedStmt
	queryLabelByIDStmt         *sqlx.NamedStmt
	insertLabelByOwnerStmt     *sqlx.NamedStmt
	countUserByIDStmt          *sqlx.NamedStmt
	updateUserPasswordByIDStmt *sqlx.NamedStmt
	deleteUserByIDStmt         *sqlx.NamedStmt
	updateUserByIDStmt         *sqlx.NamedStmt
	queryUserByIDStmt          *sqlx.NamedStmt
	insertUserStmt             *sqlx.NamedStmt
	queryUserByUsernameStmt    *sqlx.NamedStmt
)

func init() {
	var err error
	godotenv.Load()

	database, err := sqlx.Connect("mysql", os.Getenv("MYSQL_DB_URI"))
	if err != nil {
		panic(err)
	}

	// 设置 Name 映射方法
	database.MapperFunc(func(field string) string { return field })

	// 创建用户表
	createUserTableStmt = MustPreparex(
		database,
		" CREATE TABLE IF NOT EXISTS `users`(",
		" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
		" `Type` varchar(128) NOT NULL COMMENT '账户类型',",
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
	)
	_, err = createUserTableStmt.Exec()
	if err != nil {
		panic(err)
	}

	// 创建标签表
	createLabelTableStmt = MustPreparex(
		database,
		" CREATE TABLE IF NOT EXISTS `labels` (",
		" `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',",
		" `Type` varchar(128) NOT NULL COMMENT '类型',",
		" `State` varchar(128) DEFAULT '' COMMENT '状态',",
		" `Value` varchar(512) DEFAULT '' COMMENT '值',",
		" `Owner` int(11) NOT NULL COMMENT '所属',",
		" `CreateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,",
		" PRIMARY KEY (`ID`,`Type`)",
		" )ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	)
	_, err = createLabelTableStmt.Exec()
	if err != nil {
		panic(err)
	}

	// 通过 ID 查询用户
	queryUserByIDStmt = MustPreparexNamed(
		database,
		" SELECT * FROM `users`",
		" WHERE `ID` = :ID",
		" ;",
	)
	// 通过 用户名 查询用户
	queryUserByUsernameStmt = MustPreparexNamed(
		database,
		" SELECT * FROM `users`",
		" WHERE `Username` = :Username",
		" ;",
	)
	// 通过 ID 统计用户
	countUserByIDStmt = MustPreparexNamed(
		database,
		" SELECT COUNT(*) FROM `users`",
		" WHERE `ID` = :ID",
		" ;",
	)
	// 通过 Username 统计用户
	countUserByUsernameStmt = MustPreparexNamed(
		database,
		" SELECT COUNT(*) FROM `users`",
		" WHERE `Username` = :Username",
		" ;",
	)
	// 通过 ID 更新用户密码
	updateUserPasswordByIDStmt = MustPreparexNamed(
		database,
		" UPDATE `users` SET",
		" `Password` = :Password",
		" WHERE `ID` = :ID",
		" ;",
	)
	// 通过 ID 删除用户
	deleteUserByIDStmt = MustPreparexNamed(
		database,
		" DELETE FROM `users`",
		" WHERE`ID` = :ID",
		" ;",
	)
	// 通过 Id 更新用户信息
	updateUserByIDStmt = MustPreparexNamed(
		database,
		" UPDATE `users` SET",
		" `Type` = :Type,",
		" `Avatar` = :Avatar,",
		" `Nickname` = :Nickname,",
		" `Username` = :Username",
		" WHERE `ID` = :ID",
		" ;",
	)
	// 插入用户
	insertUserStmt = MustPreparexNamed(
		database,
		" INSERT INTO `users`",
		" (`Type`, `Avatar`, `Inviter`, `Nickname`, `Username`, `Password`)",
		" VALUES",
		" (:Type, :Avatar, :Inviter, :Nickname, :Username, :Password)",
		" ;",
	)

	countLabelByOwnerStmt = MustPreparexNamed(
		database,
		" SELECT COUNT(*) FROM `labels`",
		" WHERE `Owner` = :Owner",
		" ;",
	)

	// 通过 Owner 查询标签
	queryLabelByOwnerStmt = MustPreparexNamed(
		database,
		" SELECT * FROM `labels`",
		" WHERE `Owner` = :Owner",
		" LIMIT :Limit",
		" OFFSET :Offset",
		" ;",
	)

	// 通过 ID 统计 lable 数量
	countLabelByIDStmt = MustPreparexNamed(
		database,
		" SELECT COUNT(*) FROM `labels`",
		" WHERE `ID` = :ID",
		" ;",
	)
	// 通过 ID 统计 lable 数量
	deleteLabelByIDStmt = MustPreparexNamed(
		database,
		" DELETE FROM `labels`",
		" WHERE `ID` = :ID",
		" ;",
	)
	// 更新指定标签
	updateLabelByIDStmt = MustPreparexNamed(
		database,
		" UPDATE `labels` SET",
		" `Type` = :Type,",
		" `State` = :State,",
		" `Value` = :Value",
		" WHERE `ID` = :ID",
		" ;",
	)
	// 通过ID查询记录
	queryLabelByIDStmt = MustPreparexNamed(
		database,
		" SELECT * FROM `labels`",
		" WHERE `ID` = :ID",
		" ;",
	)
	// 插入记录给 制定用户 Owner
	insertLabelByOwnerStmt = MustPreparexNamed(
		database,
		" INSERT INTO `labels`",
		" (`Type`, `State`, `Value`, `Owner`)",
		" VALUES",
		" (:Type, :State, :Value, :Owner)",
		" ;",
	)
}

// MustPreparex 解析 query
func MustPreparex(database *sqlx.DB, querys ...string) *sqlx.Stmt {
	var queryBuf bytes.Buffer

	for _, s := range querys {
		queryBuf.WriteString(s)
	}

	stmp, err := database.Preparex(queryBuf.String())
	if err != nil {
		panic(err)
	}
	return stmp
}

// MustPreparexNamed 解析 query
func MustPreparexNamed(database *sqlx.DB, querys ...string) *sqlx.NamedStmt {
	var queryBuf bytes.Buffer

	for _, s := range querys {
		queryBuf.WriteString(s)
	}

	stmp, err := database.PrepareNamed(queryBuf.String())
	if err != nil {
		panic(err)
	}
	return stmp
}
