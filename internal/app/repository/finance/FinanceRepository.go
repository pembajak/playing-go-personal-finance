package finance

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

type repo struct {
	DB Repository
}

func NewRepo(db Repository) Repository {
	return &repo{
		DB: db,
	}
}

// CreateFinance ...
func (r *repo) CreateFinance(ctx context.Context, finance models.Finance) (returnData models.Finance, err error) {
	returnData, err = r.DB.CreateFinance(ctx, finance)
	if err != nil {
		return
	}

	return
}

// UpdateFinance ...
func (r *repo) UpdateFinance(ctx context.Context, finance models.Finance) (returnData models.Finance, err error) {
	returnData, err = r.DB.UpdateFinance(ctx, finance)
	if err != nil {
		return
	}

	return
}

// DeleteFinanceByID ...
func (r *repo) DeleteFinanceByID(ctx context.Context, id int64) (err error) {
	err = r.DB.DeleteFinanceByID(ctx, id)
	return
}

// GetFinanceByID ...
func (r *repo) GetFinanceByID(ctx context.Context, id int64) (returnData models.Finance, err error) {
	returnData, err = r.DB.GetFinanceByID(ctx, id)
	if err != nil {
		return
	}

	return
}

// GetAllFinance ...
func (r *repo) GetAllFinance(ctx context.Context, req models.GetAllFinanceReq) (returnData []models.Finance, err error) {
	returnData, err = r.DB.GetAllFinance(ctx, req)
	if err != nil {
		return
	}

	return
}

// CountTotalFinance ...
func (r *repo) CountTotalFinance(ctx context.Context, req models.GetAllFinanceReq) (total int64, err error) {
	total, err = r.DB.CountTotalFinance(ctx, req)
	if err != nil {
		return
	}

	return
}
