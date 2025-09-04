package externalapi

import (
	"github.com/cynx-io/ananke-reservation/internal/dependencies/config"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
	"time"
)

func NewGoogleRecaptchaClient() *recaptcha.ReCAPTCHA {
	captcha, err := recaptcha.NewReCAPTCHA(config.Config.Perintis.Google.RecaptchaSecretKey, recaptcha.V3, 10*time.Second)
	if err != nil {
		panic(err)
	}
	return &captcha
}
