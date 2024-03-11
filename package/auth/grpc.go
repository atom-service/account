package auth

import (
	"context"

	"github.com/atom-service/account/internal/model"
	protos "github.com/atom-service/account/package/protos"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

type AuthCredentials struct {
	Token     string
	SecretID  string
	SecretKey string
}

func (x *AuthCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": x.Token,
	}, nil
}

func (x *AuthCredentials) RequireTransportSecurity() bool {
	return false
}

type AuthInterceptor struct {
	accountClient    protos.AccountServiceClient
	permissionClient protos.PermissionServiceClient
}

func (ai *AuthInterceptor) resolveUserIncomingContext(ctx context.Context) context.Context {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}
	token, ok := metadata["authorization"]
	if !ok || len(token) == 0 {
		return ctx
	}

	// 查询用户信息
	ai.accountClient.QueryUsers(ctx, &protos.QueryUsersRequest{

	})

	return context.WithValue(ctx, "auth", nil)
}

func (ai *AuthInterceptor) ServerUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	newCtx := ai.resolveUserIncomingContext(ctx)
	resp, err = handler(newCtx, req)
	return resp, err
}

func NewAuthInterceptor(address string) *AuthInterceptor {
	credentials := grpc.WithPerRPCCredentials(&AuthCredentials{})
	conn, err := grpc.Dial(address, credentials)
	if err != nil {
		panic(err)
	}

	return &AuthInterceptor{
		protos.NewAccountServiceClient(conn),
		protos.NewPermissionServiceClient(conn),
	}
}

func ResolveUserFromIncomingContext(ctx context.Context) (*model.User) {
	// TODO
	return nil
}
