package middleware

import (
	"context"
	"github.com/quniverse26/miniproject/utils"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			utils.Response(w, 401, "unauthorized", nil)
			return
		}

		user, err := utils.ValidateToken(accessToken)
		if err != nil {
			utils.Response(w, 401, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userinfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}