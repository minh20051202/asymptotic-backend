package catalog

import (
	"fmt"
	"os"
	"time"

	"github.com/minh20051202/ticket-system-backend/internal/utils"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

type CatalogService interface {
	CreateProvider(req *CreateProviderRequest) error

	CreateEndpoint(req *CreateEndpointRequest) error

	CreatePricing(req *CreatePricingRequest) error
	SetNewPrice(endpointId int64, newCost int64) error

	GetRouteConfigByNames(providerName, endpointName, endpointPath string) (*ReadyRoute, error)
	GetRouteConfigById(endpointId int64) (*RouteConfig, error)
	GetAllAvailableTools() ([]*ToolInfo, error)
}

type service struct {
	repo CatalogRepository
}

func NewService(repo CatalogRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProvider(req *CreateProviderRequest) error {
	if SECRET_KEY == "" {
		return fmt.Errorf("key should be set")
	}

	encryptedApiKey, err := utils.Encrypt(req.ApiKey, SECRET_KEY)
	if err != nil {
		return err
	}

	newProvider := &Provider{
		Name:            req.Name,
		BaseUrl:         req.BaseUrl,
		EncryptedApiKey: encryptedApiKey,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}
	err = s.repo.CreateProvider(newProvider)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateEndpoint(req *CreateEndpointRequest) error {
	if req.ProviderId <= 0 {
		return fmt.Errorf("invalid provider id")
	}

	endpoint := &Endpoint{
		ProviderId: req.ProviderId,
		Name:       req.Name,
		HttpMethod: req.HttpMethod,
		Path:       req.Path,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	err := s.repo.CreateEndpoint(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreatePricing(req *CreatePricingRequest) error {
	if req.EndpointId <= 0 {
		return fmt.Errorf("invalid endpoint id")
	}
	pricing := &Pricing{
		EndpointId: req.EndpointId,
		Cost:       req.Cost,
		CreatedAt:  time.Now().UTC(),
	}
	err := s.repo.CreatePricing(pricing)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) SetNewPrice(endpointId int64, newCost int64) error {
	if newCost < 0 {
		return fmt.Errorf("invalid cost: %d. price cannot be negative", newCost)
	}

	err := s.repo.SetNewPrice(endpointId, newCost)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetRouteConfigById(endpointId int64) (*RouteConfig, error) {
	routeConfig, err := s.repo.GetRouteConfigById(endpointId)

	if err != nil {
		return nil, err
	}

	return routeConfig, nil
}

func (s *service) GetAllAvailableTools() ([]*ToolInfo, error) {
	toolsInfo, err := s.repo.GetAllAvailableTools()

	if err != nil {
		return nil, err
	}

	return toolsInfo, nil
}

func (s *service) GetRouteConfigByNames(providerName, endpointName, endpointPath string) (*ReadyRoute, error) {
	route, err := s.repo.GetRouteConfigByNames(providerName, endpointName, endpointPath)
	if err != nil {
		return nil, err
	}

	apiKey, err := utils.Decrypt(route.EncryptedApiKey, SECRET_KEY)
	if err != nil {
		return nil, err
	}

	readyRoute := &ReadyRoute{
		FullURL: route.ProviderBaseUrl + route.EndpointPath,
		Method:  route.EndpointHttpMethod,
		APIKey:  apiKey,
		Cost:    route.Cost,
	}

	return readyRoute, nil
}
