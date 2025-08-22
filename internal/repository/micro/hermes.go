package micro

import (
	pbhermes "github.com/cynx-io/ananke-reservation/api/proto/gen/hermes"
	"github.com/cynx-io/ananke-reservation/internal/dependencies/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHermesUserClient() pbhermes.HermesUserServiceClient {

	conn, err := grpc.NewClient(config.Config.Hermes.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Hermes gRPC server: " + err.Error())
	}

	userClient := pbhermes.NewHermesUserServiceClient(conn)
	return userClient
}
