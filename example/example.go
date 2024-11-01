package main

import (
	"fmt"
	"log"

	"github.com/ayayaakasvin/exchrateclient"
)

func main() {
	// Create a new client instance with your API key
	cl := exchrateclient.New("your-api-key")

	// Example 1: Fetch available currency codes
	fetchCurrencyCodes(cl)

	// Example 2: Fetch exchange rate for a specific currency pair
	fetchCurrencyPair(cl, "USD", "EUR")

	// Example 3: Fetch index rates for a base currency
	fetchCurrencyIndex(cl, "USD")
}

// Example 1: Fetch available currency codes
func fetchCurrencyCodes(cl exchrateclient.Client) {
	fmt.Println("Fetching available currency codes...")
	codesMap, err := cl.FetchCodes()
	if err != nil {
		log.Fatalf("Error fetching codes: %s", err)
		return
	}

	if len(codesMap) == 0 {
		fmt.Println("No currency codes found.")
		return
	}

	fmt.Println("Available currency codes:")
	for code, name := range codesMap {
		fmt.Printf("%s: %s\n", code, name)
	}
}

// Example 2: Fetch exchange rate for a specific currency pair
func fetchCurrencyPair(cl exchrateclient.Client, base string, target string) {
	fmt.Printf("Fetching exchange rate from %s to %s...\n", base, target)
	pair, err := cl.FetchPair(base, target)
	if err != nil {
		log.Fatalf("Error fetching pair: %s", err)
		return
	}

	if pair == nil {
		fmt.Println("No exchange rate data found for this pair.")
		return
	}

	fmt.Printf("%s to %s rate: %.4f\n", pair.BaseCode, pair.TargetCode, pair.Rate)
}

// Example 3: Fetch index rates for a base currency
func fetchCurrencyIndex(cl exchrateclient.Client, base string) {
	fmt.Printf("Fetching index rates for base currency %s...\n", base)
	index, err := cl.FetchIndex(base)
	if err != nil {
		log.Fatalf("Error fetching index: %s", err)
		return
	}

	if index == nil || len(index.RateMap) == 0 {
		fmt.Println("No index rates found for this base currency.")
		return
	}

	fmt.Printf("Exchange rates for %s:\n", base)
	for code, rate := range index.RateMap {
		fmt.Printf("%s: %.4f\n", code, rate)
	}
}
