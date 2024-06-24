// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/phoebe-go/internal/apiquery"
	"github.com/stainless-sdks/phoebe-go/internal/param"
	"github.com/stainless-sdks/phoebe-go/internal/requestconfig"
	"github.com/stainless-sdks/phoebe-go/option"
)

// ProductTop100Service contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductTop100Service] method instead.
type ProductTop100Service struct {
	Options []option.RequestOption
}

// NewProductTop100Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProductTop100Service(opts ...option.RequestOption) (r *ProductTop100Service) {
	r = &ProductTop100Service{}
	r.Options = opts
	return
}

// Get the top 100 contributors on a given date for a country or region.
//
// #### Notes
//
// The results are updated every 15 minutes.
//
// When ordering by the number of completed checklists, the number of species seen
// will always be zero. Similarly when ordering by the number of species seen the
// number of completed checklists will always be zero. <b>Selected Response Field
// Notes</b>
//
// profileHandle - if a user has enabled their profile, this is the handle to reach
// it via ebird.org/ebird/profile/{profileHandle}
//
// numSpecies - always zero when checklistSort parameter is true. Invalid
// observations ARE included in this total numCompleteChecklists - always zero when
// checklistSort parameter is false
func (r *ProductTop100Service) List(ctx context.Context, regionCode string, y int64, m int64, d int64, query ProductTop100ListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if regionCode == "" {
		err = errors.New("missing required regionCode parameter")
		return
	}
	path := fmt.Sprintf("product/top100/%s/%v/%v/%v", regionCode, y, m, d)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type ProductTop100ListParams struct {
	// Only fetch this number of contributors.
	MaxResults param.Field[int64] `query:"maxResults"`
	// Order by number of complete checklists (cl) or by number of species seen (spp).
	RankedBy param.Field[ProductTop100ListParamsRankedBy] `query:"rankedBy"`
}

// URLQuery serializes [ProductTop100ListParams]'s query parameters as
// `url.Values`.
func (r ProductTop100ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Order by number of complete checklists (cl) or by number of species seen (spp).
type ProductTop100ListParamsRankedBy string

const (
	ProductTop100ListParamsRankedBySpp ProductTop100ListParamsRankedBy = "spp"
	ProductTop100ListParamsRankedByCl  ProductTop100ListParamsRankedBy = "cl"
)

func (r ProductTop100ListParamsRankedBy) IsKnown() bool {
	switch r {
	case ProductTop100ListParamsRankedBySpp, ProductTop100ListParamsRankedByCl:
		return true
	}
	return false
}
