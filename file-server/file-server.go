package main

import (
	"net/http"
)

func main() {
	// Specify the directory you want to serve files from
	dir := "./files"

	// Create a file server handler
	fileServer := http.FileServer(http.Dir(dir))

	// Register the file server handler with a specific route
	http.Handle("/files/", http.StripPrefix("/files/", fileServer))

	// Set up a basic handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is a file server!"))
	})

	// Set the port to listen on
	port := "8080"

	// Start the server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
