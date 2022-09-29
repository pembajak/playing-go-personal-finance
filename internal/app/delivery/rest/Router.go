package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pembajak/personal-finance/internal/app/middleware"
	httppkg "github.com/pembajak/personal-finance/internal/pkg/http"
)

func (s *server) Router(delivery Delivery) (w httppkg.Router) {
	w = httppkg.NewRouter(chi.NewRouter())

	w.Route("/v1", func(r chi.Router) {
		router := r.(httppkg.Router)
		router.Action(httppkg.NewRest(http.MethodPost, "/login", delivery.UserDelivery().Login))
		router.Action(httppkg.NewRest(http.MethodPost, "/register", delivery.UserDelivery().CreateUser))

		router.Route("/user", func(r chi.Router) {
			router := r.(httppkg.Router)
			router.Use(middleware.JWTAuthorization)
			router.Action(httppkg.NewRest(http.MethodGet, "/profile", delivery.UserDelivery().Profile))
		})

		router.Route("/account", func(r chi.Router) {
			router := r.(httppkg.Router)
			router.Use(middleware.JWTAuthorization)
			router.Action(httppkg.NewRest(http.MethodPost, "/", delivery.AccountDelivery().CreateAccount))
			router.Action(httppkg.NewRest(http.MethodPatch, "/{id}", delivery.AccountDelivery().UpdateAccount))
			router.Action(httppkg.NewRest(http.MethodGet, "/", delivery.AccountDelivery().GetAllAccount))
			router.Action(httppkg.NewRest(http.MethodGet, "/{id}", delivery.AccountDelivery().GetByAccountByID))
			router.Action(httppkg.NewRest(http.MethodDelete, "/{id}", delivery.AccountDelivery().DeleteAccount))
		})

		router.Route("/finance", func(r chi.Router) {
			router := r.(httppkg.Router)
			router.Use(middleware.JWTAuthorization)
			router.Action(httppkg.NewRest(http.MethodPost, "/", delivery.FinanceDelivery().CreateFinance))
			router.Action(httppkg.NewRest(http.MethodPatch, "/{id}", delivery.FinanceDelivery().UpdateFinance))
			router.Action(httppkg.NewRest(http.MethodGet, "/{id}", delivery.FinanceDelivery().GetFinanceByID))
			router.Action(httppkg.NewRest(http.MethodDelete, "/{id}", delivery.FinanceDelivery().DeleteFinanceByID))
		})

		router.Route("/report", func(r chi.Router) {
			router := r.(httppkg.Router)
			router.Use(middleware.JWTAuthorization)
			router.Action(httppkg.NewRest(http.MethodGet, "/daily", delivery.ReportDelivery().GetTotalTransactionDaily))
			router.Action(httppkg.NewRest(http.MethodGet, "/monthly", delivery.ReportDelivery().GetTotalTransactionMonthly))
		})

	})
	return
}
