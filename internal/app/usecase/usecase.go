package usecase

import (
	"github.com/pembajak/personal-finance/internal/app/usecase/user"
)

type UseCase struct {
	User user.UserUseCase
}
