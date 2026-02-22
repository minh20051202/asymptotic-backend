package server

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/minh20051202/ticket-system-backend/internal/shared"
)

type contextKey string

const userContextKey contextKey = "userId"

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func createJWT(user *shared.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userId":    user.UserId,
		"username":  user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecretKey))
}

func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		tokenString := parts[1]
		token, err := validateJWT(tokenString)

		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}

		if !token.Valid {
			WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			WriteJSON(w, http.StatusUnauthorized, ApiError{Error: "invalid token claims"})
			return
		}

		claim, err := uuid.Parse(claims["userId"].(string))

		if err != nil {
			WriteJSON(w, http.StatusUnauthorized, ApiError{Error: "invalid token claims"})
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claim)

		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
}
