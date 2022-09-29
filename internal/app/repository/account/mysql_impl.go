package account

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

// CreateAccount ...
func (d *db) CreateAccount(ctx context.Context, account models.Account) (returnData models.Account, err error) {
	err = d.DB.Instance.WithContext(ctx).Create(&account).Error
	if err != nil {
		return
	}

	returnData = models.Account{
		ID:          account.ID,
		Name:        account.Name,
		Type:        account.Type,
		Description: account.Description,
		UserID:      account.UserID,
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}

	return
}

// UpdateAccount ...
func (d *db) UpdateAccount(ctx context.Context, account models.Account) (returnData models.Account, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("id =?", account.ID).Updates(&account).Error
	if err != nil {
		return
	}

	returnData = models.Account{
		ID:          account.ID,
		Name:        account.Name,
		Type:        account.Type,
		Description: account.Description,
		UserID:      account.UserID,
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}

	return
}

// GetAccountByID ...
func (d *db) GetAccountByID(ctx context.Context, id int64) (returnData models.Account, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("id =?", id).First(&returnData).Error
	if err != nil {
		return
	}

	return
}

// GetAllAccount ...
func (d *db) GetAllAccount(ctx context.Context, req models.GetAllAccountReq) (returnData []models.Account, err error) {
	query := d.DB.Instance.WithContext(ctx).Table("accounts").Where("user_id =?", req.UserID)
	if req != (models.GetAllAccountReq{}) {
		offset := req.Limit * (req.Page - 1)
		query = query.Offset(offset).Limit(req.Limit)
	}

	err = query.Find(&returnData).Error

	if err != nil {
		return
	}

	return
}

// DeleteAccount ...
func (d *db) DeleteAccount(ctx context.Context, id int64) (returnData models.Account, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("id = ?", id).Delete(&models.Account{}).Error
	if err != nil {
		return
	}
	return
}

// CountTotalAccount ...
func (d *db) CountTotalAccount(ctx context.Context, userID int64) (total int64, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("user_id =?", userID).Table("accounts").Where("deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return
	}

	return
}
