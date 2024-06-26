package server

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/atom-service/account/internal/config"
	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain(m *testing.M) {
	config.MustInit("../../config.yaml")
	model.Init(context.TODO())
	m.Run()
}

type testServer struct {
	address string
}

func (t *testServer) CreateClientConn(token string) *grpc.ClientConn {
	authCredentials := grpc.WithPerRPCCredentials(auth.NewTokenCredential(token))
	nonSafeCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(t.address, authCredentials, nonSafeCredentials)
	if err != nil {
		panic(err)
	}

	return conn
}

func (t *testServer) CreateAccountClientWithToken(token string) proto.AccountServiceClient {
	return proto.NewAccountServiceClient(t.CreateClientConn(token))
}

func (t *testServer) CreatePermissionClientWithToken(token string) proto.PermissionServiceClient {
	return proto.NewPermissionServiceClient(t.CreateClientConn(token))
}

func createTestServer() *testServer {
	randAddress := ":" + strconv.Itoa(10000+rand.Intn(10000))
	go func() { StartServer(randAddress) }() // 如何 close
	return &testServer{address: randAddress}
}
