package externalapi

import (
	"context"
	brevo "github.com/getbrevo/brevo-go/lib"
)

func New(ctx context.Context) *brevo.APIClient {

	cfg := brevo.NewConfiguration()

	cfg.AddDefaultHeader("api-key", "YOUR_API_KEY")
	br := brevo.NewAPIClient(cfg)
	return br
}
