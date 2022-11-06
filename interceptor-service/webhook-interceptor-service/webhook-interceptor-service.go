package webhookis

import (
	is "github.com/devingen/sepet-api/interceptor-service"
	"github.com/go-resty/resty/v2"
	"strings"
)

// WebhookInterceptorService implements ISepetInterceptorService interface with web integration
type WebhookInterceptorService struct {
	Address string
	Client  *resty.Client
}

// New generates new WebhookInterceptorService
// address: The complete URL of the webhook api.
// headersValue: Key=value list of headers joined by ",". E.g.("X-Api-Key=abc,X-Client=web-hook")
func New(address string, headersValue string) is.ISepetInterceptorService {

	if address == "" {
		return WebhookInterceptorService{}
	}

	httpClient := resty.New().
		SetBaseURL(address).
		SetHeader("Content-Type", "application/json")

	if headersValue != "" {
		for _, keyAndValue := range strings.Split(headersValue, ",") {
			headerParts := strings.SplitN(keyAndValue, "=", 2)
			httpClient.SetHeader(headerParts[0], headerParts[1])
		}
	}

	return WebhookInterceptorService{
		Address: address,
		Client:  httpClient,
	}
}
