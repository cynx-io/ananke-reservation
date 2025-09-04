package app

import (
	pbhermes "github.com/cynx-io/ananke-reservation/api/proto/gen/hermes"
	pbplutus "github.com/cynx-io/ananke-reservation/api/proto/gen/plutus"
	"github.com/cynx-io/ananke-reservation/internal/repository/database"
	"github.com/cynx-io/ananke-reservation/internal/repository/externalapi"
	"github.com/cynx-io/ananke-reservation/internal/repository/micro"
	brevo "github.com/getbrevo/brevo-go/lib"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

type Repos struct {
	PreorderRepo     *database.TblPreorder
	PreorderTypeRepo *database.TblPreorderType
	WaitlistRepo     *database.TblWaitlist
	WaitlistTypeRepo *database.TblWaitlistType

	HermesUserClient    pbhermes.HermesUserServiceClient
	PlutusPaymentClient pbplutus.PaymentServiceClient

	BrevoClient     *brevo.APIClient
	RecaptchaClient *recaptcha.ReCAPTCHA
}

func NewRepos(dependencies *Dependencies) *Repos {
	return &Repos{
		PreorderRepo:     database.NewPreorderRepo(dependencies.DatabaseClient.Db),
		PreorderTypeRepo: database.NewPreorderTypeRepo(dependencies.DatabaseClient.Db),
		WaitlistRepo:     database.NewWaitlistRepo(dependencies.DatabaseClient.Db),
		WaitlistTypeRepo: database.NewWaitlistTypeRepo(dependencies.DatabaseClient.Db),

		HermesUserClient:    micro.NewHermesUserClient(),
		PlutusPaymentClient: micro.NewPlutusPaymentClient(),

		BrevoClient:     externalapi.NewBrevoClient(),
		RecaptchaClient: externalapi.NewGoogleRecaptchaClient(),
	}
}
