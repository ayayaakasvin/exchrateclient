package exchrateclient

// Client interface combines methods for setting the API and fetching exchange rate data.
type Client interface {
	SetAPI(string) // Method to set the API key or endpoint.
	fetcher       // Embedding the fetcher interface to inherit its methods.
}

// New creates a new Client instance with the provided API string.
// It initializes the client with the specified API endpoint.
func New(API string) Client {
	return clientWithApi(API) // Calls a function to create a client with the given API.
}

// Default returns a default Client instance.
// This function can be used to obtain a pre-configured client without needing to specify an API key.
func Default() Client {
	return defaultClient() // Calls a function to create a default client instance.
}
