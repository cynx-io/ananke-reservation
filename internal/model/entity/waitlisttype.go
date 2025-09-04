package entity

import (
	"encoding/json"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	core "github.com/cynx-io/cynx-core/proto/gen"
	"github.com/cynx-io/cynx-core/src/entity"
)

type TblWaitlistType struct {
	entity.EssentialEntity
	Name         string `json:"name" gorm:"type:varchar(255);not null"`
	Description  string `json:"description" gorm:"type:varchar(255);"`
	ListIds      string `json:"list_ids" gorm:"type:varchar(255);"`
	ChannelType  int32  `json:"channel_type" gorm:"type:tinyint(4);not null"`
	ActiveStatus int32  `json:"active_status" gorm:"type:tinyint(4);not null"`
}

func (u TblWaitlistType) GetListIds() []int64 {
	var ids []int64
	if u.ListIds == "" || u.ListIds == "[]" {
		return ids
	}
	if len(u.ListIds) <= 2 {
		return ids
	}

	err := json.Unmarshal([]byte(u.ListIds), &ids)
	if err != nil {
		return ids
	}

	return ids
}

func (u TblWaitlistType) Response() *proto.WaitlistType {
	return &proto.WaitlistType{
		Id:           u.Id,
		Name:         u.Name,
		Description:  u.Description,
		ChannelType:  core.ChannelType(u.ChannelType),
		ActiveStatus: proto.ActiveStatus(u.ActiveStatus),
	}
}
