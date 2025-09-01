package waitlistservice

import (
	pbhermes "github.com/cynx-io/ananke-reservation/api/proto/gen/hermes"
	"github.com/cynx-io/ananke-reservation/internal/repository/database"
)

type Service struct {
	TblWaitlist     *database.TblWaitlist
	TblWaitlistType *database.TblWaitlistType

	HermesUserClient pbhermes.HermesUserServiceClient
}
