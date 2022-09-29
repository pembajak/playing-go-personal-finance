package user

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/pkg/driver/driversql"
)

type db struct {
	DB *driversql.Database
}

// NewDB ...
func NewDB(d *driversql.Database) Repository {
	return &db{
		DB: d,
	}
}

// CreateUser ...
func (d *db) CreateUser(ctx context.Context, user models.User) (returnData models.User, err error) {
	err = d.DB.Instance.WithContext(ctx).Create(&user).Error
	if err != nil {
		return
	}

	returnData = models.User{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}

// GetUserByEmail ...
func (d *db) GetUserByEmail(ctx context.Context, user models.User) (returnData models.User, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("email = ?", user.Email).Find(&returnData).Error
	if err != nil {
		return
	}

	return
}

// GetUserByID ...
func (d *db) GetUserByID(ctx context.Context, id int64) (returnData models.User, err error) {
	err = d.DB.Instance.WithContext(ctx).Where("id = ?", id).Find(&returnData).Error
	if err != nil {
		return
	}
	return
}
