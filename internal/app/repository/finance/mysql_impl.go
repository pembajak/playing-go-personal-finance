package finance

import (
	"context"
	"strings"

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

// GetAllFinance ...
func (d *db) GetAllFinance(ctx context.Context, req models.GetAllFinanceReq) (returnData []models.Finance, err error) {
	query := d.DB.Instance.WithContext(ctx).Preload("Account").Where("user_id =?", req.UserID)
	if req != (models.GetAllFinanceReq{}) {
		offset := req.Limit * (req.Page - 1)
		query = query.Offset(offset).Limit(req.Limit)
	}

	if strings.Compare(req.Title, "") != 0 {
		searchTitle := strings.ToLower(req.Title)
		query = query.Where("LOWER(title) LIKE ?", "%"+searchTitle+"%")
	}

	if strings.Compare(req.Type, "") != 0 {
		query = query.Where("type =?", req.Type)
	}

	if strings.Compare(req.StartDate, "") != 0 && strings.Compare(req.EndDate, "") != 0 {
		query = query.Where("transaction_date >= ?", req.StartDate).Where("transaction_date < ?", req.EndDate)
	}

	err = query.Find(&returnData).Error

	if err != nil {
		return
	}

	return
}

// CountTotalFinance ...
func (d *db) CountTotalFinance(ctx context.Context, req models.GetAllFinanceReq) (total int64, err error) {
	query := d.DB.Instance.WithContext(ctx).Where("user_id =?", req.UserID).Table("finances").Where("deleted_at IS NULL")

	if strings.Compare(req.Title, "") != 0 {
		searchTitle := strings.ToLower(req.Title)
		query = query.Where("LOWER(title) LIKE ?", "%"+searchTitle+"%")
	}

	if strings.Compare(req.Type, "") != 0 {
		query = query.Where("type =?", req.Type)
	}

	if strings.Compare(req.StartDate, "") != 0 && strings.Compare(req.EndDate, "") != 0 {
		query = query.Where("transaction_date >= ?", req.StartDate).Where("transaction_date < ?", req.EndDate)
	}

	err = query.Count(&total).Error
	if err != nil {
		return
	}
	return
}
