package server

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/minh20051202/ticket-system-backend/internal/auth"
	"github.com/minh20051202/ticket-system-backend/internal/crypto"
	db "github.com/minh20051202/ticket-system-backend/internal/database"
	"github.com/minh20051202/ticket-system-backend/internal/shared"
)

const PREFIX string = "asym_sk_"

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	storage    db.Storage
}

func NewAPIServer(listenAddr string, storage db.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/login", makeHTTPHandleFunc(s.handleLogin))
	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/user/{uuid}", withJWTAuth(makeHTTPHandleFunc(s.handleUserById)))
	router.HandleFunc("/transaction", makeHTTPHandleFunc(s.handleTransaction))
	router.HandleFunc("/api-keys", withJWTAuth(makeHTTPHandleFunc(s.handleCreateApiKey)))
	log.Println("Server is running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUser(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleUserById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUserById(w, r)
	}

	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	users, err := s.storage.GetAllUsers()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserReq := new(CreateUserRequest)

	if err := json.NewDecoder(r.Body).Decode(createUserReq); err != nil {
		return err
	}

	defer r.Body.Close()

	hashedPassword, err := auth.HashPassword(createUserReq.Password)

	if err != nil {
		return err
	}

	newUser := &shared.User{
		UserId:    uuid.New(),
		Username:  createUserReq.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.storage.CreateUserWithBalance(newUser); err != nil {
		return err
	}

	jwt, err := createJWT(newUser)

	if err != nil {
		return err
	}

	fmt.Println("JWT Token:", jwt)

	return WriteJSON(w, http.StatusOK, jwt)
}

func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	uuidVar, err := getUUID(r)

	if err != nil {
		return err
	}

	user, err := s.storage.GetUserById(uuidVar)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleTransaction(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetTransaction(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateTransaction(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetTransaction(w http.ResponseWriter, r *http.Request) error {
	transactions, err := s.storage.GetAllTransactions()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, transactions)
}

func (s *APIServer) handleCreateTransaction(w http.ResponseWriter, r *http.Request) error {
	createTransactionRequest := new(CreateTransactionRequest)

	if err := json.NewDecoder(r.Body).Decode(createTransactionRequest); err != nil {
		return err
	}

	defer r.Body.Close()

	switch createTransactionRequest.Type {
	case "CHARGE":
		newTransaction := &shared.Transaction{
			TransactionId:  uuid.New(),
			UserId:         createTransactionRequest.UserId,
			IdempotencyKey: createTransactionRequest.IdempotencyKey,
			Amount:         createTransactionRequest.Amount,
			Type:           "CHARGE",
			CreatedAt:      time.Now().UTC(),
		}
		tx, err := s.storage.Charge(newTransaction)

		if err != nil {
			if errors.Is(err, db.ErrInsufficientFunds) {
				return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "insufficient funds"})
			} else if errors.Is(err, db.ErrAmountNotGreaterThanZero) {
				return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "amount not greater than 0"})
			} else if strings.Contains(err.Error(), "conflict") {
				return WriteJSON(w, http.StatusServiceUnavailable, ApiError{Error: "system busy, please try again"})
			} else {
				return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
			}
		}
		return WriteJSON(w, http.StatusOK, tx)
	case "DEPOSIT":
		newTransaction := &shared.Transaction{
			TransactionId:  uuid.New(),
			UserId:         createTransactionRequest.UserId,
			IdempotencyKey: createTransactionRequest.IdempotencyKey,
			Amount:         createTransactionRequest.Amount,
			Type:           "DEPOSIT",
			CreatedAt:      time.Now().UTC(),
		}
		tx, err := s.storage.Deposit(newTransaction)

		if err != nil {
			if errors.Is(err, db.ErrAmountNotGreaterThanZero) {
				return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "amount not greater than 0"})
			} else if strings.Contains(err.Error(), "conflict") {
				return WriteJSON(w, http.StatusServiceUnavailable, ApiError{Error: "system busy, please try again"})
			} else {
				return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
			}
		}
		return WriteJSON(w, http.StatusOK, tx)
	default:
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "invalid transaction type, must be CHARGE or DEPOSIT"})
	}
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	loginRequest := new(LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
		return err
	}

	defer r.Body.Close()

	user, err := s.storage.GetUserByUsername(loginRequest.Username)

	if err != nil {
		return fmt.Errorf("Invalid username or password")
	}

	if err := auth.CheckPasswordHash(loginRequest.Password, user.Password); err != nil {
		return err
	}

	jwt, err := createJWT(user)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, jwt)
}

func (s *APIServer) handleCreateApiKey(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value("userId")

	parsedUserId := userId.(uuid.UUID)

	apiKeyReq := new(CreateApiKeyRequest)

	if err := json.NewDecoder(r.Body).Decode(apiKeyReq); err != nil {
		return err
	}

	defer r.Body.Close()

	key, err := crypto.GenerateSecureToken(32)

	key = fmt.Sprintf("%v%v", PREFIX, key)

	hashedKey := sha256.Sum256([]byte(key))

	apiKey := &shared.ApiKey{
		ApiKey: hex.EncodeToString(hashedKey[:]),
		UserId: parsedUserId,
		Name:   apiKeyReq.Name,
	}

	err = s.storage.CreateApiKey(apiKey)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, CreateApiKeyResponse{ApiKey: key})
}

func getUUID(r *http.Request) (uuid.UUID, error) {
	uuidStr := mux.Vars(r)["uuid"]
	uuid, err := uuid.Parse(uuidStr)

	if err != nil {
		return uuid, fmt.Errorf("Invalid uuid given %s", uuid)
	}

	return uuid, nil
}
