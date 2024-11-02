package api

import (
	"fmt"
	"net/http"
	"strings"
)

func adminOnly(accountService AccountsService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := extractAuthToken(r)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			err = accountService.IsAdmin(r.Context(), token)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func extractAuthToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("unauthorized")
	}

	// Split the header into scheme and token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("unauthorized")
	}

	return parts[1], nil
}
