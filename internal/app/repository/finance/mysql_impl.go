package finance

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/pkg/driver/driversql"
)

type db struct {
	DB *driversql.Database
}

// NewDB ...
func NewDB(d *driversql.Database) Repository {
	return &db{
		DB: d,
	}
}

// CreateFinance ...
func (d *db) CreateFinance(ctx context.Context, finance models.Finance) (returnData models.Finance, err error) {
	err = d.DB.Instance.WithContext(ctx).Create(&finance).Error
	if err != nil {
		return
	}

	returnData = models.Finance{
		ID:          finance.ID,
		Title:       finance.Title,
		Description: finance.Description,
		Amount:      finance.Amount,
		UserID:      finance.UserID,
		AccountID:   finance.AccountID,
		CreatedAt:   finance.CreatedAt,
		UpdatedAt:   finance.UpdatedAt,
	}

	return
}

// UpdateFinance ...
func (d *db) UpdateFinance(ctx context.Context, finance models.Finance) (returnData models.Finance, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("id =?", finance.ID).Updates(&finance).Error
	if err != nil {
		return
	}

	return
}

// GetFinanceByID ...
func (d *db) GetFinanceByID(ctx context.Context, id int64) (returnData models.Finance, err error) {
	err = d.DB.Instance.WithContext(ctx).Preload("Account").Where("id =?", id).First(&returnData).Error

	if err != nil {
		return
	}

	return
}

// DeleteFinanceByID ...
func (d *db) DeleteFinanceByID(ctx context.Context, id int64) (err error) {
	err = d.DB.Instance.WithContext(ctx).Where("id =?", id).Delete(&models.Finance{}).Error
	return
}
