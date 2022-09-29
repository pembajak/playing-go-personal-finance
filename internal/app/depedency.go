package app

import (
	"github.com/pembajak/personal-finance/internal/app/repository"
	userRepository "github.com/pembajak/personal-finance/internal/app/repository/user"
	"github.com/pembajak/personal-finance/internal/app/usecase"
	UserUseCase "github.com/pembajak/personal-finance/internal/app/usecase/user"
	"github.com/pembajak/personal-finance/internal/pkg/driver/driversql"
	"github.com/pembajak/personal-finance/internal/pkg/token"
)

func WiringRepository(db *driversql.Database) *repository.Repositories {
	return &repository.Repositories{
		User: userRepository.NewRepo(userRepository.NewDB(db)),
	}
}

func WiringUsecase(repo *repository.Repositories, token token.IToken) *usecase.UseCase {
	return &usecase.UseCase{
		User: UserUseCase.NewUsrCase(repo, token),
	}
}
