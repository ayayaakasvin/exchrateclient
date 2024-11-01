package exchrateclient

import (
	"fmt"
	"time"
)

// fetcher is an interface defining the methods for fetching exchange rate data.
type fetcher interface {
	FetchCodes() (map[string]iso4217, error) // Fetches supported currency codes.
	FetchPair(first string, second string) (*pair, error) // Fetches the exchange rate between two currencies.
	FetchIndex(code string) (*index, error) // Fetches the index for a specific base currency.
	Fetch(endpoint, ...string) (*responseFromServer, error) // Generic method for making API calls.
}

// responseFromServer defines the structure of the API response.
// It includes common fields that are returned in every request.
type responseFromServer struct {
	Result             string    `json:"result"` // General result message from the API.
	LastUpdateString   string    `json:"time_last_update_utc,omitempty"` // Last update time as a string (UTC).
	LastUpdateDateTime time.Time `json:"-"` // Last update time as a time.Time object, not exported.
	NextUpdateString   string    `json:"time_next_update_utc,omitempty"` // Next scheduled update time as a string (UTC).
	NextUpdateDateTime time.Time `json:"-"` // Next update time as a time.Time object, not exported.

	BaseCode string `json:"base_code,omitempty"` // Base currency code, omitted if empty.

	// ConversionRateMap is a map of conversion rates for different currencies.
	// For example, `{"USD": 1, "AUD": 1.4817}`.
	ConversionRateMap map[string]float64 `json:"conversion_rates,omitempty"`

	// Fields specific to the pair endpoint.
	TargetCode          string  `json:"target_code,omitempty"` // Target currency code.
	ConversionRateFloat float64 `json:"conversion_rate,omitempty"` // Conversion rate from base to target currency.

	// SupportedCodes contains a list of supported currency codes and their names.
	// Format: [["AED", "UAE Dirham"], ...].
	SupportedCodes [][]string `json:"supported_codes,omitempty"`
}

// endpoint is a custom type for defining API endpoints.
type endpoint string

// Defining constants for the various API endpoints.
const (
	IndexEndpoint endpoint = "index" // Endpoint for fetching index data.
	PairEndpoint  endpoint = "pair"  // Endpoint for fetching pair exchange rates.
	CodesEndpoint endpoint = "codes"  // Endpoint for fetching supported currency codes.
)

// pair represents the exchange rate between two currencies.
type pair struct {
	BaseCode   string  `json:"base"`   // Base currency code.
	TargetCode string  `json:"target"` // Target currency code.
	Rate       float64 `json:"rate"`   // Exchange rate from base to target.
}

// index represents the index for a specific base currency.
type index struct {
	BaseCode string             `json:"base"` // Base currency code.
	RateMap  map[string]float64 `json:"index_of_code"` // Map of currency codes to their rates relative to the base currency.
}

// iso4217 represents the structure of a currency code according to ISO 4217 standard.
type iso4217 struct {
	Name string `json:"name"` // Name of the currency.
	Code string `json:"code"` // ISO 4217 code of the currency.
}

// String method for pair struct, providing a string representation of the exchange rate.
func (p *pair) String() string {
	if p == nil {
		return "empty pair"
	}
	return fmt.Sprintf("%s rate to %s: %.4f", p.BaseCode, p.TargetCode, p.Rate)
}

// String method for index struct, providing a string representation of the currency rates.
func (i *index) String() string {
	if i == nil {
		return "empty index"
	}

	var result string = fmt.Sprintf("%s has rates of:\n", i.BaseCode)

	for key, value := range i.RateMap {
		result += fmt.Sprintf("\t%s : %.4f\n", key, value)
	}

	return result
}

// String method for iso4217 struct, providing a string representation of the currency.
func (iso *iso4217) String() string {
	if iso == nil {
		return "empty ISO4217"
	}
	return fmt.Sprintf("%s -> %s", iso.Code, iso.Name)
}
