package database

import (
	"context"
	"github.com/cynx-io/ananke-reservation/internal/model/entity"
	"gorm.io/gorm"
)

type TblPreorderType struct {
	DB *gorm.DB
}

func NewPreorderTypeRepo(DB *gorm.DB) *TblPreorderType {
	return &TblPreorderType{
		DB: DB,
	}
}

func (r *TblPreorderType) GetPreorderTypeById(ctx context.Context, id int32) (*entity.TblPreorderType, error) {
	var existingPreorderType entity.TblPreorderType

	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existingPreorderType).Error
	if err != nil {
		return nil, err
	}

	return &existingPreorderType, nil
}
