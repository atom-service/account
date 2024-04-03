package auth

import (
	"context"
	"time"

	"github.com/atom-service/account/internal/model"
	proto "github.com/atom-service/account/package/proto"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	metadata "google.golang.org/grpc/metadata"
)

type contextSymbol struct{ name string }

var (
	ContextUserSymbol        = contextSymbol{"ContextUserSymbol"}
	ContextSecretSymbol      = contextSymbol{"ContextSecretSymbol"}
	ContextPermissionsSymbol = contextSymbol{"ContextPermissionsSymbol"}
)

type AuthWithTokenCredentials struct {
	Token string
}

func (x *AuthWithTokenCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": x.Token,
	}, nil
}

func (x *AuthWithTokenCredentials) RequireTransportSecurity() bool {
	return false
}

type AuthWithSecretCredentials struct {
	SecretKey   string
	SecretValue string
}

func (x *AuthWithSecretCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": SignToken(x.SecretKey, x.SecretValue, SignData{
			ExpiresAt: time.Now().Add(24 * 7 * time.Hour), // 1 周有效期
		}),
	}, nil
}

func (x *AuthWithSecretCredentials) RequireTransportSecurity() bool {
	return false
}

type serverAuthInterceptor struct {
	secretKey        string
	secretValue      string
	accountClient    proto.AccountServiceClient
	permissionClient proto.PermissionServiceClient
}

func NewServerAuthInterceptor(accountServerHost, secretKey, secretValue string) *serverAuthInterceptor {
	authCredentials := grpc.WithPerRPCCredentials(&AuthWithSecretCredentials{secretKey, secretValue})
	nonSafeCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(accountServerHost, authCredentials, nonSafeCredentials)
	if err != nil {
		panic(err)
	}

	return &serverAuthInterceptor{
		secretKey:        secretKey,
		secretValue:      secretValue,
		accountClient:    proto.NewAccountServiceClient(conn),
		permissionClient: proto.NewPermissionServiceClient(conn),
	}
}

// 通过 ctx 中的 authorization 解析用户信息，并设置到 ctx，以便程序访问
func (ai *serverAuthInterceptor) resolveUserIncomingContext(ctx context.Context) context.Context {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	tokens, ok := metadata["authorization"]
	if !ok || len(tokens) == 0 {
		return ctx
	}

	tokenInfo, err := ParseToken(tokens[0])
	if err != nil || tokenInfo.SecretKey == "" {
		return ctx
	}

	// 说明是 serverAuthInterceptor 自己的请求
	// 比如下面代码中的 QuerySecrets、QueryUsers 等
	// 如果是自己的请求，则直接构造一个 Secret，不然就死循环了
	isSelfBackToken := VerifyToken(ai.secretKey, ai.secretValue, tokens[0])
	if isSelfBackToken {
		ctx = context.WithValue(ctx, ContextSecretSymbol, &proto.Secret{
			Key:   ai.secretKey,
			Value: ai.secretValue,
		})
		return ctx
	}

	paginationLimit := int64(1)
	secretSelector := &proto.SecretSelector{Key: &tokenInfo.SecretKey}
	paginationOption := &proto.PaginationOption{Limit: &paginationLimit}
	querySecretsRequest := &proto.QuerySecretsRequest{Selector: secretSelector, Pagination: paginationOption}
	querySecretsResponse, err := ai.accountClient.QuerySecrets(ctx, querySecretsRequest)
	if err != nil || querySecretsResponse.State != proto.State_SUCCESS {
		return ctx
	}
	if querySecretsResponse.Data.Total == 0 {
		return ctx
	}

	// 验证 token 是否有效
	secret := querySecretsResponse.Data.Secrets[0]

	// 是否已经被禁用
	secretModel := new(model.Secret)
	secretModel.LoadProto(secret)
	if secretModel.IsDisabled() {
		return ctx
	}

	if !VerifyToken(secret.Key, secret.Value, tokens[0]) {
		return ctx
	}

	userSelector := &proto.UserSelector{ID: &secret.UserID}
	queryUserResponse, err := ai.accountClient.QueryUsers(ctx, &proto.QueryUsersRequest{Selector: userSelector})
	if err != nil || queryUserResponse.State != proto.State_SUCCESS || querySecretsResponse.Data.Total == 0 {
		return ctx
	}

	summaryForUserRequest := &proto.SummaryForUserRequest{UserSelector: userSelector}
	summaryForUserResponse, err := ai.permissionClient.SummaryForUser(ctx, summaryForUserRequest)
	if err != nil || queryUserResponse.State != proto.State_SUCCESS {
		return ctx
	}

	user := queryUserResponse.Data.Users[0]
	userModel := new(model.User)
	userModel.LoadProto(user)

	permissions := []*model.RoleResource{}
	for _, summary := range summaryForUserResponse.Data {
		summaryMode := new(model.RoleResource)
		summaryMode.LoadProto(summary)
		permissions = append(permissions, summaryMode)
	}

	ctx = context.WithValue(ctx, ContextUserSymbol, userModel)
	ctx = context.WithValue(ctx, ContextSecretSymbol, secretModel)
	ctx = context.WithValue(ctx, ContextPermissionsSymbol, permissions)
	return ctx
}

func (ai *serverAuthInterceptor) ServerUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	newCtx := ai.resolveUserIncomingContext(ctx)
	resp, err = handler(newCtx, req)
	return resp, err
}

type AuthData struct {
	User        *model.User
	Secret      *model.Secret
	Permissions []*model.UserResourcePermissionSummary
}

func ResolveAuth(ctx context.Context) *AuthData {
	data := &AuthData{}
	user := ctx.Value(ContextUserSymbol)
	secret := ctx.Value(ContextSecretSymbol)
	permissions := ctx.Value(ContextPermissionsSymbol)

	if user != nil {
		if passUser, ok := user.(*model.User); ok {
			data.User = passUser
		}
	}

	if secret != nil {
		if passSecret, ok := secret.(*model.Secret); ok {
			data.Secret = passSecret
		}
	}

	if permissions != nil {
		if passPermissions, ok := permissions.([]*model.UserResourcePermissionSummary); ok {
			data.Permissions = passPermissions
		}
	}

	return data
}

func ResolvePermission(ctx context.Context, handler func(*model.User, *model.UserResourcePermissionSummary) bool) bool {
	authData := ResolveAuth(ctx)

	if authData == nil || authData.User == nil || authData.Secret == nil || authData.Permissions == nil {
		return false
	}

	if len(authData.Permissions) != 0 {
		for _, permission := range authData.Permissions {
			if permission.ResourceName == model.AllResourceName {
				return true
			}

			if handler(authData.User, permission) {
				return true
			}
		}
	}

	return false
}
