# Golang Exchange Rate API Client ([exchangerate-api](https://www.exchangerate-api.com/))

## Overview

The Go Exchange Rate API Client is a simple and lightweight client for accessing exchange rate data. It provides a convenient interface for retrieving current exchange rates with daily updates for various currencies. An API key is free to obtain; you must register at the link above to get a free API key for 1500 requests. You can also find the documentation for the Exchange Rate API there.

## Features

- Fetch current exchange rates for pairs of specific currencies.
- Fetch current exchange rates indexed by a specific currency.
- Over 150+ supported *ISO4217* currency codes.

## Getting Started

### Prerequisites

- Go 1.23 or higher

### Installation

You can clone the package or import it into your project.

1. ### Clone the repository:

   ```bash
   git clone https://github.com/ayayaakasvin/exchrateclient.git
   cd exchrateclient
   ```
    You can modify files and extend the current state of the package. Feel free to contribute!

2. ### Import into Your Project
    ```bash
    go get github.com/ayayaakasvin/exchrateclient
    ```
    This package uses only standard libraries, so there’s no need for external dependencies.

## Usage
To  start using the package imported into your project, follow these examples:

## Create client:
```go
cl := exchrateclient.New("your-api-key") // change with your actual key
// how to get key was explained above
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
```
        
## Fetch the Index of an ISO4217 Code
```go
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
```

## Fetch the Index of an ISO4217 Code
```go
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
```
## Fetch the Index of an ISO4217 Code
```go
resp, err := cl.Fetch(exchrateclient.PairEndpoint, "USD", "KZT")
if err != nil {
    log.Fatalf("Error fetching codes: %s", err)
}
```
The Fetch method is useful for obtaining the full response, which can be beneficial if you plan to use the client in your project more extensively.

## About response from API server
Response from server([exchangerate-api](https://www.exchangerate-api.com/)) can be differ, depending on endpoint, therefore can send JSON file with different fields. You can learn about it inside types.go file 

## Contact
For questions or feedback, please reach out to kozhamseitov06@gmail.com.