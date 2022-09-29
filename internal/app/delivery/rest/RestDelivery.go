package rest

import (
	"github.com/pembajak/personal-finance/internal/app/delivery/rest/user" // for rest
	"github.com/pembajak/personal-finance/internal/app/usecase"
)

type Delivery interface {
	UserDelivery() user.UserDeliverer
}

type delivery struct {
	user user.UserDeliverer
}

func NewRestDelivery(usecase *usecase.UseCase) Delivery {
	h := new(delivery)
	h.user = user.NewDelivery(usecase)

	return h
}

func (delivery *delivery) UserDelivery() user.UserDeliverer {
	return delivery.user
}
