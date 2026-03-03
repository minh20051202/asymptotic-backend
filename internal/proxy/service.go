package proxy

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/minh20051202/ticket-system-backend/internal/catalog"
)

type ReadyRoute struct {
	FullURL string
	Method  string
	APIKey  string
	Cost    int64
}

type CatalogProvider interface {
	GetRouteConfigByNames(providerName, endpointName, endpointPath string) (*catalog.ReadyRoute, error)
}

type LedgerProvider interface {
	Charge(userId uuid.UUID, amount int64, idempotencyKey string) error
}

type ProxyService interface {
	Run(req *ProxyRequest, userId uuid.UUID) (*Response, error)
}

type service struct {
	catalog  CatalogProvider
	ledger   LedgerProvider
	client   *http.Client
	adapters map[string]ProviderAdapter
}

func NewService(catalog CatalogProvider, ledger LedgerProvider) *service {
	return &service{
		catalog: catalog,
		ledger:  ledger,
		client: &http.Client{
			Timeout: time.Second * 30,
		},
		adapters: make(map[string]ProviderAdapter),
	}
}

func (s *service) Run(req *ProxyRequest, userId uuid.UUID) (*Response, error) {
	route, err := s.catalog.GetRouteConfigByNames(req.ProviderName, req.EndpointName, req.EndpointPath)
	if err != nil {
		return nil, err
	}

	idempotencyKey := uuid.New().String()
	err = s.ledger.Charge(userId, route.Cost, idempotencyKey)
	if err != nil {
		return nil, err
	}

	adapter, ok := s.adapters[req.ProviderName]
	if !ok {
		return nil, fmt.Errorf("unsupported provider: %s", req.ProviderName)
	}

	httpReq, err := adapter.PrepareRequest(route, req.Payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return adapter.ParseResponse(resp)

}

func (s *service) RegisterAdapter(name string, a ProviderAdapter) {
	s.adapters[name] = a
}
