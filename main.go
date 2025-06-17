package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Default port to listen on
const defaultPort = "8080"

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Register the handler function for all paths
	http.HandleFunc("/", handleRequest)

	// Start the server
	address := fmt.Sprintf(":%s", port)
	fmt.Printf("Starting HTTP header dump server on port %s...\n", port)
	fmt.Println("Press Ctrl+C to stop the server")
	
	log.Fatal(http.ListenAndServe(address, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Create a map to store headers
	headers := make(map[string][]string)

	// Copy all headers from request
	for name, values := range r.Header {
		headers[name] = values
	}

	// Marshal headers to JSON
	jsonData, err := json.MarshalIndent(headers, "", "  ")
	if err != nil {
		log.Printf("Error marshalling headers to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Print the JSON to console
	fmt.Printf("Request from %s %s\n", r.Method, r.URL.Path)
	fmt.Println(string(jsonData))
	
	// Return 200 OK response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
