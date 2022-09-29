package account

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

type AccountDeliverer interface {
	// CreateAccount ...
	CreateAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// UpdateAccount ...
	UpdateAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// DeleteAccount ...
	DeleteAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// GetAllAccount ...
	GetAllAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// GetByAccountByID ...
	GetByAccountByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
}

type delivery struct {
	usecase *usecase.UseCase
}

func NewDelivery(usecase *usecase.UseCase) AccountDeliverer {
	return &delivery{
		usecase: usecase,
	}
}

// CreateAccount ...
func (h *delivery) CreateAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request models.Account
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (models.Account{}) {
		err = errors.New("error payload")
		return
	}

	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)
	request.UserID = int64(userID)
	returnData, err = h.usecase.Account.CreateAccount(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// UpdateAccount ...
func (h *delivery) UpdateAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request models.Account
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (models.Account{}) {
		err = errors.New("error payload")
		return
	}

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	request.ID = id
	returnData, err = h.usecase.Account.UpdateAccount(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

// DeleteAccount ...
func (h *delivery) DeleteAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	returnData, err = h.usecase.Account.DeleteAccount(ctx, id)
	w.Header().Add("Content-Type", "application/json")
	return
}

// GetAllAccount ...
func (h *delivery) GetAllAccount(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)

	request := models.GetAllAccountReq{
		Page:   page,
		Limit:  limit,
		UserID: int64(userID),
	}

	returnData, err = h.usecase.Account.GetAllAccount(ctx, request)
	w.Header().Add("Content-Type", "application/json")
	return
}

// GetByAccountByID ...
func (h *delivery) GetByAccountByID(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	returnData, err = h.usecase.Account.GetAccountByID(ctx, id)
	w.Header().Add("Content-Type", "application/json")

	return
}
