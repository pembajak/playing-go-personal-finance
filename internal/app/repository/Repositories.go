package repository

import (
	"github.com/pembajak/personal-finance/internal/app/repository/account"
	"github.com/pembajak/personal-finance/internal/app/repository/finance"
	"github.com/pembajak/personal-finance/internal/app/repository/user"
)

type Repositories struct {
	User    user.Repository
	Account account.Repository
	Finance finance.Repository
}
