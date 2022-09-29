package user

import (
	"context"
	"errors"
	"time"

	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/app/repository"
	pkgError "github.com/pembajak/personal-finance/internal/pkg/errors"
	"github.com/pembajak/personal-finance/internal/pkg/token"
	"github.com/pembajak/personal-finance/internal/pkg/utils"
	"github.com/ulule/deepcopier"
)

// srv ...
type srv struct {
	repo  *repository.Repositories
	token token.IToken
}

// NewSrv ..
func NewUsrCase(repo *repository.Repositories, token token.IToken) UserUseCase {
	return &srv{
		repo:  repo,
		token: token,
	}
}

// CreateUser ...
func (s *srv) CreateUser(ctx context.Context, param models.User) (returnData models.User, err error) {
	userRepo := models.User{}

	// encrypt password
	password, err := utils.EncryptString(param.Password)
	if err != nil {
		return
	}

	param.Password = password
	_ = deepcopier.Copy(param).To(&userRepo)

	res, err := s.repo.User.CreateUser(ctx, userRepo)
	if err != nil {
		return
	}
	_ = deepcopier.Copy(res).To(&returnData)

	return
}

func (s *srv) Login(ctx context.Context, param models.User) (returnData models.LoginRes, err error) {
	var userRepo models.User
	_ = deepcopier.Copy(param).To(&userRepo)

	res, err := s.repo.User.GetUserByEmail(ctx, userRepo)
	if err != nil {
		return
	}

	validPassword, err := utils.CompareStringValid(param.Password, res.Password)
	if !validPassword {
		err = &pkgError.BadRequest{Err: errors.New("email or password invalid")}
		return
	}

	var getTokenRes token.GetToken
	var accessTokenExpiresAt int64
	var expiredAt uint64

	getTokenRes.AccessToken, accessTokenExpiresAt, err = s.token.GetImplicitToken(
		token.Payload{
			ID:       res.ID,
			Email:    res.Email,
			Phone:    res.Phone,
			FullName: res.FullName,
		}, expiredAt)

	getTokenRes.AccessTokenExpiresAt = time.Unix(accessTokenExpiresAt, 0)

	returnData = models.LoginRes{
		AccessToken: getTokenRes.AccessToken,
		ExpiredAt:   getTokenRes.AccessTokenExpiresAt,
	}

	return
}

// Profile ...
func (s *srv) Profile(ctx context.Context, id int64) (returnData models.Profile, err error) {
	res, err := s.repo.User.GetUserByID(ctx, id)
	if err != nil {
		return
	}
	_ = deepcopier.Copy(res).To(&returnData)

	return
}
