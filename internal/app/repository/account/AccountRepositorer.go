package account

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

type Repository interface {
	// CreateAccount ...
	CreateAccount(ctx context.Context, account models.Account) (returnData models.Account, err error)

	// UpdateAccount ...
	UpdateAccount(ctx context.Context, account models.Account) (returnData models.Account, err error)

	// GetAccountByID ...
	GetAccountByID(ctx context.Context, id int64) (returnData models.Account, err error)

	// GetAllAccount ...
	GetAllAccount(ctx context.Context, req models.GetAllAccountReq) (returnData []models.Account, err error)

	// DeleteAccount ...
	DeleteAccount(ctx context.Context, id int64) (returnData models.Account, err error)

	// CountTotalAccount ...
	CountTotalAccount(ctx context.Context, userID int64) (total int64, err error)
}
