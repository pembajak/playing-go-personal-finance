package finance

import (
	"context"
	"math"

	"github.com/go-playground/validator/v10"
	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/app/repository"
	"github.com/ulule/deepcopier"
)

type srv struct {
	repo      *repository.Repositories
	validator *validator.Validate
}

func NewFinanceUsecase(repo *repository.Repositories) FinanceUseCase {
	return &srv{
		repo:      repo,
		validator: validator.New(),
	}
}

// CreateFinance ...
func (s *srv) CreateFinance(ctx context.Context, param models.Finance) (returnData models.Finance, err error) {

	financeParam := models.FinanceReq{}
	_ = deepcopier.Copy(param).To(&financeParam)

	err = s.validator.Struct(financeParam)
	if err != nil {
		return
	}

	financeRepo := models.Finance{}
	_ = deepcopier.Copy(param).To(&financeRepo)

	retData, err := s.repo.Finance.CreateFinance(ctx, financeRepo)
	if err != nil {
		return
	}

	finance, err := s.repo.Finance.GetFinanceByID(ctx, retData.ID)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(finance).To(&returnData)
	return
}

// UpdateFinance ...
func (s *srv) UpdateFinance(ctx context.Context, param models.Finance) (returnData models.Finance, err error) {

	financeParam := models.FinanceReq{}
	_ = deepcopier.Copy(param).To(&financeParam)

	err = s.validator.Struct(financeParam)
	if err != nil {
		return
	}

	financeRepo := models.Finance{}
	_ = deepcopier.Copy(param).To(&financeRepo)

	_, err = s.repo.Finance.UpdateFinance(ctx, financeRepo)
	if err != nil {
		return
	}

	finance, err := s.repo.Finance.GetFinanceByID(ctx, param.ID)
	if err != nil {
		return
	}

	_ = deepcopier.Copy(finance).To(&returnData)
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

// GetAllFinance ...
func (s *srv) GetAllFinance(ctx context.Context, req models.GetAllFinanceReq) (returnData models.GetAllFinanceRes, err error) {
	var financeRepo []models.Finance
	var financeReq models.GetAllFinanceReq
	_ = deepcopier.Copy(req).To(&financeReq)
	res, err := s.repo.Finance.GetAllFinance(ctx, financeReq)
	if err != nil {
		return
	}

	for _, v := range res {
		var financeTmp models.Finance
		_ = deepcopier.Copy(v).To(&financeTmp)
		financeTmp.Account = models.Account(v.Account)
		financeRepo = append(financeRepo, financeTmp)
	}

	totalData, err := s.repo.Finance.CountTotalFinance(ctx, financeReq)
	if err != nil {
		return
	}

	if len(financeRepo) == 0 {
		financeRepo = []models.Finance{}
	}

	returnData.Data = financeRepo
	returnData.Total = totalData
	returnData.Page = 1
	returnData.Limit = 10

	if req != (models.GetAllFinanceReq{}) {
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
