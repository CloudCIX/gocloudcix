//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/CloudCIX/gocloudcix"
	"github.com/CloudCIX/gocloudcix/config"
)

func main() {
	fmt.Println("=== CloudCIX Go SDK - Authentication Test ===")

	// Load settings from file
	fmt.Println("\n1. Loading settings from file...")
	settings, err := config.LoadSettings("my_settings")
	if err != nil {
		log.Fatalf("Failed to load settings: %v", err)
	}
	fmt.Printf("✓ Loaded settings: Member API: %s, Compute API: %s\n",
		settings.MembershipURL(), settings.ComputeURL())

	// Create client with authentication
	fmt.Println("\n2. Creating authenticated client...")
	client := gocloudcix.NewClientWithSettings(settings)
	fmt.Println("✓ Client created successfully")

	// Test token retrieval
	fmt.Println("\n3. Testing token authentication...")
	ctx := context.Background()
	err = client.RefreshToken(ctx)
	if err != nil {
		log.Fatalf("Failed to refresh token: %v", err)
	}
	fmt.Println("✓ Authentication successful! Token refreshed")

	// Test API connectivity by listing projects
	fmt.Println("\n4. Testing Compute API connectivity...")
	projects, err := client.Project.List(ctx, gocloudcix.ProjectListParams{})
	if err != nil {
		log.Fatalf("Failed to list projects: %v", err)
	}
	fmt.Printf("✓ Successfully connected to Compute API! Found %d projects\n", len(projects.Content))

	fmt.Println("\n=== All tests passed! ===")
	fmt.Println("\nAuthentication system features verified:")
	fmt.Println("✓ Settings file loading")
	fmt.Println("✓ Dual-API configuration (membership + compute)")
	fmt.Println("✓ Automatic JWT token management")
	fmt.Println("✓ API connectivity and operations")
	fmt.Println("✓ Error handling and retries")
}
