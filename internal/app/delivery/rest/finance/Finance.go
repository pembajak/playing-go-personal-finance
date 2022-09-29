package finance

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/app/usecase"
)

type FinanceDeliverer interface {
	CreateFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
	DeleteFinanceByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
	UpdateFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
	GetFinanceByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
}

type delivery struct {
	usecase *usecase.UseCase
}

func NewDelivery(usecase *usecase.UseCase) FinanceDeliverer {
	return &delivery{
		usecase: usecase,
	}
}

// CreateFinance ...
func (d *delivery) CreateFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request models.Finance
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (models.Finance{}) {
		err = errors.New("error payload")
		return
	}

	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)
	request.UserID = int64(userID)
	returnData, err = d.usecase.Finance.CreateFinance(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// UpdateFinance ...
func (d *delivery) UpdateFinance(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request models.Finance
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (models.Finance{}) {
		err = errors.New("error payload")
		return
	}

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)
	request.UserID = int64(userID)
	request.ID = id

	returnData, err = d.usecase.Finance.UpdateFinance(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// GetFinanceByID ...
func (d *delivery) GetFinanceByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	returnData, err = d.usecase.Finance.GetFinanceByID(ctx, id)

	w.Header().Add("Content-Type", "application/json")

	return
}

// DeleteFinanceByID ...
func (d *delivery) DeleteFinanceByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	returnData, err = d.usecase.Finance.DeleteFinanceByID(ctx, id)

	w.Header().Add("Content-Type", "application/json")

	return
}
