package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CloudCIX/gocloudcix/config"
	"github.com/CloudCIX/gocloudcix/option"
)

// GetToken attempts to generate a CloudCIX token for the given credentials
// If apiURL is empty, it defaults to "https://api.cloudcix.com/"
func GetToken(email, password, apiKey string, opts ...option.RequestOption) (string, error) {
	return GetTokenWithURL(email, password, apiKey, "", opts...)
}

// GetTokenWithURL attempts to generate a CloudCIX token for the given credentials with a custom API URL
// If apiURL is empty, it defaults to "https://api.cloudcix.com/"
func GetTokenWithURL(email, password, apiKey, apiURL string, opts ...option.RequestOption) (string, error) {
	if apiURL == "" {
		apiURL = "https://api.cloudcix.com/"
	}

	settings := &config.Settings{
		CLOUDCIX_API_V2_URL:   apiURL,
		CLOUDCIX_API_USERNAME: email,
		CLOUDCIX_API_PASSWORD: password,
		CLOUDCIX_API_KEY:      apiKey,
	}

	tokenManager := NewTokenManager(settings)
	return tokenManager.GetValidToken(context.Background(), opts...)
}

// GetAdminToken generates an admin token using environment variables or settings file
func GetAdminToken(settingsFile ...string) (string, error) {
	settings, err := config.LoadSettings(settingsFile...)
	if err != nil {
		return "", fmt.Errorf("failed to load settings: %w", err)
	}

	if err := settings.Validate(); err != nil {
		return "", fmt.Errorf("invalid settings: %w", err)
	}

	tokenManager := NewTokenManager(settings)
	return tokenManager.GetValidToken(context.Background())
}

// WithTokenAuth creates a request option that sets the X-Auth-Token header
func WithTokenAuth(token string) option.RequestOption {
	return option.WithHeader("X-Auth-Token", token)
}

// WithAutoAuth creates a request option that automatically adds authentication
func WithAutoAuth(tokenManager *TokenManager) option.RequestOption {
	return option.WithMiddleware(func(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
		// Get valid token
		token, err := tokenManager.GetValidToken(req.Context())
		if err != nil {
			return nil, err
		}

		// Add auth header
		req.Header.Set("X-Auth-Token", token)

		return next(req)
	})
}
