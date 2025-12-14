package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		fmt.Fprintln(w, "Hello, Server!")
	})

	// const addString string = "127.0.0.1:8080"
	// port := 3000
	// println("Server is running at address", ":3000")

	http.ListenAndServe(":3000", nil)

	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	} )

	http.ListenAndServe(":3000", r)
	// 
	r.HandleFunc("/orders",orders)
	r.HandleFunc("/carts",cart)
	r.HandleFunc("/products",products)

}
