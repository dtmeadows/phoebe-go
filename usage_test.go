// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/phoebe-go"
	"github.com/stainless-sdks/phoebe-go/internal/testutil"
	"github.com/stainless-sdks/phoebe-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := phoebe.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	refHotspotInfoGetResponse, err := client.Ref.Hotspot.Info.Get(context.TODO(), "L99381")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", refHotspotInfoGetResponse.CountryCode)
}
