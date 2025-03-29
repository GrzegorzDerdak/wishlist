package internal

import (
	"fmt"
	"net/http"
)

// Middleware to check for a specific header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the presence of the "Authorization" header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// If the header is missing, return a 401 Unauthorized response
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Unauthorized: Missing Authorization header")
			return
		}

		// Validate the header (e.g., check if it matches a specific value)
		if authHeader != "Bearer my-secret-token" {
			// If the header is invalid, return a 403 Forbidden response
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Forbidden: Invalid Authorization header")
			return
		}

		// If the header is valid, call the next handler
		next.ServeHTTP(w, r)
	})
}
