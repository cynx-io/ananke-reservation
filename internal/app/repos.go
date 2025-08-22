package app

import (
	pbhermes "github.com/cynx-io/ananke-reservation/api/proto/gen/hermes"
	pbplutus "github.com/cynx-io/ananke-reservation/api/proto/gen/plutus"
	"github.com/cynx-io/ananke-reservation/internal/repository/database"
	"github.com/cynx-io/ananke-reservation/internal/repository/micro"
)

type Repos struct {
	PreorderRepo     *database.TblPreorder
	PreorderTypeRepo *database.TblPreorderType

	HermesUserClient    pbhermes.HermesUserServiceClient
	PlutusPaymentClient pbplutus.PaymentServiceClient
}

func NewRepos(dependencies *Dependencies) *Repos {
	return &Repos{
		PreorderRepo:     database.NewPreorderRepo(dependencies.DatabaseClient.Db),
		PreorderTypeRepo: database.NewPreorderTypeRepo(dependencies.DatabaseClient.Db),

		HermesUserClient:    micro.NewHermesUserClient(),
		PlutusPaymentClient: micro.NewPlutusPaymentClient(),
	}
}
