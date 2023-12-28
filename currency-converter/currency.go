package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	apiEndpoint = "https://api.exchangerate-api.com/v4/latest/USD"
)

// ExchangeRates represents the structure of the response from the exchange rate API.
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

// getExchangeRates retrieves the latest exchange rates from the API.
func getExchangeRates() (ExchangeRates, error) {
	resp, err := resty.New().R().Get(apiEndpoint)
	if err != nil {
		return ExchangeRates{}, err
	}

	var exchangeRates ExchangeRates
	if err := json.Unmarshal(resp.Body(), &exchangeRates); err != nil {
		return ExchangeRates{}, err
	}

	return exchangeRates, nil
}

// convertCurrency converts the given amount from one currency to another.
func convertCurrency(amount float64, fromCurrency, toCurrency string) (float64, error) {
	// Get the latest exchange rates
	exchangeRates, err := getExchangeRates()
	if err != nil {
		return 0, err
	}

	// Retrieve the exchange rates for the specified currencies
	fromRate, fromExists := exchangeRates.Rates[fromCurrency]
	toRate, toExists := exchangeRates.Rates[toCurrency]

	// Check if both currencies are valid
	if !fromExists || !toExists {
		return 0, fmt.Errorf("invalid currency code")
	}

	// Perform the currency conversion
	result := (amount / fromRate) * toRate
	return result, nil
}

func main() {
	// Set the initial values for currency conversion
	amount := 100.0
	fromCurrency := "USD"
	toCurrency := "CAD"

	// Perform the currency conversion
	convertedAmount, err := convertCurrency(amount, fromCurrency, toCurrency)
	if err != nil {
		log.Fatal(err)
	}

	// Display the result
	fmt.Printf("%.2f %s is equal to %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
