package identity

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/minh20051202/ticket-system-backend/internal/utils"
)

func withApiKeyAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		tokenString := parts[1]
		token, err := validateJWT(tokenString)

		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, utils.ApiError{Error: "permission denied"})
			return
		}

		if !token.Valid {
			utils.WriteJSON(w, http.StatusForbidden, utils.ApiError{Error: "permission denied"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid token claims"})
			return
		}

		claim, err := uuid.Parse(claims["userId"].(string))

		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid token claims"})
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claim)

		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		tokenString := parts[1]
		token, err := validateJWT(tokenString)

		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, utils.ApiError{Error: "permission denied"})
			return
		}

		if !token.Valid {
			utils.WriteJSON(w, http.StatusForbidden, utils.ApiError{Error: "permission denied"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid token claims"})
			return
		}

		claim, err := uuid.Parse(claims["userId"].(string))

		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid token claims"})
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claim)

		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}
