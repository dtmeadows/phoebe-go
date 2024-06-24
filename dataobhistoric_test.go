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

func TestDataObHistoricListWithOptionalParams(t *testing.T) {
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
	err := client.Data.Obs.Historic.List(
		context.TODO(),
		"string",
		int64(0),
		int64(1),
		int64(1),
		phoebe.DataObHistoricListParams{
			Cat:                phoebe.F(phoebe.DataObHistoricListParamsCatSpecies),
			Detail:             phoebe.F(phoebe.DataObHistoricListParamsDetailSimple),
			Hotspot:            phoebe.F(true),
			IncludeProvisional: phoebe.F(true),
			MaxResults:         phoebe.F(int64(1)),
			R:                  phoebe.F([]string{"string"}),
			Rank:               phoebe.F(phoebe.DataObHistoricListParamsRankMrec),
			SppLocale:          phoebe.F("string"),
		},
	)
	if err != nil {
		var apierr *phoebe.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
