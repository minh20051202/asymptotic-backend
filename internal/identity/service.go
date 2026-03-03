package identity

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const userContextKey contextKey = "userId"

const API_KEY_PREFIX string = "asym_sk_"

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

type IdentityService interface {
	CreateUser(req *CreateUserRequest) (string, error)
	Login(req *LoginRequest) (string, error)
	GetAllUsers() ([]*User, error)
	GetUserById(uuid uuid.UUID) (*User, error)

	CreateApiKey(userId uuid.UUID, req *CreateApiKeyRequest) (string, error)
	GetUserIdByApiKeyHash(apiKeyHash string) (uuid.UUID, error)
}

type service struct {
	repo IdentityRepository
}

func NewService(repo IdentityRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(req *CreateUserRequest) (string, error) {
	hashedPassword, err := hashPassword(req.Password)

	if err != nil {
		return "", err
	}

	newUser := &User{
		UserId:    uuid.New(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.CreateUserWithBalance(newUser); err != nil {
		return "", err
	}

	jwt, err := createJWT(newUser)

	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *service) Login(req *LoginRequest) (string, error) {

	user, err := s.repo.GetUserByUsername(req.Username)

	if err != nil {
		return "", err
	}

	if err := checkPasswordHash(user.Password, req.Password); err != nil {
		return "", err
	}

	jwt, err := createJWT(user)

	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (s *service) GetAllUsers() ([]*User, error) {
	users, err := s.repo.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) GetUserById(uuid uuid.UUID) (*User, error) {
	user, err := s.repo.GetUserById(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) CreateApiKey(userId uuid.UUID, req *CreateApiKeyRequest) (string, error) {
	key, err := generateSecureToken(32)

	key = fmt.Sprintf("%v%v", API_KEY_PREFIX, key)

	hashedKey := sha256.Sum256([]byte(key))

	apiKey := &ApiKey{
		ApiKey: hex.EncodeToString(hashedKey[:]),
		UserId: userId,
		Name:   req.Name,
	}

	err = s.repo.CreateApiKey(apiKey)

	if err != nil {
		return "", err
	}
	return apiKey.ApiKey, err
}

func (s *service) GetUserIdByApiKeyHash(apiKeyHash string) (uuid.UUID, error) {
	return s.repo.GetUserIdByApiKeyHash(apiKeyHash)
}

func hashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return fmt.Errorf("invalid password: %w\n", err)
	}
	return nil
}

func generateSecureToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func createJWT(user *User) (string, error) {
	claims := jwt.MapClaims{
		"userId":   user.UserId.String(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecretKey))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
}
