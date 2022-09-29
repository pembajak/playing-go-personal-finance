package finance

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

type FinanceUseCase interface {
	CreateFinance(ctx context.Context, finance models.Finance) (returnData models.Finance, err error)
	UpdateFinance(ctx context.Context, finance models.Finance) (returnData models.Finance, err error)
	DeleteFinanceByID(ctx context.Context, id int64) (returnData models.Finance, err error)
	GetFinanceByID(ctx context.Context, id int64) (returnData models.Finance, err error)
}
