package server

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	db "github.com/minh20051202/ticket-system-backend/internal/database"
	"github.com/minh20051202/ticket-system-backend/internal/models"
)

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
	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/user/{uuid}", withJWTAuth(makeHTTPHandleFunc(s.handleUserById)))
	router.HandleFunc("/wallet", makeHTTPHandleFunc(s.handleWallet))
	router.HandleFunc("/wallet/{uuid}", makeHTTPHandleFunc(s.handleWalletById))
	router.HandleFunc("/transaction", makeHTTPHandleFunc(s.handleTransaction))
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

	if r.Method == "DELETE" {
		return s.handleDeleteUserById(w, r)
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

	newUser := &models.User{
		UserID:    uuid.New(),
		Username:  createUserReq.Username,
		Password:  createUserReq.Password,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.storage.CreateUser(newUser); err != nil {
		return err
	}

	jwt, err := createJWT(newUser)

	if err != nil {
		return err
	}

	fmt.Println("JWT Token:", jwt)

	return WriteJSON(w, http.StatusOK, newUser.UserID)
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

func (s *APIServer) handleDeleteUserById(w http.ResponseWriter, r *http.Request) error {
	uuidVar, err := getUUID(r)

	if err != nil {
		return err
	}

	if err := s.storage.DeleteUserById(uuidVar); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]uuid.UUID{"deleted": uuidVar})
}

func (s *APIServer) handleWallet(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetWallet(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateWallet(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
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

func (s *APIServer) handleWalletById(w http.ResponseWriter, r *http.Request) error {
	uuidVar, err := getUUID(r)

	wallet, err := s.storage.GetWalletById(uuidVar)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, wallet)
}

func (s *APIServer) handleGetWallet(w http.ResponseWriter, r *http.Request) error {
	wallets, err := s.storage.GetAllWallets()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, wallets)
}

func (s *APIServer) handleCreateWallet(w http.ResponseWriter, r *http.Request) error {
	createWalletReq := new(CreateWalletRequest)

	if err := json.NewDecoder(r.Body).Decode(createWalletReq); err != nil {
		return err
	}

	defer r.Body.Close()

	newWallet := &models.Wallet{
		WalletID:     uuid.New(),
		WalletName:   createWalletReq.WalletName,
		TotalQuota:   createWalletReq.TotalQuota,
		AvailableQty: createWalletReq.AvailableQty,
		CreatedAt:    time.Now().UTC(),
	}

	err := s.storage.CreateWallet(newWallet)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newWallet)
}

func (s *APIServer) handleDeleteWallet(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetTransaction(w http.ResponseWriter, r *http.Request) error {
	transactions, err := s.storage.GetAllTransactions()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, transactions)
}

func (s *APIServer) handleCreateTransaction(w http.ResponseWriter, r *http.Request) error {
	var err error
	maxRetries := 20
	createTransactionRequest := new(CreateTransactionRequest)

	if err := json.NewDecoder(r.Body).Decode(createTransactionRequest); err != nil {
		return err
	}

	defer r.Body.Close()

	newTransaction := &models.Transaction{
		TransactionID: uuid.New(),
		WalletID:      createTransactionRequest.WalletID,
		UserID:        createTransactionRequest.UserID,
		Amount:        createTransactionRequest.Amount,
		CreatedAt:     time.Now().UTC(),
	}
	for range maxRetries {
		err = s.storage.CreateTransactionPessimistic(newTransaction)
		if err == nil {
			return WriteJSON(w, http.StatusOK, newTransaction)
		}
		if !strings.Contains(err.Error(), "conflict") {
			break
		}
		sleepTime := time.Duration(rand.Intn(20)+5) * time.Millisecond
		time.Sleep(sleepTime)
	}

	if err != nil {
		if strings.Contains(err.Error(), "sold out") {
			return WriteJSON(w, http.StatusConflict, ApiError{Error: "sold out"})
		} else if strings.Contains(err.Error(), "conflict") {
			return WriteJSON(w, http.StatusServiceUnavailable, ApiError{Error: "system busy, please try again"})
		}
	}
	return WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
}

func getUUID(r *http.Request) (uuid.UUID, error) {
	uuidStr := mux.Vars(r)["uuid"]
	uuid, err := uuid.Parse(uuidStr)

	if err != nil {
		return uuid, fmt.Errorf("Invalid uuid given %s", uuid)
	}

	return uuid, nil
}
