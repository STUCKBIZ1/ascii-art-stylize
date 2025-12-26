package main

import (
	"fmt"
	"net/http"

	ascii "ascii_art_web/src"
)

func main() {
	// Register the handler for the home page (GET /)
	http.HandleFunc("/", ascii.HomeHandler)

	// Register the handler for ASCII art generation (POST /ascii-art)
	http.HandleFunc("/ascii-art", ascii.AsciiArtHandler)

	// Print server URL to the console
	fmt.Println("Server running at http://localhost:8080")

	// Start the HTTP server on port 8080
	// nil means we use the default ServeMux
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// If the server fails to start, print the error
		fmt.Println("Error starting server:", err)
	}
}
