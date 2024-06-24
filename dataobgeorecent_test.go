// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/phoebe-go"
	"github.com/stainless-sdks/phoebe-go/internal/testutil"
	"github.com/stainless-sdks/phoebe-go/option"
)

func TestDataObGeoRecentListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := phoebe.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Data.Obs.Geo.Recent.List(context.TODO(), phoebe.DataObGeoRecentListParams{
		Lat:                phoebe.F(-90.000000),
		Lng:                phoebe.F(-180.000000),
		Back:               phoebe.F(int64(1)),
		Cat:                phoebe.F(phoebe.DataObGeoRecentListParamsCatSpecies),
		Dist:               phoebe.F(int64(0)),
		Hotspot:            phoebe.F(true),
		IncludeProvisional: phoebe.F(true),
		MaxResults:         phoebe.F(int64(1)),
		Sort:               phoebe.F(phoebe.DataObGeoRecentListParamsSortDate),
		SppLocale:          phoebe.F("string"),
	})
	if err != nil {
		var apierr *phoebe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
