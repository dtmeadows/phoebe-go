// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package phoebe

import (
	"context"
	"fmt"
	"net/http"

	"github.com/phoebe-bird/phoebe-go/internal/apijson"
	"github.com/phoebe-bird/phoebe-go/internal/param"
	"github.com/phoebe-bird/phoebe-go/internal/requestconfig"
	"github.com/phoebe-bird/phoebe-go/option"
)

// RefTaxonomyLocaleService contains methods and other services that help with
// interacting with the phoebe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefTaxonomyLocaleService] method instead.
type RefTaxonomyLocaleService struct {
	Options []option.RequestOption
}

// NewRefTaxonomyLocaleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRefTaxonomyLocaleService(opts ...option.RequestOption) (r *RefTaxonomyLocaleService) {
	r = &RefTaxonomyLocaleService{}
	r.Options = opts
	return
}

// Returns the list of supported locale codes and names for species common names,
// with the last time they were updated. Use the accept-language header to get
// translated language names when available.
//
// NOTE: The locale codes and names are stable but the other fields in this result
// are not yet finalized and should be used with caution.
func (r *RefTaxonomyLocaleService) List(ctx context.Context, query RefTaxonomyLocaleListParams, opts ...option.RequestOption) (res *[]RefTaxonomyLocaleListResponse, err error) {
	if query.AcceptLanguage.Present {
		opts = append(opts, option.WithHeader("Accept-Language", fmt.Sprintf("%s", query.AcceptLanguage)))
	}
	opts = append(r.Options[:], opts...)
	path := "ref/taxa-locales/ebird"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RefTaxonomyLocaleListResponse struct {
	Code        string                            `json:"code"`
	LastUpdated string                            `json:"lastUpdated"`
	Name        string                            `json:"name"`
	JSON        refTaxonomyLocaleListResponseJSON `json:"-"`
}

// refTaxonomyLocaleListResponseJSON contains the JSON metadata for the struct
// [RefTaxonomyLocaleListResponse]
type refTaxonomyLocaleListResponseJSON struct {
	Code        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefTaxonomyLocaleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refTaxonomyLocaleListResponseJSON) RawJSON() string {
	return r.raw
}

type RefTaxonomyLocaleListParams struct {
	AcceptLanguage param.Field[string] `header:"Accept-Language"`
}
