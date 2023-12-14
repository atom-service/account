package user

import (
	"context"

	"github.com/protect-we-network/server/internal/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var secretCollection *mongo.Collection
var userCollection *mongo.Collection

func Init(database *mongo.Database) {
	userCollection = database.Collection("user")
	secretCollection = database.Collection("user-secret")
}

func CreateUser(ctx context.Context, user *User) error {
	_, err := userCollection.InsertOne(ctx, bson.M{
		"email":    user.Email,
		"password": user.Password,
		"username": user.Username,
	})

	logger.Error(err)
	return err
}

type QueryUserParams struct {
	ID       *string
	Email    *string
	Username *string
}

func QueryUser(ctx context.Context, params *QueryUserParams) (*User, error) {
	query := bson.M{}

	// 检查是否提供了 ID 参数
	if params.ID != nil {
		query["_id"] = *params.ID
	}

	// 检查是否提供了 Email 参数
	if params.Email != nil {
		query["email"] = *params.Email
	}

	// 检查是否提供了 Username 参数
	if params.Username != nil {
		query["username"] = *params.Username
	}

	var user User
	err := userCollection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &user, nil
}

type CreateUserSecretParams struct {
	User string
}

func CreateUserSecret(ctx context.Context, params *CreateUserSecretParams) error {
	_, err := secretCollection.InsertOne(ctx, bson.M{
		"user": params.User,
	})

	logger.Error(err)
	return err
}

type RemoveUserSecretParams struct {
	User   string
	secret string
}

func RemoveUserSecret(ctx context.Context, params *RemoveUserSecretParams) error {
	update := bson.M{"$pull": bson.M{"secrets": params.secret}}
	_, err := userCollection.UpdateOne(ctx, bson.M{"_id": params.User}, update)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func ListUserBySecrets(ctx context.Context, secrets []string) ([]*User, error) {
	filter := bson.M{"secrets": bson.M{"$in": secrets}}

	cursor, err := userCollection.Find(ctx, filter)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*User

	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		logger.Error(err)
		return nil, err
	}

	return users, nil
}
