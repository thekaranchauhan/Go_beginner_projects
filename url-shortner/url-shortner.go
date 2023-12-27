package main

import (
	"fmt"
	"net/http"
)

// URLShortener is a simple URL shortener.
type URLShortener struct {
	urlMap map[string]string // Map to store the short key to original URL mapping
}

// NewURLShortener creates a new instance of URLShortener.
func NewURLShortener() *URLShortener {
	return &URLShortener{
		urlMap: make(map[string]string),
	}
}

// shortenURL is a handler for the /shorten endpoint. It shortens a given URL and returns the short URL.
func (u *URLShortener) shortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Get the original URL from the form
	originalURL := r.Form.Get("url")
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Generate a unique short key (for simplicity, you might want to use a hash function)
	shortKey := fmt.Sprintf("%d", len(u.urlMap)+1)

	// Store the mapping of short key to original URL
	u.urlMap[shortKey] = originalURL

	// Send the short URL back in the response
	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)
	fmt.Fprintf(w, "Shortened URL: %s\n", shortURL)
}

// redirectToOriginalURL is a handler for redirecting to the original URL based on the short key.
func (u *URLShortener) redirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[1:]

	// Retrieve the original URL from the map
	originalURL, ok := u.urlMap[shortKey]
	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

func main() {
	// Create a new instance of URLShortener
	urlShortener := NewURLShortener()

	// Handle URL shortening requests at the "/shorten" endpoint
	http.HandleFunc("/shorten", urlShortener.shortenURL)

	// Handle redirects to original URLs at the "/" endpoint
	http.HandleFunc("/", urlShortener.redirectToOriginalURL)

	// Start the server on port 8080
	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
