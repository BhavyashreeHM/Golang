package main

import (
	"fmt"
	"net/http"
)

type Employee struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Project    string `json:"project"`
	Department string `json:"department"`
}

var employeeOne Employee

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("This is main function")
		fmt.Fprintln(w, "This is Home Page")
	})
	mux.HandleFunc("GET /employee", GetEmployee)
	mux.HandleFunc("POST /employee", POSTEmployee)
	mux.HandleFunc("PUT /employee", PUTEmployee)
	mux.HandleFunc("PATCH /employee", PATCHEmployee)
	mux.HandleFunc("DELETE /employee", DELETEEmployee)

	port := ":3000"
	fmt.Println("Server Up and Running on port", port)
	http.ListenAndServe(port, mux)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is GetEmployee Page")
}

func POSTEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is PostEmployee Page")
}

func PUTEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is PUtEmployee Page")
}

func PATCHEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is PAtchEmployee Page")
}

func DELETEEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is DELEEmployee Page")
}
