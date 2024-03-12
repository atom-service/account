package server

import (
	"context"

	"github.com/atom-service/account/package/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SecretServer struct {
	protos.UnimplementedSecretServiceServer
}

func NewSecretServer() *SecretServer {
	return &SecretServer{}
}

func (s *SecretServer) CreateSecret(ctx context.Context, request *protos.CreateSecretRequest) (result *protos.CreateSecretResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}

func (s *SecretServer) DisableSecret(ctx context.Context, request *protos.DisableSecretRequest) (result *protos.DisableSecretResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableSecret not implemented")
}

func (s *SecretServer) DeleteSecret(ctx context.Context, request *protos.DeleteSecreteRequest) (result *protos.DeleteSecreteResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}

func (s *SecretServer) QuerySecrets(ctx context.Context, request *protos.QuerySecretsRequest) (result *protos.QuerySecretsResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySecrets not implemented")
}
