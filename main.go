package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// Get All Books
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Get All Books
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Get All Books
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Get All Books
func deleteBook(w http.ResponseWriter, r *http.Request) {

}
