package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	db "github.com/minh20051202/ticket-system-backend/internal/database"
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
	router.HandleFunc("/user/{uuid}", makeHTTPHandleFunc(s.handleGetUserById))
	router.HandleFunc("/event", makeHTTPHandleFunc(s.handleEvent))
	router.HandleFunc("/event/{uuid}", makeHTTPHandleFunc(s.handleGetEventById))
	router.HandleFunc("/order", makeHTTPHandleFunc(s.handleOrder))
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
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) error {
	uuidStr := mux.Vars(r)["uuid"]
	uuid, err := uuid.Parse(uuidStr)

	if err != nil {
		return fmt.Errorf("Invalid uuid given %s", uuidStr)
	}

	user, err := s.storage.GetUserById(uuid)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
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

	newUser := &db.User{
		UserID:    uuid.New(),
		Username:  createUserReq.Username,
		Password:  createUserReq.Password,
		CreatedAt: time.Now().UTC(),
	}

	err := s.storage.CreateUser(newUser)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newUser.UserID)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleEvent(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetEvent(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateEvent(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteEvent(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleOrder(w http.ResponseWriter, r *http.Request) error {
	// if r.Method == "GET" {
	// 	return s.handleGetOrder(w, r)
	// }
	if r.Method == "POST" {
		return s.handleCreateOrder(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetEventById(w http.ResponseWriter, r *http.Request) error {
	uuidStr := mux.Vars(r)["uuid"]
	uuid, err := uuid.Parse(uuidStr)

	if err != nil {
		return fmt.Errorf("Invalid uuid given %s", uuidStr)
	}

	event, err := s.storage.GetEventById(uuid)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, event)
}

func (s *APIServer) handleGetEvent(w http.ResponseWriter, r *http.Request) error {
	events, err := s.storage.GetAllEvents()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, events)
}

func (s *APIServer) handleCreateEvent(w http.ResponseWriter, r *http.Request) error {
	createEventReq := new(CreateEventRequest)

	if err := json.NewDecoder(r.Body).Decode(createEventReq); err != nil {
		return err
	}

	newEvent := &db.Event{
		EventID:      uuid.New(),
		EventName:    createEventReq.EventName,
		TotalQuota:   createEventReq.TotalQuota,
		AvailableQty: createEventReq.AvailableQty,
		CreatedAt:    time.Now().UTC(),
	}

	err := s.storage.CreateEvent(newEvent)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newEvent)
}

func (s *APIServer) handleDeleteEvent(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// func (s *APIServer) handleGetOrder(w http.ResponseWriter, r *http.Request) error {
// 	orders, err := s.storage.GetAllOrders()

// 	if err != nil {
// 		return err
// 	}

// 	return WriteJSON(w, http.StatusOK, orders)
// }

func (s *APIServer) handleCreateOrder(w http.ResponseWriter, r *http.Request) error {
	createOrderRequest := new(CreateOrderRequest)

	if err := json.NewDecoder(r.Body).Decode(createOrderRequest); err != nil {
		return err
	}

	newOrder := &db.Order{
		OrderID:   uuid.New(),
		EventID:   createOrderRequest.EventID,
		UserID:    createOrderRequest.UserID,
		Amount:    createOrderRequest.Amount,
		CreatedAt: time.Now().UTC(),
	}

	err := s.storage.CreateOrder(newOrder)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newOrder)
}
