package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// --------------------------------------------
// Model for Book and Author
// --------------------------------------------

// Book struct defines the structure of a book
type Book struct {
	BookId string  `json:"bookId"` // Tip: Tag defines how this field will appear in JSON
	Name   string  `json:"name"`
	Author *Author `json:"author"`
	Price  float32 `json:"price"`
}

// Author struct represents book's author
type Author struct {
	AuthorName string `json:"authorName"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
}

// --------------------------------------------
// In-memory "fake" DB
// --------------------------------------------
var books []Book

// --------------------------------------------
// Middleware / Helper Functions
// --------------------------------------------

// IsEmpty checks if the Book is empty (missing Name)
// Tip: Helps validate input during creation/update
func (b *Book) IsEmpty() bool {
	return b.Name == ""
}

// --------------------------------------------
// Main function
// --------------------------------------------
func main() {
	// Tip: Set up router and routes
	r := mux.NewRouter()

	// seeding
	books = append(books, Book{BookId: "1", Name: "The Art of Being Alone", Price: 300, Author: &Author{AuthorName: "Renuka Gavrani", Email: "renuka@gavrani.com", IsVerified: true}})
	books = append(books, Book{BookId: "2", Name: "The Art of Being Alone PT-II", Price: 300, Author: &Author{AuthorName: "Renuka Gavrani", Email: "renuka@gavrani.com", IsVerified: true}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{bookId}", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/book/{bookId}", updateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", deleteBook).Methods("DELETE")

	// Tip: Start server
	fmt.Println("Server is running on :4000")
	http.ListenAndServe(":4000", r)
}

// --------------------------------------------
// Controllers / Handlers
// --------------------------------------------

// serveHome serves a simple welcome message
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to home page</h1>"))
}

// getBooks returns all books in JSON format
func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all books")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// getBook fetches a book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a book")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Tip: Extract path parameters
	for _, book := range books {
		if book.BookId == params["bookId"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode("Unable to find a book with given ID")
}

// createBook handles adding a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a book")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	if book.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// Tip: Use dedicated rand.Rand instance (Go 1.20+ best practice)
	randInstance := rand.New(rand.NewSource(time.Now().UnixNano()))
	book.BookId = strconv.Itoa(randInstance.Intn(1000)) // Tip: Generate unique ID

	books = append(books, book)

	json.NewEncoder(w).Encode("Book created successfully")
}

// updateBook replaces a book by ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a book")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var updatedBook Book
	_ = json.NewDecoder(r.Body).Decode(&updatedBook)

	if updatedBook.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	params := mux.Vars(r)

	// Tip: Update by removing old and adding updated
	for index, book := range books {
		if book.BookId == params["bookId"] {
			books = append(books[:index], books[index+1:]...) // Remove old book
			updatedBook.BookId = params["bookId"]             // Keep same ID
			books = append(books, updatedBook)                // Add updated book
			json.NewEncoder(w).Encode("Book updated successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("Book not found")
}

// deleteBook removes a book by ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a book")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, book := range books {
		if book.BookId == params["bookId"] {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode("Book deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("Book not found")
}
