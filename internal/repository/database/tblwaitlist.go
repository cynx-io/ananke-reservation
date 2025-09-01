package database

import (
	"context"
	"github.com/cynx-io/ananke-reservation/internal/model/entity"
	"gorm.io/gorm"
)

type TblWaitlist struct {
	DB *gorm.DB
}

func NewWaitlistRepo(DB *gorm.DB) *TblWaitlist {
	return &TblWaitlist{
		DB: DB,
	}
}

func (r *TblWaitlist) GetWaitlistById(ctx context.Context, id int32) (*entity.TblWaitlist, error) {
	var existingWaitlist entity.TblWaitlist

	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&existingWaitlist).Error
	if err != nil {
		return nil, err
	}

	return &existingWaitlist, nil
}

func (r *TblWaitlist) GetWaitlistByEmail(ctx context.Context, email string) (*entity.TblWaitlist, error) {
	var existingWaitlist entity.TblWaitlist

	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&existingWaitlist).Error
	if err != nil {
		return nil, err
	}

	return &existingWaitlist, nil
}

func (r *TblWaitlist) CreateWaitlist(ctx context.Context, waitlist *entity.TblWaitlist) error {
	err := r.DB.WithContext(ctx).Create(waitlist).Error
	if err != nil {
		return err
	}

	return nil
}
