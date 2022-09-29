package user

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
)

type repo struct {
	DB Repository
}

func NewRepo(db Repository) Repository {
	return &repo{
		db,
	}
}

// CreateUser ...
func (r *repo) CreateUser(ctx context.Context, user models.User) (returnData models.User, err error) {
	returnData, err = r.DB.CreateUser(ctx, user)
	if err != nil {
		return
	}

	return
}

// GetUserByEmail ...
func (r *repo) GetUserByEmail(ctx context.Context, user models.User) (returnData models.User, err error) {
	returnData, err = r.DB.GetUserByEmail(ctx, user)
	if err != nil {
		return
	}

	return
}

// GetUserByID ...
func (r *repo) GetUserByID(ctx context.Context, id int64) (returnData models.User, err error) {
	returnData, err = r.DB.GetUserByID(ctx, id)
	if err != nil {
		return
	}

	return
}
