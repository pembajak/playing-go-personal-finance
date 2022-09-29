package report

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

type Repository interface {
	// GetTotalTransactionDaily ...
	GetTotalTransactionDaily(ctx context.Context, req models.GetTotalTransaction) (returnData []models.TotalTransactionRes, err error)

	// GetTotalTransactionMonthly ...
	GetTotalTransactionMonthly(ctx context.Context, req models.GetTotalTransaction) (returnData []models.TotalTransactionMonthlyRes, err error)
}
