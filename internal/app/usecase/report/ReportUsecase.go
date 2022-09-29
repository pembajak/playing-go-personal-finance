package report

import (
	"context"

	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/app/repository"
	"github.com/ulule/deepcopier"
)

type srv struct {
	repo *repository.Repositories
}

func NewReportUsecase(repo *repository.Repositories) ReportUseCase {
	return &srv{
		repo: repo,
	}
}

// GetTotalTransactionDaily ...
func (s *srv) GetTotalTransactionDaily(ctx context.Context, req models.GetTotalTransaction) (returnData interface{}, err error) {
	financeRepo := models.GetTotalTransaction{}
	_ = deepcopier.Copy(req).To(&financeRepo)

	res, err := s.repo.Report.GetTotalTransactionDaily(ctx, financeRepo)

	returnData = res
	return
}

// GetTotalTransactionMonthly ...
func (s *srv) GetTotalTransactionMonthly(ctx context.Context, req models.GetTotalTransaction) (returnData interface{}, err error) {
	financeRepo := models.GetTotalTransaction{}
	_ = deepcopier.Copy(req).To(&financeRepo)

	res, err := s.repo.Report.GetTotalTransactionMonthly(ctx, financeRepo)

	returnData = res

	return
}
