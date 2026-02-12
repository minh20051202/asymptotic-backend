package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/minh20051202/ticket-system-backend/internal/models"
)

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func createJWT(user *models.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userId":    user.UserID,
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

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
}
