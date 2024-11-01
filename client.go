package exchrateclient

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/ayayaakasvin/exchrateclient/lib/errorhand" // Custom error handling package
)

// Base URL for the exchange rate API
const (
	host = "https://v6.exchangerate-api.com/v6"
)

// client struct contains necessary information for making API requests.
type client struct {
	basePath string     // The full path for API requests, including the base URL and API key.
	api      string     // API key for accessing the exchange rate API.
	client   *http.Client // HTTP client for making requests.
	*apiFetcher // Embedded struct for fetching API data.
}

// defaultClient initializes a client with default settings and a new HTTP client.
func defaultClient() *client {
	cl := &client{
		client:   &http.Client{}, // Create a new HTTP client.
		basePath: host,           // Set the base path to the API host.
	}

	cl.apiFetcher = NewFetcher(cl) // Create a new apiFetcher for making requests.

	return cl // Return the initialized client.
}

// clientWithApi initializes a client with a specified API key.
func clientWithApi(API string) *client {
	cl := &client{
		api:    API,                     // Set the API key.
		client: &http.Client{},          // Create a new HTTP client.
	}

	cl.basePath = pathConstruct(host, cl.api) // Construct the full path using the host and API key.

	cl.apiFetcher = NewFetcher(cl) // Create a new apiFetcher for making requests.

	return cl // Return the initialized client.
}

// SetAPI sets the API key and updates the base path accordingly.
func (cl *client) SetAPI(APIkey string) {
	cl.api = APIkey // Update the API key.
	cl.basePath = path.Join(host, cl.api) // Update the base path using the new API key.
}

// doRequest executes an HTTP GET request for a given query string.
func (c *client) doRequest(query string) (result []byte, err error) {
	defer func() {
		// Ensure any errors are handled properly using the custom error handler.
		err = errorhand.IfError("failed to do request", err)
	}()

	if query == "" {
		return nil, fmt.Errorf("empty query") // Return an error if the query string is empty.
	}

	req, err := http.NewRequest(http.MethodGet, query, nil) // Create a new GET request.
	if err != nil {
		return nil, err // Return any errors encountered while creating the request.
	}

	responce, err := c.client.Do(req) // Execute the request using the HTTP client.
	if err != nil {
		return nil, err // Return any errors encountered during the request.
	} else if responce.StatusCode > 400 && responce.StatusCode < 500 {
		return nil, fmt.Errorf("client error: status code %d", responce.StatusCode) // Handle client errors based on the status code.
	}

	defer responce.Body.Close() // Ensure the response body is closed after reading.

	if result, err = io.ReadAll(responce.Body); err != nil {
		return nil, err // Return any errors encountered while reading the response body.
	}

	return result, nil // Return the result of the request.
}

// pathConstruct joins the provided path segments into a single path string.
func pathConstruct(args ...string) string {
	return strings.Join(args, "/") // Join the segments with a "/" separator.
}
