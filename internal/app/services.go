package app

import (
	paymentservice "github.com/cynx-io/ananke-reservation/internal/service/preorderservice"
	"github.com/cynx-io/ananke-reservation/internal/service/waitlistservice"
)

type Services struct {
	PaymentService  *paymentservice.Service
	WaitlistService *waitlistservice.Service
}

func NewServices(repos *Repos) *Services {
	return &Services{
		PaymentService: &paymentservice.Service{
			TblPreorder:         repos.PreorderRepo,
			TblPreorderType:     repos.PreorderTypeRepo,
			HermesUserClient:    repos.HermesUserClient,
			PlutusPaymentClient: repos.PlutusPaymentClient,
		},
		WaitlistService: &waitlistservice.Service{
			TblWaitlist:      repos.WaitlistRepo,
			TblWaitlistType:  repos.WaitlistTypeRepo,
			HermesUserClient: repos.HermesUserClient,
			BrevoClient:      repos.BrevoClient,
			RecaptchaClient:  repos.RecaptchaClient,
		},
	}
}
