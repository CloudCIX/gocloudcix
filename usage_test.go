// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocloudcix_test

import (
	"context"
	"os"
	"testing"

	"github.com/CloudCIX/gocloudcix"
	"github.com/CloudCIX/gocloudcix/internal/testutil"
	"github.com/CloudCIX/gocloudcix/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gocloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Prism tests are disabled")
	backups, err := client.Compute.Backups.List(context.TODO(), gocloudcix.ComputeBackupListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", backups.Metadata)
}
