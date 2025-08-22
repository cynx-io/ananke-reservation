package micro

import (
	pbplutus "github.com/cynx-io/ananke-reservation/api/proto/gen/plutus"
	"github.com/cynx-io/ananke-reservation/internal/dependencies/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPlutusPaymentClient() pbplutus.PaymentServiceClient {

	conn, err := grpc.NewClient(config.Config.Plutus.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to Plutus gRPC server: " + err.Error())
	}

	client := pbplutus.NewPaymentServiceClient(conn)
	return client
}
