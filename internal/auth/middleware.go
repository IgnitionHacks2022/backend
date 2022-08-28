package auth

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/statistics" {
			log.Println("Auth route!")
		}
		next.ServeHTTP(w, r)
	})
}
