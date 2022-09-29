package report

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/app/usecase"
)

type ReportDeliverer interface {
	GetTotalTransactionDaily(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
	GetTotalTransactionMonthly(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
}

type delivery struct {
	usecase *usecase.UseCase
}

func NewDelivery(usecase *usecase.UseCase) ReportDeliverer {
	return &delivery{
		usecase: usecase,
	}
}

// GetTotalTransactionDaily ...
func (d *delivery) GetTotalTransactionDaily(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)
	request := models.GetTotalTransaction{
		UserID: int64(userID),
		Type:   r.URL.Query().Get("type"),
	}

	returnData, err = d.usecase.Report.GetTotalTransactionDaily(ctx, request)

	w.Header().Add("Content-Type", "application/json")
	return
}

func (d *delivery) GetTotalTransactionMonthly(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)
	request := models.GetTotalTransaction{
		UserID: int64(userID),
		Type:   r.URL.Query().Get("type"),
	}

	returnData, err = d.usecase.Report.GetTotalTransactionMonthly(ctx, request)

	w.Header().Add("Content-Type", "application/json")
	return
}
