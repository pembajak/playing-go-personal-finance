package finance

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

type Repository interface {
	// CreateFinance ...
	CreateFinance(ctx context.Context, finance models.Finance) (returnData models.Finance, err error)
	// UpdateFinance ...
	UpdateFinance(ctx context.Context, finance models.Finance) (returnData models.Finance, err error)
	// DeleteFinanceByID ...
	DeleteFinanceByID(ctx context.Context, id int64) (err error)
	// GetFinanceByID ...
	GetFinanceByID(ctx context.Context, id int64) (returnData models.Finance, err error)
	// GetAllFinance ...
	GetAllFinance(ctx context.Context, param models.GetAllFinanceReq) (returnData []models.Finance, err error)
	// CountTotalFinance ...
	CountTotalFinance(ctx context.Context, req models.GetAllFinanceReq) (total int64, err error)
}
