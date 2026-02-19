package main

import (
	"fmt"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r) // Pass execution to the next handler
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP server!")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	// Wrap the multiplexer with the middleware
	wrappedMux := loggingMiddleware(mux)

	fmt.Println("Server is starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
