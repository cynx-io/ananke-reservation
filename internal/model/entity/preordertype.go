package entity

import (
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	core "github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/cynx-core/src/entity"
)

type TblPreorderType struct {
	entity.EssentialEntity
	Name         string  `json:"name" gorm:"type:varchar(255);not null"`
	Description  string  `json:"description" gorm:"type:varchar(255);"`
	Currency     string  `json:"currency" gorm:"type:varchar(10);not null"`
	Amount       float32 `json:"amount" gorm:"type:decimal(10,2);not null"`
	ChannelType  int32   `json:"channel_type" gorm:"type:tinyint(4);not null"`
	ActiveStatus int32   `json:"active_status" gorm:"type:tinyint(4);not null"`
}

func (u TblPreorderType) Response() *proto.PreorderType {
	return &proto.PreorderType{
		Id:           u.Id,
		Name:         u.Name,
		Description:  u.Description,
		Amount:       u.Amount,
		Currency:     u.Currency,
		ChannelType:  core.ChannelType(u.ChannelType),
		ActiveStatus: proto.ActiveStatus(u.ActiveStatus),
	}
}
