package auth

import (
	"context"

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

	paginationLimit := int64(1)
	paginationOption := &model.Pagination{Limit: &paginationLimit}
	secretSelector := model.SecretSelector{Key: &tokenInfo.SecretKey}
	querySecretsResponse, err := model.SecretTable.QuerySecrets(ctx, secretSelector, paginationOption, nil)
	if err != nil || len(querySecretsResponse) == 0 {
		return ctx
	}

	// 验证 token 是否有效
	firstSecret := querySecretsResponse[0]

	// 是否已经被禁用
	if firstSecret.IsDisabled() {
		return ctx
	}

	if !publicAuth.VerifyToken(*firstSecret.Key, *firstSecret.Value, tokens[0]) {
		return ctx
	}

	userSelector := model.UserSelector{ID: firstSecret.UserID}
	queryUserResponse, err := model.UserTable.QueryUsers(ctx, userSelector, paginationOption, nil)
	if err != nil || len(queryUserResponse) == 0 {
		return ctx
	}

	// 验证 token 是否有效
	firstUser := queryUserResponse[0]

	summaryForUserRequest := model.UserResourceSummarySelector{UserID: firstUser.ID}
	summaryForUserResponse, err := model.Permission.QueryUserResourceSummaries(ctx, summaryForUserRequest)
	if err != nil {
		return ctx
	}

	ctx = context.WithValue(ctx, publicAuth.ContextUserSymbol, firstUser)
	ctx = context.WithValue(ctx, publicAuth.ContextSecretSymbol, firstSecret)
	ctx = context.WithValue(ctx, publicAuth.ContextPermissionsSymbol, summaryForUserResponse)
	return ctx
}

func (ai *serverAuthInterceptor) ServerUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	newCtx := ai.resolveUserIncomingContext(ctx)
	resp, err = handler(newCtx, req)
	return resp, err
}
