package auth

import (
	"context"
	"log/slog"

	"github.com/atom-service/account/internal/config"
	"github.com/atom-service/account/internal/model"
	publicAuth "github.com/atom-service/account/package/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type serverAuthInterceptor struct {
}

func NewServerAuthInterceptor() *serverAuthInterceptor {
	return &serverAuthInterceptor{}
}

func (ai *serverAuthInterceptor) resolveUserIncomingContext(ctx context.Context) context.Context {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	tokens, ok := metadata["authorization"]
	if !ok || len(tokens) == 0 {
		return ctx
	}

	tokenInfo, err := publicAuth.ParseToken(tokens[0])
	if err != nil || tokenInfo.SecretKey == "" {
		return ctx
	}

	var secretInfo *model.Secret

	paginationLimit := int64(1)
	paginationOption := &model.Pagination{Limit: &paginationLimit}

	if config.Secret != nil && config.Secret.Key == tokenInfo.SecretKey {
		// 管理员身份
		secretInfo = &model.Secret{
			UserID: &model.AdminUserID,
			Key:    &config.Secret.Key,
			Value:  &config.Secret.Value,
		}
	}

	if secretInfo == nil {
		secretSelector := model.SecretSelector{Key: &tokenInfo.SecretKey}
		querySecretsResponse, err := model.SecretTable.QuerySecrets(ctx, secretSelector, paginationOption, nil)
		if err != nil || len(querySecretsResponse) == 0 {
			slog.InfoContext(ctx, " Invalid token, possibly invalid secret")
			return ctx
		}

		secretInfo = querySecretsResponse[0]
	}

	// 是否已经被禁用
	if secretInfo.IsDisabled() {
		slog.InfoContext(ctx, " Invalid token, possibly invalid secret")
		return ctx
	}

	if !publicAuth.VerifyToken(*secretInfo.Key, *secretInfo.Value, tokens[0]) {
		slog.InfoContext(ctx, " Invalid token, possibly invalid secret")
		return ctx
	}

	userSelector := model.UserSelector{ID: secretInfo.UserID}
	queryUserResponse, err := model.UserTable.QueryUsers(ctx, userSelector, paginationOption, nil)
	if err != nil || len(queryUserResponse) == 0 {
		slog.InfoContext(ctx, " Invalid token, possibly invalid secret")
		return ctx
	}

	// 验证 token 是否有效
	firstUser := queryUserResponse[0]

	summaryForUserRequest := model.UserResourceSummarySelector{UserID: firstUser.ID}
	summaryForUserResponse, err := model.Permission.QueryUserResourceSummaries(ctx, summaryForUserRequest)
	if err != nil {
		slog.InfoContext(ctx, " Invalid token, possibly invalid secret")
		return ctx
	}

	ctx = context.WithValue(ctx, publicAuth.ContextUserSymbol, firstUser)
	ctx = context.WithValue(ctx, publicAuth.ContextSecretSymbol, secretInfo)
	ctx = context.WithValue(ctx, publicAuth.ContextPermissionsSymbol, summaryForUserResponse)
	return ctx
}

func (ai *serverAuthInterceptor) ServerUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	newCtx := ai.resolveUserIncomingContext(ctx)
	resp, err = handler(newCtx, req)
	return resp, err
}
