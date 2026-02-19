// go get -u golang.org/x/net/http2

package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http2.ConfigureServer(server, &http2.Server{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP/2!")
	})

	log.Printf("Starting server at port 8080")
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
