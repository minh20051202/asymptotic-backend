package catalog

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minh20051202/ticket-system-backend/internal/identity"
	"github.com/minh20051202/ticket-system-backend/internal/utils"
)

type catalogHandler struct {
	service CatalogService
}

func NewHandler(service CatalogService) *catalogHandler {
	return &catalogHandler{
		service: service,
	}
}

func (h *catalogHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/provider", identity.WithAdminAuth(utils.MakeHTTPHandleFunc(h.handleCreateProvider))).Methods(http.MethodPost)
	router.HandleFunc("/endpoint", identity.WithAdminAuth(utils.MakeHTTPHandleFunc(h.handleCreateEndpoint))).Methods(http.MethodPost)
	router.HandleFunc("/pricing", identity.WithAdminAuth(utils.MakeHTTPHandleFunc(h.handleCreatePricing))).Methods(http.MethodPost)
	router.HandleFunc("/tools", utils.MakeHTTPHandleFunc(h.handleGetTools)).Methods(http.MethodGet)
}

func (h *catalogHandler) handleCreateProvider(w http.ResponseWriter, r *http.Request) error {
	createProviderReq := new(CreateProviderRequest)

	if err := json.NewDecoder(r.Body).Decode(createProviderReq); err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "invalid credentials"})
	}

	defer r.Body.Close()

	err := h.service.CreateProvider(createProviderReq)

	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: "Internal server error. Please try again later!"})
	}

	return utils.WriteJSON(w, http.StatusOK, CreateProviderResponse{Name: createProviderReq.Name, BaseUrl: createProviderReq.BaseUrl})
}

func (h *catalogHandler) handleCreateEndpoint(w http.ResponseWriter, r *http.Request) error {
	createEndpointReq := new(CreateEndpointRequest)

	if err := json.NewDecoder(r.Body).Decode(createEndpointReq); err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "invalid credentials"})
	}

	defer r.Body.Close()

	err := h.service.CreateEndpoint(createEndpointReq)

	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: "Internal server error. Please try again later!"})
	}

	return utils.WriteJSON(w, http.StatusOK, CreateEndpointRequest{ProviderId: createEndpointReq.ProviderId, Name: createEndpointReq.Name, HttpMethod: createEndpointReq.HttpMethod, Path: createEndpointReq.Path})
}

func (h *catalogHandler) handleCreatePricing(w http.ResponseWriter, r *http.Request) error {
	createPricingReq := new(CreatePricingRequest)

	if err := json.NewDecoder(r.Body).Decode(createPricingReq); err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: "invalid credentials"})
	}

	defer r.Body.Close()

	err := h.service.CreatePricing(createPricingReq)

	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: "Internal server error. Please try again later!"})
	}

	return utils.WriteJSON(w, http.StatusOK, CreatePricingResponse{EndpointId: createPricingReq.EndpointId, Cost: createPricingReq.Cost})
}

func (h *catalogHandler) handleGetTools(w http.ResponseWriter, r *http.Request) error {
	tools, err := h.service.GetAllAvailableTools()
	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, utils.ApiError{Error: "Failed to fetch available tools"})
	}

	return utils.WriteJSON(w, http.StatusOK, tools)
}
