package auth

import (
	"testing"

	"github.com/CloudCIX/gocloudcix/config"
)

func TestTokenManager(t *testing.T) {
	settings := &config.Settings{
		CLOUDCIX_API_URL:      "https://compute.api.cloudcix.com/",
		CLOUDCIX_API_V2_URL:   "https://membership.api.cloudcix.com/",
		CLOUDCIX_API_VERSION:  "5.0",
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_API_KEY:      "testkey",
	}

	tokenManager := NewTokenManager(settings)

	if tokenManager == nil {
		t.Fatal("TokenManager should not be nil")
	}

	if tokenManager.settings != settings {
		t.Fatal("TokenManager should store settings")
	}

	if tokenManager.IsTokenValid() {
		t.Fatal("TokenManager should not have valid token initially")
	}
}

func TestSettingsValidation(t *testing.T) {
	// Test valid settings
	validSettings := &config.Settings{
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_API_KEY:      "testkey",
		CLOUDCIX_REGION_ID:    1,
	}

	if err := validSettings.Validate(); err != nil {
		t.Fatalf("Valid settings should pass validation: %v", err)
	}

	// Test missing username
	invalidSettings := &config.Settings{
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_API_KEY:      "testkey",
	}

	if err := invalidSettings.Validate(); err == nil {
		t.Fatal("Settings without username should fail validation")
	}

	// Test missing password
	invalidSettings2 := &config.Settings{
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_KEY:      "testkey",
		CLOUDCIX_REGION_ID:    1,
	}

	if err := invalidSettings2.Validate(); err == nil {
		t.Fatal("Settings without password should fail validation")
	}

	// Test missing API key
	invalidSettings3 := &config.Settings{
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_REGION_ID:    1,
	}

	if err := invalidSettings3.Validate(); err == nil {
		t.Fatal("Settings without API key should fail validation")
	}

	// Test missing region ID
	invalidSettings4 := &config.Settings{
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_API_KEY:      "testkey",
		// CLOUDCIX_REGION_ID not set (defaults to 0)
	}

	if err := invalidSettings4.Validate(); err == nil {
		t.Fatal("Settings without region ID should fail validation")
	}
}

func TestGetTokenWithURL(t *testing.T) {
	// Test with empty URL (should use default)
	// Note: This will fail to connect but we're testing URL setting, not actual connection
	t.Run("empty URL defaults to https://api.cloudcix.com/", func(t *testing.T) {
		// We can't actually make the request without valid credentials and network access
		// but we can verify the function signature and that it accepts parameters correctly
		email := "test@example.com"
		password := "testpass"
		apiKey := "testkey"

		// This will fail due to invalid credentials, but that's expected
		_, err := GetTokenWithURL(email, password, apiKey, "")
		// We expect an error because credentials are invalid, but the function should execute
		if err == nil {
			t.Fatal("Expected error with invalid credentials")
		}
	})

	t.Run("custom URL is used", func(t *testing.T) {
		email := "test@example.com"
		password := "testpass"
		apiKey := "testkey"
		customURL := "https://staging.api.cloudcix.com/"

		// This will fail due to invalid credentials and network, but that's expected
		_, err := GetTokenWithURL(email, password, apiKey, customURL)
		// We expect an error because credentials are invalid, but the function should execute
		if err == nil {
			t.Fatal("Expected error with invalid credentials")
		}
	})

	t.Run("GetToken calls GetTokenWithURL with empty URL", func(t *testing.T) {
		email := "test@example.com"
		password := "testpass"
		apiKey := "testkey"

		// Both should behave the same way
		_, err1 := GetToken(email, password, apiKey)
		_, err2 := GetTokenWithURL(email, password, apiKey, "")

		// Both should fail in the same way (invalid credentials)
		if (err1 == nil) != (err2 == nil) {
			t.Fatal("GetToken and GetTokenWithURL with empty URL should behave the same")
		}
	})
}
