package finance

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/app/repository"
	"github.com/ulule/deepcopier"
)

type srv struct {
	repo *repository.Repositories
}

func NewFinanceUsecase(repo *repository.Repositories) FinanceUseCase {
	return &srv{
		repo: repo,
	}
}

// CreateFinance ...
func (s *srv) CreateFinance(ctx context.Context, param models.Finance) (returnData models.Finance, err error) {
	financeRepo := models.Finance{}
	_ = deepcopier.Copy(param).To(&financeRepo)

	res, err := s.repo.Finance.CreateFinance(ctx, financeRepo)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)
	return
}

// UpdateFinance ...
func (s *srv) UpdateFinance(ctx context.Context, param models.Finance) (returnData models.Finance, err error) {
	financeRepo := models.Finance{}
	_ = deepcopier.Copy(param).To(&financeRepo)

	res, err := s.repo.Finance.UpdateFinance(ctx, financeRepo)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)
	return
}

// DeleteFinanceByID ...
func (s *srv) DeleteFinanceByID(ctx context.Context, id int64) (returnData models.Finance, err error) {
	res, err := s.repo.Finance.GetFinanceByID(ctx, id)
	if err != nil {
		return
	}

	err = s.repo.Finance.DeleteFinanceByID(ctx, id)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(res).To(&returnData)
	return
}

// GetFinanceByID ...
func (s *srv) GetFinanceByID(ctx context.Context, id int64) (returnData models.Finance, err error) {
	res, err := s.repo.Finance.GetFinanceByID(ctx, id)
	if err != nil {
		return
	}

	returnData = models.Finance{
		ID:        res.ID,
		Title:     res.Title,
		AccountID: res.AccountID,
		Account: models.Account{
			ID:          res.Account.ID,
			Name:        res.Account.Name,
			Type:        res.Account.Type,
			Description: res.Account.Description,
			UserID:      res.Account.UserID,
			CreatedAt:   res.Account.CreatedAt,
			UpdatedAt:   res.Account.UpdatedAt,
		},
		Amount:          res.Amount,
		Description:     res.Description,
		UserID:          res.UserID,
		Type:            res.Type,
		TransactionDate: res.TransactionDate,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       res.UpdatedAt,
	}

	return
}
