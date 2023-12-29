package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Quote structure to unmarshal JSON response
type Quote struct {
	Content  string   `json:"content"`
	Author   string   `json:"author"`
	Category []string `json:"tags"`
}

func getRandomQuote() (Quote, error) {
	// Make a GET request to the Quotable API
	response, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		return Quote{}, err
	}
	defer response.Body.Close()

	// Decode the JSON response into a Quote struct
	var quote Quote
	err = json.NewDecoder(response.Body).Decode(&quote)
	if err != nil {
		return Quote{}, err
	}

	return quote, nil
}

func main() {
	// Fetch a random quote
	quote, err := getRandomQuote()
	if err != nil {
		fmt.Printf("Error fetching quote: %s\n", err)
		return
	}

	// Display the quote
	fmt.Printf("Random Quote:\n%s\n- %s\nCategory: %s\n", quote.Content, quote.Author, quote.Category)
}
