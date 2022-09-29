package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pembajak/personal-finance/config"
	httppkg "github.com/pembajak/personal-finance/internal/pkg/http"
	"github.com/pembajak/personal-finance/internal/pkg/token"
)

func JWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		cfg := config.NewConfig()
		if !strings.Contains(authorizationHeader, "Bearer") {
			httppkg.WriteError(w, r, jwt.NewValidationError("Invalid Authorization Header", jwt.ValidationErrorUnverifiable), http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		claims, err := token.New(token.WithSecretKey(cfg.GetString("app.secret_key"))).GetClaims(tokenStr)
		if err != nil {
			httppkg.WriteError(w, r, err, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
