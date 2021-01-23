package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
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

// Init Books Variable as a slice Book Struct
var books []Book

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, 
		Book{
			ID: "1", 
			ISBN: "7777", 
			Title: "First Book", 
			Author: &Author{
				FirstName: "Younhong", 
				LastName: "Ko",
			},
		},
	)
	books = append(books, 
		Book{
			ID: "2", 
			ISBN: "5555", 
			Title: "Second Book", 
			Author: &Author{
				FirstName: "Sumin", 
				LastName: "Lee",
			},
		},
	)

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get All Books
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Loop through book and find by id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)		
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Get All Books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// Get All Books
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Get All Books
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
