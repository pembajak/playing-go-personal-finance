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

	})
	return
}
