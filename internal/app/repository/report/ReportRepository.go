package report

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

// GetTotalTransactionDaily ...
func (r *repo) GetTotalTransactionDaily(ctx context.Context, req models.GetTotalTransaction) (returnData []models.TotalTransactionRes, err error) {
	returnData, err = r.DB.GetTotalTransactionDaily(ctx, req)
	if err != nil {
		return
	}
	return
}

// GetTotalTransactionMonthly ...
func (r *repo) GetTotalTransactionMonthly(ctx context.Context, req models.GetTotalTransaction) (returnData []models.TotalTransactionMonthlyRes, err error) {
	returnData, err = r.DB.GetTotalTransactionMonthly(ctx, req)
	if err != nil {
		return
	}

	return
}
