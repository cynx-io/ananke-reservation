package paymentservice

import (
	pbhermes "github.com/cynx-io/ananke-reservation/api/proto/gen/hermes"
	pbplutus "github.com/cynx-io/ananke-reservation/api/proto/gen/plutus"
	"github.com/cynx-io/ananke-reservation/internal/repository/database"
)

type Service struct {
	TblPreorder     *database.TblPreorder
	TblPreorderType *database.TblPreorderType

	HermesUserClient    pbhermes.HermesUserServiceClient
	PlutusPaymentClient pbplutus.PaymentServiceClient
}
