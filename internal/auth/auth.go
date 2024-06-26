package auth

import (
	"context"
	"log/slog"
	"strings"

	"github.com/atom-service/account/internal/config"
	"github.com/atom-service/account/internal/model"
	publicAuth "github.com/atom-service/account/package/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func matchRuntimeSecretIfExits(ctx context.Context, secretKey string) *model.Secret {
	for _, secret := range config.Secrets {
		if secret != nil && secret.Key == secretKey {
			return &model.Secret{
				UserID: &model.AdminUserID,
				Key:    &secret.Key,
				Value:  &secret.Value,
			}
		}
	}
	return nil
}

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

	firstToken := tokens[0]

	// 标准的 Bearer token
	if strings.HasPrefix(firstToken, "Bearer") {
		// 从 Bearer token 中提取 token 值
		firstToken = strings.TrimPrefix(firstToken, "Bearer ")
	}

	tokenInfo, err := publicAuth.ParseToken(firstToken)
	if err != nil {
		slog.ErrorContext(ctx, "Invalid token, possibly invalid secret", slog.String("token", firstToken), slog.Any("error", err))
		return ctx
	}

	if tokenInfo.SecretKey == "" {
		slog.DebugContext(ctx, "Invalid token, possibly invalid secret", slog.String("token", firstToken))
		return ctx
	}

	paginationLimit := int64(1)
	paginationOption := &model.Pagination{Limit: &paginationLimit}
	var secretInfo = matchRuntimeSecretIfExits(ctx, tokenInfo.SecretKey)

	if secretInfo == nil {
		secretSelector := model.SecretSelector{Key: &tokenInfo.SecretKey}
		querySecretsResponse, err := model.SecretTable.QuerySecrets(ctx, secretSelector, paginationOption, nil)
		if err != nil {
			slog.ErrorContext(ctx, "Invalid token, possibly invalid secret", slog.String("token", firstToken), slog.Any("error", err))
			return ctx
		}

		if len(querySecretsResponse) == 0 {
			slog.DebugContext(ctx, "Invalid token, possibly invalid secret", slog.String("token", firstToken))
			return ctx
		}

		secretInfo = querySecretsResponse[0]
	}

	// 是否已经被禁用
	if secretInfo.IsDisabled() {
		slog.DebugContext(ctx, "Invalid token, possibly invalid secret, secret is disabled", slog.String("token", firstToken))
		return ctx
	}

	if !publicAuth.VerifyToken(*secretInfo.Key, *secretInfo.Value, firstToken) {
		slog.DebugContext(ctx, "Invalid token, possibly invalid secret, verify token failed", slog.String("token", firstToken))
		return ctx
	}

	userSelector := model.UserSelector{ID: secretInfo.UserID}
	queryUserResponse, err := model.UserTable.QueryUsers(ctx, userSelector, paginationOption, nil)
	if err != nil {
		slog.ErrorContext(ctx, "Invalid token, possibly invalid secret", slog.String("token", firstToken), slog.Any("error", err))
		return ctx
	}

	if len(queryUserResponse) == 0 {
		slog.DebugContext(ctx, "Invalid token, possibly invalid secret, secret user not found", slog.String("token", firstToken))
		return ctx
	}

	// 验证 token 是否有效
	firstUser := queryUserResponse[0]
	summaryForUserRequest := model.UserResourceSummarySelector{UserID: firstUser.ID}
	summaryForUserResponse, err := model.Permission.QueryUserResourceSummaries(ctx, summaryForUserRequest)
	if err != nil {
		slog.ErrorContext(ctx, "Invalid token, possibly invalid secret", slog.Any("error", err))
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
