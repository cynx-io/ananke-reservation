package database

import (
	"context"
	"github.com/cynx-io/ananke-reservation/internal/model/entity"
	"gorm.io/gorm"
)

type TblWaitlistType struct {
	DB *gorm.DB
}

func NewWaitlistTypeRepo(DB *gorm.DB) *TblWaitlistType {
	return &TblWaitlistType{
		DB: DB,
	}
}

func (r *TblWaitlistType) GetWaitlistTypeById(ctx context.Context, id int32) (*entity.TblWaitlistType, error) {
	var existingWaitlistType entity.TblWaitlistType

	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existingWaitlistType).Error
	if err != nil {
		return nil, err
	}

	return &existingWaitlistType, nil
}
