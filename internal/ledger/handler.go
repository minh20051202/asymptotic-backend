package ledger

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/minh20051202/ticket-system-backend/internal/utils"
)

type ledgerHandler struct {
	service LedgerService
}

func NewHandler(service LedgerService) *ledgerHandler {
	return &ledgerHandler{
		service: service,
	}
}

func (h *ledgerHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/transaction", utils.MakeHTTPHandleFunc(h.handleTransaction))
}

func (h *ledgerHandler) handleTransaction(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return h.handleGetTransaction(w, r)
	}
	if r.Method == "POST" {
		return h.handleCreateTransaction(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (h *ledgerHandler) handleGetTransaction(w http.ResponseWriter, r *http.Request) error {
	transactions, err := h.service.GetAllTransactions()

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, transactions)
}

func (h *ledgerHandler) handleCreateTransaction(w http.ResponseWriter, r *http.Request) error {
	createTransactionRequest := new(CreateTransactionRequest)

	if err := json.NewDecoder(r.Body).Decode(createTransactionRequest); err != nil {
		return err
	}

	defer r.Body.Close()

	switch createTransactionRequest.Type {
	case "CHARGE":

		tx, err := h.service.Charge(createTransactionRequest)

		if err != nil {
			if errors.Is(err, ErrInsufficientFunds) {
				return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "insufficient funds"})
			} else if errors.Is(err, ErrAmountNotGreaterThanZero) {
				return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "amount not greater than 0"})
			} else if strings.Contains(err.Error(), "conflict") {
				return utils.WriteJSON(w, http.StatusServiceUnavailable, utils.ApiError{Error: "system busy, please try again"})
			} else {
				return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: err.Error()})
			}
		}
		return utils.WriteJSON(w, http.StatusOK, tx)
	case "DEPOSIT":
		tx, err := h.service.Deposit(createTransactionRequest)

		if err != nil {
			if errors.Is(err, ErrAmountNotGreaterThanZero) {
				return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "amount not greater than 0"})
			} else if strings.Contains(err.Error(), "conflict") {
				return utils.WriteJSON(w, http.StatusServiceUnavailable, utils.ApiError{Error: "system busy, please try again"})
			} else {
				return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: err.Error()})
			}
		}
		return utils.WriteJSON(w, http.StatusOK, tx)
	default:
		return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "invalid transaction type, must be CHARGE or DEPOSIT"})
	}
}
