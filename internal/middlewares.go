package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"wishlist/logger"

	"github.com/rs/xid"
	"github.com/rs/zerolog"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const (
	UserIDCtxKey    = contextKey("userID")
	UserEmailCtxKey = contextKey("userEmail")
	RequestIDCtxKey = contextKey("request_id")
)

type AuthMiddleware struct {
	jwks *keyfunc.JWKS
}

func CreateAuthMiddleware(jwksURL string) (*AuthMiddleware, error) {
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval:   time.Hour,
		RefreshTimeout:    10 * time.Second,
		RefreshUnknownKID: true,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create JWKS from URL %s: %w", jwksURL, err)
	}

	return &AuthMiddleware{jwks: jwks}, nil
}

func (am *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		l := zerolog.Ctx(r.Context())

		if authHeader == "" {
			writeJSONError(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			writeJSONError(w, "Authorization header must be in 'Bearer <token>' format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, am.jwks.Keyfunc)
		if err != nil {
			l.Err(err).Msg("Failed to validate token.")
			writeJSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			l.Warn().Msg("Invalid token")
			writeJSONError(w, "Token is invalid", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			l.Err(err).Msg("Failed to parse token claims")
			writeJSONError(w, "Could not parse token claims", http.StatusInternalServerError)
			return
		}

		headerUserID, err := getClaim(claims, "user_id")
		if err != nil {
			l.Err(err).Msg("Failed to get user_id from token claims")
			writeJSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		email, err := getClaim(claims, "email")
		if err != nil {
			l.Err(err).Msg("Failed to get email from token claims")
			writeJSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDCtxKey, headerUserID)
		ctx = context.WithValue(ctx, UserEmailCtxKey, email)

		l.Debug().Str("UserID", headerUserID).Str("Email", email).Msg("Authenticated request")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func getClaim(claims jwt.MapClaims, key string) (string, error) {
	value, ok := claims[key]
	if !ok {
		return "", fmt.Errorf("missing required claim: %s", key)
	}
	strValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("claim '%s' is not a string", key)
	}
	return strValue, nil
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = xid.New().String()
		}

		ctx := context.WithValue(r.Context(), RequestIDCtxKey, requestID)

		// Create request-scoped logger with request_id
		reqLogger := logger.Get().With().
			Str(string(RequestIDCtxKey), requestID).
			Logger()

		ctx = reqLogger.WithContext(ctx)
		r = r.WithContext(ctx)

		w.Header().Set("X-Request-ID", requestID)

		lrw := newLoggingResponseWriter(w)

		defer func() {
			panicVal := recover()
			if panicVal != nil {
				lrw.statusCode = http.StatusInternalServerError
				panic(panicVal)
			}

			reqLogger.
				Info().
				Str("method", r.Method).
				Str("url", r.URL.RequestURI()).
				Str("user_agent", r.UserAgent()).
				Str("remote_addr", r.RemoteAddr).
				Int("status_code", lrw.statusCode).
				Dur("elapsed_ms", time.Since(start)).
				Msg("incoming request")
		}()

		next.ServeHTTP(lrw, r)
	})
}
