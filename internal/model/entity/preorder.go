package entity

import (
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	"github.com/cynx-io/cynx-core/src/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TblPreorder struct {
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	entity.EssentialEntity
	PaymentInvoiceId string          `gorm:"size:255;not null" json:"payment_invoice_id"`
	PaymentLinkUrl   string          `gorm:"size:255;not null" json:"payment_link_url"`
	Currency         string          `gorm:"size:10;not null" json:"currency"`
	Description      string          `gorm:"size:255;not null" json:"description"`
	Provider         string          `gorm:"size:50;not null" json:"provider"`
	PreorderType     TblPreorderType `gorm:"foreignkey:PreorderTypeId" json:"preorder_type"`
	Status           int32           `json:"status" gorm:"type:tinyint(4);not null"`
	Amount           float32         `gorm:"not null" json:"amount"`
	PreorderTypeId   int32           `gorm:"not null" json:"preorder_type_id"`
	UserId           int32           `gorm:"not null" json:"user_id"`
}

func (p TblPreorder) Response() *proto.Preorder {
	return &proto.Preorder{
		Id:                p.Id,
		PaymentInvoiceId:  p.PaymentInvoiceId,
		TransactionStatus: proto.TransactionStatus(p.Status),
		PaymentLinkUrl:    p.PaymentLinkUrl,
		ExpiresAt:         timestamppb.New(p.ExpiresAt),
		Amount:            p.Amount,
		Currency:          p.Currency,
		Description:       p.Description,
		Provider:          p.Provider,
		PreorderTypeId:    p.PreorderTypeId,
		UserId:            p.UserId,
	}

}
