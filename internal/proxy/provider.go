package proxy

import (
	"net/http"

	"github.com/minh20051202/ticket-system-backend/internal/catalog"
)

type ProviderAdapter interface {
	PrepareRequest(route *catalog.ReadyRoute, payload []byte) (*http.Request, error)

	ParseResponse(resp *http.Response) (*Response, error)
}
