package identity

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/minh20051202/ticket-system-backend/internal/shared"
	"github.com/minh20051202/ticket-system-backend/internal/utils"
)

type identityHandler struct {
	service IdentityService
}

func NewHandler(service IdentityService) *identityHandler {
	return &identityHandler{
		service: service,
	}
}

func (h *identityHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", utils.MakeHTTPHandleFunc(h.handleLogin))
	router.HandleFunc("/user", utils.MakeHTTPHandleFunc(h.handleUser))
	router.HandleFunc("/user/{uuid}", withJWTAuth(utils.MakeHTTPHandleFunc(h.handleUserById)))
	router.HandleFunc("/api-keys", withJWTAuth(utils.MakeHTTPHandleFunc(h.handleCreateApiKey)))
}

func (h *identityHandler) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return h.handleGetUser(w, r)
	}
	if r.Method == "POST" {
		return h.handleCreateUser(w, r)
	}
	return fmt.Errorf("method not allowed: %h", r.Method)
}

func (h *identityHandler) handleUserById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return h.handleGetUserById(w, r)
	}

	return fmt.Errorf("method not allowed: %h", r.Method)
}

func (h *identityHandler) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	users, err := h.service.GetAllUsers()

	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: "internal server error"})
	}

	return utils.WriteJSON(w, http.StatusOK, users)
}

func (h *identityHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserReq := new(shared.CreateUserRequest)

	if err := json.NewDecoder(r.Body).Decode(createUserReq); err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "invalid credentials"})
	}

	defer r.Body.Close()

	jwt, err := h.service.CreateUser(createUserReq)

	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: "Internal server error. Please try again later!"})
	}

	return utils.WriteJSON(w, http.StatusOK, shared.CreateUserResponse{Username: createUserReq.Username, Email: createUserReq.Email, JWT: jwt})
}

func (h *identityHandler) handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	uuidVar, err := getUUID(r)

	if err != nil {
		return err
	}

	user, err := h.service.GetUserById(uuidVar)

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, user)
}

func (h *identityHandler) handleCreateApiKey(w http.ResponseWriter, r *http.Request) error {
	userId, ok := getUserIdFromContext(r.Context())

	if !ok {
		return utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid credentials"})
	}

	apiKeyReq := new(shared.CreateApiKeyRequest)

	if err := json.NewDecoder(r.Body).Decode(apiKeyReq); err != nil {
		return err
	}

	defer r.Body.Close()

	apiKey, err := h.service.CreateApiKey(userId, apiKeyReq)

	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: "Internal server error. Please try again later!"})
	}

	return utils.WriteJSON(w, http.StatusOK, shared.CreateApiKeyResponse{ApiKey: apiKey})
}

func getUUID(r *http.Request) (uuid.UUID, error) {
	uuidStr := mux.Vars(r)["uuid"]
	uuid, err := uuid.Parse(uuidStr)

	if err != nil {
		return uuid, fmt.Errorf("Invalid uuid given %h", uuid)
	}

	return uuid, nil
}

func (h *identityHandler) handleLogin(w http.ResponseWriter, r *http.Request) error {
	loginRequest := new(shared.LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(loginRequest); err != nil {
		return err
	}

	defer r.Body.Close()

	jwt, err := h.service.Login(loginRequest)

	if err != nil {
		return utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "invalid credentials"})
	}

	return utils.WriteJSON(w, http.StatusOK, jwt)
}

func (h *identityHandler) handleRun(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *identityHandler) handleSearch(w http.ResponseWriter, r *http.Request) error {
	return nil
}
