package usecase

import (
	"github.com/pembajak/personal-finance/internal/app/usecase/account"
	"github.com/pembajak/personal-finance/internal/app/usecase/finance"
	"github.com/pembajak/personal-finance/internal/app/usecase/user"
)

type UseCase struct {
	User    user.UserUseCase
	Account account.AccountUseCase
	Finance finance.FinanceUseCase
}
