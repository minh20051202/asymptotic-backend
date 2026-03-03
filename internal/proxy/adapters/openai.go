package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/minh20051202/ticket-system-backend/internal/catalog"
	"github.com/minh20051202/ticket-system-backend/internal/proxy"
)

type OpenAIAdapter struct{}

type openAIResp struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		TotalTokens int64 `json:"total_tokens"`
	} `json:"usage"`
}

func (a *OpenAIAdapter) PrepareRequest(route *catalog.ReadyRoute, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(route.Method, route.FullURL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+route.APIKey)
	return req, nil
}

func (a *OpenAIAdapter) ParseResponse(r *http.Response) (*proxy.Response, error) {
	if r.StatusCode >= 400 {
		return nil, fmt.Errorf("openai error: %d", r.StatusCode)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var res openAIResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &proxy.Response{
		Text:     res.Choices[0].Message.Content,
		Usage:    res.Usage.TotalTokens,
		Provider: "openai",
		Payload:  body,
	}, nil
}
