package user

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

// Repository ...
type Repository interface {
	// CreateUser ...
	CreateUser(ctx context.Context, user models.User) (returnData models.User, err error)

	// GetUserByEmail ...
	GetUserByEmail(ctx context.Context, user models.User) (returnData models.User, err error)

	// GetUserByID ...
	GetUserByID(ctx context.Context, id int64) (returnData models.User, err error)
}
