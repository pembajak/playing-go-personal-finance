package user

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

type UserUseCase interface {
	// CreateUser ...
	CreateUser(ctx context.Context, param models.User) (returnData models.User, err error)

	// Login ...
	Login(ctx context.Context, param models.User) (returnData models.LoginRes, err error)

	// Profile ...
	Profile(ctx context.Context, id int64) (returnData models.Profile, err error)
}
