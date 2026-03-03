package proxy

import (
	"encoding/json"
)

type ProxyRequest struct {
	ProviderName string          `json:"providerName"`
	EndpointName string          `json:"endpointName"`
	EndpointPath string          `json:"endpointPath"`
	Payload      json.RawMessage `json:"payload"`
}

type Response struct {
	Text     string          `json:"text"`
	Usage    int64           `json:"usage"`
	Provider string          `json:"provider"`
	Payload  json.RawMessage `json:"payload"`
}
