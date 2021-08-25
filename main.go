package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

// Init books var as a slice book struct
// slice : similar to array 
var books []Book

// get all books
func getBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

// get particular book
func getBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	// get parameters
	params := mux.Vars(r) 
	// loop through books and find with id
	for _,item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// create a new book
func createBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// mock id 
	book.ID = strconv.Itoa(rand.Intn(999999))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// update book
func updateBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

//delete a particular book
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


func main(){
	// initializing router
	router := mux.NewRouter()

	// Mock Data
	books = append(books,Book{ ID:"1" ,Isbn: "433",Title: "Book1", Author: &Author{FirstName:"Aniket",LastName:"Pal"}})
	books = append(books,Book{ ID:"2" ,Isbn: "434",Title: "Book2", Author: &Author{FirstName:"Aniket2",LastName:"Pal2"}})
	books = append(books,Book{ ID:"3" ,Isbn: "435",Title: "Book3", Author: &Author{FirstName:"Aniket3",LastName:"Pal3"}})



	// routes
	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

	// running server
	log.Fatal(http.ListenAndServe(":8080",router));




}