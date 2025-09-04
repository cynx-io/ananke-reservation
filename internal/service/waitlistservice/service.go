package waitlistservice

import (
	pbhermes "github.com/cynx-io/ananke-reservation/api/proto/gen/hermes"
	"github.com/cynx-io/ananke-reservation/internal/repository/database"
	brevo "github.com/getbrevo/brevo-go/lib"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

type Service struct {
	TblWaitlist     *database.TblWaitlist
	TblWaitlistType *database.TblWaitlistType

	HermesUserClient pbhermes.HermesUserServiceClient

	BrevoClient     *brevo.APIClient
	RecaptchaClient *recaptcha.ReCAPTCHA
}
