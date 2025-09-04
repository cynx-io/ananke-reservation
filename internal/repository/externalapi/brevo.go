package externalapi

import (
	"github.com/cynx-io/ananke-reservation/internal/dependencies/config"
	brevo "github.com/getbrevo/brevo-go/lib"
)

func NewBrevoClient() *brevo.APIClient {

	cfg := brevo.NewConfiguration()

	cfg.AddDefaultHeader("api-key", config.Config.Perintis.Brevo.ApiKey)
	br := brevo.NewAPIClient(cfg)
	return br
}
