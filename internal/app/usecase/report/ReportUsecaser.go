package report

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

type ReportUseCase interface {
	GetTotalTransactionDaily(ctx context.Context, req models.GetTotalTransaction) (returnData interface{}, err error)
	GetTotalTransactionMonthly(ctx context.Context, req models.GetTotalTransaction) (returnData interface{}, err error)
}
