package database

import (
	"context"
	proto "github.com/cynx-io/ananke-reservation/api/proto/gen/ananke"
	"github.com/cynx-io/ananke-reservation/internal/model/entity"
	"gorm.io/gorm"
)

type TblPreorder struct {
	DB *gorm.DB
}

func NewPreorderRepo(DB *gorm.DB) *TblPreorder {
	return &TblPreorder{
		DB: DB,
	}
}

func (r *TblPreorder) GetPreorderById(ctx context.Context, id int32) (*entity.TblPreorder, error) {
	var existingPreorder entity.TblPreorder

	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existingPreorder).Error
	if err != nil {
		return nil, err
	}

	return &existingPreorder, nil
}

func (r *TblPreorder) GetPreorderByInvoiceId(ctx context.Context, invoiceId string) (*entity.TblPreorder, error) {
	var existingPreorder entity.TblPreorder

	err := r.DB.WithContext(ctx).Where("payment_invoice_id = ?", invoiceId).First(&existingPreorder).Error
	if err != nil {
		return nil, err
	}

	return &existingPreorder, nil
}

func (r *TblPreorder) CreatePreorder(ctx context.Context, preorder *entity.TblPreorder) error {
	err := r.DB.WithContext(ctx).Create(preorder).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TblPreorder) UpdatePreorderStatus(ctx context.Context, id int32, status int32) error {
	result := r.DB.WithContext(ctx).
		Model(&entity.TblPreorder{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TblPreorder) GetCompletedPreorderByUserIdAndType(ctx context.Context, userId int32, preorderTypeId int32) (*entity.TblPreorder, error) {
	var preorders *entity.TblPreorder

	err := r.DB.WithContext(ctx).
		Where("user_id = ? AND preorder_type_id = ? AND status = ?", userId, preorderTypeId, int32(proto.TransactionStatus_TRANSACTION_STATUS_COMPLETED)).
		Order("created_date DESC").
		First(&preorders).Error
	if err != nil {
		return nil, err
	}

	return preorders, nil
}

func (r *TblPreorder) GetPendingPreorderByUserIdAndType(ctx context.Context, userId int32, preorderTypeId int32) (*entity.TblPreorder, error) {
	var preorders *entity.TblPreorder

	err := r.DB.WithContext(ctx).
		Where("user_id = ? AND preorder_type_id = ? AND expires_at > NOW() AND status = ?", userId, preorderTypeId, int32(proto.TransactionStatus_TRANSACTION_STATUS_PENDING)).
		Order("created_date DESC").
		First(&preorders).Error
	if err != nil {
		return nil, err
	}

	return preorders, nil
}
