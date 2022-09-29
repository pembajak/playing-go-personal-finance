package rest

import (
	"github.com/pembajak/personal-finance/internal/app/delivery/rest/account" // for rest
	"github.com/pembajak/personal-finance/internal/app/delivery/rest/finance" // for rest
	"github.com/pembajak/personal-finance/internal/app/delivery/rest/report"  // for rest
	"github.com/pembajak/personal-finance/internal/app/delivery/rest/user"    // for rest
	"github.com/pembajak/personal-finance/internal/app/usecase"
)

type Delivery interface {
	UserDelivery() user.UserDeliverer
	AccountDelivery() account.AccountDeliverer
	FinanceDelivery() finance.FinanceDeliverer
	ReportDelivery() report.ReportDeliverer
}

type delivery struct {
	user    user.UserDeliverer
	account account.AccountDeliverer
	finance finance.FinanceDeliverer
	report  report.ReportDeliverer
}

func NewRestDelivery(usecase *usecase.UseCase) Delivery {
	h := new(delivery)
	h.user = user.NewDelivery(usecase)
	h.account = account.NewDelivery(usecase)
	h.finance = finance.NewDelivery(usecase)
	h.report = report.NewDelivery(usecase)

	return h
}

func (delivery *delivery) UserDelivery() user.UserDeliverer {
	return delivery.user
}

func (delivery *delivery) AccountDelivery() account.AccountDeliverer {
	return delivery.account
}

func (delivery *delivery) FinanceDelivery() finance.FinanceDeliverer {
	return delivery.finance
}

func (delivery *delivery) ReportDelivery() report.ReportDeliverer {
	return delivery.report
}
