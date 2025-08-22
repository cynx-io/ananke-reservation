package app

import (
	paymentservice "github.com/cynx-io/ananke-reservation/internal/service/preorderservice"
)

type Services struct {
	PaymentService *paymentservice.Service
}

func NewServices(repos *Repos) *Services {
	return &Services{
		PaymentService: &paymentservice.Service{
			TblPreorder:         repos.PreorderRepo,
			TblPreorderType:     repos.PreorderTypeRepo,
			HermesUserClient:    repos.HermesUserClient,
			PlutusPaymentClient: repos.PlutusPaymentClient,
		},
	}
}
