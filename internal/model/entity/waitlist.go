package entity

import (
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	"github.com/cynx-io/cynx-core/src/entity"
)

type TblWaitlist struct {
	entity.EssentialEntity
	Email          string          `gorm:"type:varchar(255);not null" json:"email"`
	WaitlistType   TblWaitlistType `gorm:"foreignkey:WaitlistTypeId" json:"waitlist_type"`
	WaitlistTypeId int32           `gorm:"not null" json:"waitlist_type_id"`
}

func (p TblWaitlist) Response() *proto.Waitlist {
	return &proto.Waitlist{
		Id:             p.Id,
		Email:          p.Email,
		WaitlistTypeId: p.WaitlistTypeId,
	}

}
