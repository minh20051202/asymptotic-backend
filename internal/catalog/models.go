package catalog

import (
	"time"
)

type Provider struct {
	ProviderId      int64     `json:"providerId"`
	Name            string    `json:"name"`
	BaseUrl         string    `json:"baseUrl"`
	EncryptedApiKey string    `json:"-"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type Endpoint struct {
	EndpointId int64     `json:"endpointId"`
	ProviderId int64     `json:"providerId"`
	Name       string    `json:"name"`
	HttpMethod string    `json:"httpMethod"`
	Path       string    `json:"path"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Pricing struct {
	PricingId  int64      `json:"pricingId"`
	EndpointId int64      `json:"endpointId"`
	Cost       int64      `json:"cost"`
	IsCurrent  bool       `json:"isCurrent"`
	CreatedAt  time.Time  `json:"createdAt"`
	ValidUntil *time.Time `json:"validUntil"`
}

type RouteConfig struct {
	ProviderBaseUrl    string
	EncryptedApiKey    string
	EndpointPath       string
	EndpointHttpMethod string
	Cost               int64
}

type ToolInfo struct {
	EndpointId   int64
	ProviderName string
	EndpointName string
	HttpMethod   string
	Path         string
	Cost         int64
}

type CreateProviderRequest struct {
	Name    string `json:"name"`
	BaseUrl string `json:"baseUrl"`
	ApiKey  string `json:"apiKey"`
}

type GetProviderResponse struct {
	ProviderId      int64     `json:"providerId"`
	Name            string    `json:"name"`
	BaseUrl         string    `json:"baseUrl"`
	EncryptedApiKey string    `json:"-"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type CreateEndpointRequest struct {
	ProviderId int64     `json:"providerId"`
	Name       string    `json:"name"`
	HttpMethod string    `json:"httpMethod"`
	Path       string    `json:"path"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type CreatePricingRequest struct {
	EndpointId int64     `json:"endpointId"`
	Cost       int64     `json:"cost"`
	IsCurrent  bool      `json:"isCurrent"`
	CreatedAt  time.Time `json:"createdAt"`
}
