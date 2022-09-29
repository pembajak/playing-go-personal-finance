package account

import (
	"context"
	"math"

	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/app/repository"
	"github.com/pembajak/personal-finance/internal/pkg/token"
	"github.com/ulule/deepcopier"
)

// srv ...
type srv struct {
	repo  *repository.Repositories
	token token.IToken
}

// NewSrv ..
func NewAccountCase(repo *repository.Repositories, token token.IToken) AccountUseCase {
	return &srv{
		repo:  repo,
		token: token,
	}
}

// CreateAccount ...
func (s srv) CreateAccount(ctx context.Context, param models.Account) (returnData models.Account, err error) {
	accountRepo := models.Account{}
	_ = deepcopier.Copy(param).To(&accountRepo)

	res, err := s.repo.Account.CreateAccount(ctx, accountRepo)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)

	return
}

// UpdateAccount ...
func (s srv) UpdateAccount(ctx context.Context, param models.Account) (returnData models.Account, err error) {
	accountRepo := models.Account{}
	_ = deepcopier.Copy(param).To(&accountRepo)
	res, err := s.repo.Account.UpdateAccount(ctx, accountRepo)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)

	return
}

// GetAccountByID ...
func (s srv) GetAccountByID(ctx context.Context, id int64) (returnData models.Account, err error) {
	res, err := s.repo.Account.GetAccountByID(ctx, id)
	if err != nil {
		return
	}
	_ = deepcopier.Copy(res).To(&returnData)
	return
}

// GetAllAccount ...
func (s srv) GetAllAccount(ctx context.Context, req models.GetAllAccountReq) (returnData models.GetAllAccountRes, err error) {
	var accountRepo []models.Account
	var accountReq models.GetAllAccountReq
	_ = deepcopier.Copy(req).To(&accountReq)
	res, err := s.repo.Account.GetAllAccount(ctx, accountReq)
	if err != nil {
		return
	}

	for _, v := range res {
		var accountTmp models.Account
		_ = deepcopier.Copy(v).To(&accountTmp)
		accountRepo = append(accountRepo, accountTmp)
	}

	totalData, err := s.repo.Account.CountTotalAccount(ctx, accountReq.UserID)
	if err != nil {
		return
	}

	if len(accountRepo) == 0 {
		accountRepo = []models.Account{}
	}

	returnData.Data = accountRepo
	returnData.Total = totalData
	returnData.Page = 1
	returnData.Limit = 10

	if req != (models.GetAllAccountReq{}) {
		if req.Limit != 0 {
			returnData.Limit = req.Limit
		}

		if req.Page != 0 {
			returnData.Page = req.Page
		}

		returnData.LastPage = int(math.Ceil(float64(returnData.Total) / float64(returnData.Limit)))
	}

	return
}

// DeleteAccount ...
func (s srv) DeleteAccount(ctx context.Context, id int64) (returnData models.Account, err error) {
	res, err := s.repo.Account.GetAccountByID(ctx, id)
	if err != nil {
		return
	}
	_, err = s.repo.Account.DeleteAccount(ctx, id)
	if err != nil {
		return
	}
	_ = deepcopier.Copy(res).To(&returnData)
	return
}
