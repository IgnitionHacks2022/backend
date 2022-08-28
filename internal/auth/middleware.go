package auth

import (
	"context"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/statistics" {
			log.Println("Auth route!")
			token := r.Header.Get("token")

			id, err := ValidateToken(token)
			if err != nil || id == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userID", id)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
