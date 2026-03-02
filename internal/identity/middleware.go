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
		if authHeader == "" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "missing authorization header"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid authorization format"})
			return
		}
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
		if authHeader == "" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "missing authorization header"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid authorization format"})
			return
		}
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

func WithAdminAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "missing authorization header"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid authorization format"})
			return
		}
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

		if role, ok := claims["role"].(string); !ok || role != "ADMIN" {
			utils.WriteJSON(w, http.StatusForbidden, utils.ApiError{Error: "permission denied"})
			return
		}

		userIdStr, idOk := claims["userId"].(string)
		if !idOk {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid token claims"})
			return
		}

		userId, err := uuid.Parse(userIdStr)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid token claims"})
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, userId)

		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}
