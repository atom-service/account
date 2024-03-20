package auth

import (
	"context"
	"time"

	"github.com/atom-service/account/internal/model"
	protos "github.com/atom-service/account/package/protos"
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
			ExpiresAt: time.Now().Add(10 * time.Minute),
		}),
	}, nil
}

func (x *AuthWithSecretCredentials) RequireTransportSecurity() bool {
	return false
}

type serverAuthInterceptor struct {
	secretKey        string
	secretValue      string
	accountClient    protos.AccountServiceClient
	permissionClient protos.PermissionServiceClient
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
		accountClient:    protos.NewAccountServiceClient(conn),
		permissionClient: protos.NewPermissionServiceClient(conn),
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
		ctx = context.WithValue(ctx, ContextSecretSymbol, &protos.Secret{
			Key:   ai.secretKey,
			Value: ai.secretValue,
		})
		return ctx
	}

	paginationLimit := int64(1)
	secretSelector := &protos.SecretSelector{Key: &tokenInfo.SecretKey}
	paginationOption := &protos.PaginationOption{Limit: &paginationLimit}
	querySecretsRequest := &protos.QuerySecretsRequest{Selector: secretSelector, Pagination: paginationOption}
	querySecretsResponse, err := ai.accountClient.QuerySecrets(ctx, querySecretsRequest)
	if err != nil || querySecretsResponse.State != protos.State_SUCCESS {
		return ctx
	}
	if querySecretsResponse.Data.Total == 0 {
		return ctx
	}

	// 验证 token 是否有效
	secret := querySecretsResponse.Data.Secrets[0]

	// 是否已经被禁用
	secretModel := new(model.Secret)
	secretModel.LoadProtoStruct(secret)
	if secretModel.IsDisabled() {
		return ctx
	}

	if !VerifyToken(secret.Key, secret.Value, tokens[0]) {
		return ctx
	}

	userSelector := &protos.UserSelector{ID: &secret.UserID}
	queryUserResponse, err := ai.accountClient.QueryUsers(ctx, &protos.QueryUsersRequest{Selector: userSelector})
	if err != nil || queryUserResponse.State != protos.State_SUCCESS || querySecretsResponse.Data.Total == 0 {
		return ctx
	}

	summaryForUserRequest := &protos.SummaryForUserRequest{UserSelector: userSelector}
	summaryForUserResponse, err := ai.permissionClient.SummaryForUser(ctx, summaryForUserRequest)
	if err != nil || queryUserResponse.State != protos.State_SUCCESS {
		return ctx
	}

	user := queryUserResponse.Data.Users[0]
	userModel := new(model.User)
	userModel.LoadProtoStruct(user)

	permissions := []*model.UserResourcePermissionSummary{}
	for _, summary := range summaryForUserResponse.Data {
		summaryMode := new(model.UserResourcePermissionSummary)
		summaryMode.LoadProtoStruct(summary)
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

func ResolveAuthFromIncomingContext(ctx context.Context) *AuthData {
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
		if passPermissions, ok := permissions.([]any); ok {
			data.Permissions = make([]*model.UserResourcePermissionSummary, len(passPermissions))
			for _, passPermission := range passPermissions {
				if passUserResourceSummary, ok := passPermission.(*model.UserResourcePermissionSummary); ok {
					data.Permissions = append(data.Permissions, passUserResourceSummary)
				}
			}
		}
	}

	return data
}

func ResolvePermissionFromIncomingContext(ctx context.Context, handler func(*model.User, *model.UserResourcePermissionSummary) bool) bool {
	authData := ResolveAuthFromIncomingContext(ctx)

	// 不一定都有，部分信息也可以处理
	if authData == nil || authData.User == nil || authData.Secret == nil {
		return false
	}

	userID := authData.User.ID

	roles, err := model.Permission.QueryUserResourceSummaries(ctx, model.UserResourceSummarySelector{UserID: *userID})

	if err != nil || len(roles) == 0 {
		return false
	}

	for _, role := range roles {
		if handler(authData.User, role) {
			return true
		}
	}

	return false
}
