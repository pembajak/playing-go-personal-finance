package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pembajak/personal-finance/internal/app/models"
	"github.com/pembajak/personal-finance/internal/app/usecase"
)

// Handler ...
type UserDeliverer interface {
	// CreateUser ...
	CreateUser(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	//Login ...
	Login(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)

	// Profile ...
	Profile(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error)
}

type delivery struct {
	usecase *usecase.UseCase
}

// NewHandler ...
func NewDelivery(usecase *usecase.UseCase) UserDeliverer {
	return &delivery{
		usecase: usecase,
	}
}

func (h *delivery) CreateUser(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request models.User
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (models.User{}) {
		err = errors.New("error payload")
		return
	}

	returnData, err = h.usecase.User.CreateUser(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

func (h *delivery) Login(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request models.User
	err = json.Unmarshal(bodyByte, &request)

	if err != nil || request == (models.User{}) {
		err = errors.New("error payload")
		return
	}

	returnData, err = h.usecase.User.Login(ctx, request)
	w.Header().Add("Content-Type", "application/json")

	return
}

func (h *delivery) Profile(w http.ResponseWriter, r *http.Request) (returnData interface{}, err error) {
	ctx := r.Context()
	userID := r.Context().Value("claims").(jwt.MapClaims)["ID"].(float64)

	returnData, err = h.usecase.User.Profile(ctx, int64(userID))
	w.Header().Add("Content-Type", "application/json")
	return
}
