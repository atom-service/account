package server

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func TestMain(m *testing.M) {
	model.Init(context.TODO())
	m.Run()
}

type testServer struct {
	address string
}

func (t *testServer) CreateAccountClientWithToken(token string) proto.AccountServiceClient {
	authCredentials := grpc.WithPerRPCCredentials(&auth.AuthWithTokenCredentials{Token: token})
	nonSafeCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(t.address, authCredentials, nonSafeCredentials)
	if err != nil {
		panic(err)
	}

	return proto.NewAccountServiceClient(conn)
}

func createTestServer() *testServer {
	randAddress := ":" + strconv.Itoa(10000+rand.Intn(10000))
	go func() { StartServer(randAddress) }() // 如何 close
	return &testServer{address: randAddress}
}
