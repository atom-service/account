package mongo

import (
	"github.com/yinxulai/goutils/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

func getUserCollection() *mongoDriver.Collection {
	return mongo.GetCollection("user")
}

// 根据 id 统计
func countUserByID(id uint64) int {
	return 1
}

// 根据用户名统计
func countUserByUsername(username string) int {
	return 1
}

// 创建用户
func createUser(username, password string) bool {
	_, err := getUserCollection().InsertOne(mongo.GetContext(), bson.M{
		"username": username,
		"password": password,
	})
	if err != nil {
		return false
	}

	return true
}

// 更新用户类型
func updateUserTypeByID(id uint64, typo string) bool {
	return true
}

// 更新用户头像
func updateUserAvatarByID(id uint64, avatar string) bool {
	return true
}

// 更新用户昵称
func updateUserNicknameByID(id uint64, nickname string) bool {
	return true
}

// 更新用户密码
func updateUserPasswordByID(id uint64, password string) bool {
	return true
}
