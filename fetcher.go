package exchrateclient

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ayayaakasvin/exchrateclient/lib/errorhand" // Custom error handling package
)

// Constants for API endpoint keys and success status.
const (
	latestUrlKey = "latest" // Key for the latest exchange rates endpoint.
	codesUrlKey  = "codes"  // Key for the supported currency codes endpoint.
	pairUrlKey   = "pair"   // Key for fetching exchange rate between two currencies.
	successKey   = "success" // Key indicating successful API response.
)

// apiFetcher is responsible for fetching data from the exchange rate API.
type apiFetcher struct {
	client *client // Reference to the client that owns this fetcher.
}

// NewFetcher creates a new instance of apiFetcher with the provided client.
func NewFetcher(cl *client) *apiFetcher {
	return &apiFetcher{client: cl}
}

// FetchCodes retrieves a map of supported currency codes from the API.
func (c *apiFetcher) FetchCodes() (result map[string]iso4217, err error) {
	defer func() { err = errorhand.IfError("failed to fetch codes", err) }() // Ensure errors are handled properly.

	result = make(map[string]iso4217) // Initialize the result map.

	// Fetch response from the Codes endpoint.
	resp, err := c.Fetch(CodesEndpoint)
	if err != nil {
		return nil, err // Return error if fetching failed.
	}

	// Populate the result map with supported currency codes.
	for _, v := range resp.SupportedCodes {
		result[v[0]] = iso4217{
			Code: v[0],
			Name: v[1],
		}
	}

	return result, nil // Return the map of supported codes.
}

// FetchPair retrieves the exchange rate between two specified currencies.
func (c *apiFetcher) FetchPair(first, second string) (result *pair, err error) {
	defer func() { err = errorhand.IfError("failed to fetch pair", err) }() // Handle errors.

	// Fetch response from the Pair endpoint with the specified currencies.
	resp, err := c.Fetch(PairEndpoint, first, second)
	if err != nil {
		return nil, err // Return error if fetching failed.
	}

	// Create and return the pair result from the response.
	result = &pair{
		BaseCode:   resp.BaseCode,
		TargetCode: resp.TargetCode,
		Rate:       resp.ConversionRateFloat,
	}

	return result, nil
}

// FetchIndex retrieves the index of a specified currency.
func (c *apiFetcher) FetchIndex(code string) (result *index, err error) {
	defer func() { err = errorhand.IfError("failed to fetch index", err) }() // Handle errors.

	// Fetch response from the Index endpoint with the specified currency code.
	resp, err := c.Fetch(IndexEndpoint, code)
	if err != nil {
		return nil, err // Return error if fetching failed.
	}

	// Create and return the index result from the response.
	result = &index{
		BaseCode: resp.BaseCode,
		RateMap:  resp.ConversionRateMap,
	}

	return result, nil
}

// Fetch performs a general fetch operation for the specified endpoint.
func (c *apiFetcher) Fetch(toFetch endpoint, args ...string) (resp *responseFromServer, err error) {
	defer func() { err = errorhand.IfError("failed to fetch", err) }() // Handle errors.

	resp = &responseFromServer{} // Initialize the response structure.

	var (
		query     string // Query string to hold the constructed URL for the request.
		dateParse bool = true // Flag to indicate if date parsing is needed.
	)

	// Construct the query URL based on the requested endpoint.
	switch toFetch {
	case CodesEndpoint:
		query = pathConstruct(c.client.basePath, codesUrlKey) // Construct URL for fetching currency codes.
		dateParse = false // No date parsing needed for this endpoint.
	case PairEndpoint:
		if len(args) != 2 {
			return nil, fmt.Errorf("invalid args len: pair endpoint requires 2 args") // Check for valid argument count.
		}
		query = pathConstruct(c.client.basePath, pairUrlKey, args[0], args[1]) // Construct URL for fetching currency pair rates.
	case IndexEndpoint:
		if len(args) != 1 {
			return nil, fmt.Errorf("invalid args len:index endpoint requires 1 args") // Check for valid argument count.
		}
		query = pathConstruct(c.client.basePath, latestUrlKey, args[0]) // Construct URL for fetching index rates.
	default:
		return nil, fmt.Errorf("Unsupported endpoint: %s", toFetch) // Return error for unsupported endpoints.
	}

	// Execute the HTTP request and handle any errors.
	reqResp, err := c.client.doRequest(query)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %s", query)
	} else if reqResp == nil {
		return nil, fmt.Errorf("failed to do request: empty response") // Handle empty response case.
	}

	// Unmarshal the JSON response into the response structure.
	if err = json.Unmarshal(reqResp, resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %v", err) // Handle JSON unmarshal errors.
	}

	// Parse the date strings if needed.
	if dateParse {
		if resp.NextUpdateDateTime, err = time.Parse(time.RFC1123, resp.NextUpdateString); err != nil {
			log.Printf("failed to parse into RFC1123: %s", resp.NextUpdateString) // Log any parsing errors.
		}

		if resp.LastUpdateDateTime, err = time.Parse(time.RFC1123, resp.LastUpdateString); err != nil {
			log.Printf("failed to parse into RFC1123: %s", resp.LastUpdateString) // Log any parsing errors.
		}
	}

	return resp, nil // Return the populated response structure.
}
