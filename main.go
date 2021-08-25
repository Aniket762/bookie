package main

import (
	// "encoding/json"
	"log"
	"net/http"

	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Author model
type Author struct{
	FirstName	string `json:"firstname"`
	LastName	string `json:"lastname"`
	
}

// Book model
type Book struct {
	ID		string		`json:"id"`
	Isbn	string		`json:"isbn"`
	Title	string		`json:"title"`
	Author	*Author		`json:"author"`
}



func main(){
	// initializing router
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

	// running server
	log.Fatal(http.ListenAndServe(":8080",router))


}