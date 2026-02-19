package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var books = map[string]Book{}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers)
	mux.HandleFunc("/books", Getbook)
	mux.HandleFunc("/books/{id}", Getbook)
	mux.HandleFunc("/addbook", Createbook)
	mux.HandleFunc("/deletebook", Deletebook)
	log.Println("Server Running on :8080")
	http.ListenAndServe(":8080", mux)
}
func handlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server up and Runing ...\n"))
}

func Createbook(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting books
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if book.ID == "" {
		book.ID = strconv.Itoa(len(books) + 1)
	}
	_, exists := books[book.ID]
	if exists {
		http.Error(w, "Book with this ID already exists", http.StatusConflict)
		return
	}
	books[book.ID] = book
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func Getbook(w http.ResponseWriter, r *http.Request) {
	// Implementation for adding a book
}

func Deletebook(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a book
}
