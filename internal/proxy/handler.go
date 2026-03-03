package proxy

import (
	"encoding/json"
	"net/http"

	"github.com/minh20051202/ticket-system-backend/internal/identity"
	"github.com/minh20051202/ticket-system-backend/internal/utils"
)

type proxyHandler struct {
	service ProxyService
}

func NewHandler(service ProxyService) *proxyHandler {
	return &proxyHandler{
		service: service,
	}
}

func (h *proxyHandler) HandleRun(w http.ResponseWriter, r *http.Request) error {
	req := new(ProxyRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "invalid request"})
	}
	userId, ok := identity.GetUserIdFromContext(r.Context())
	if !ok {
		return utils.WriteJSON(w, http.StatusUnauthorized, utils.ApiError{Error: "unauthorized"})
	}
	response, err := h.service.Run(req, userId)
	if err != nil {
		return utils.WriteJSON(w, http.StatusPaymentRequired, utils.ApiError{Error: err.Error()})

	}

	return utils.WriteJSON(w, http.StatusOK, response)
}
