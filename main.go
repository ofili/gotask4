package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type book struct {
	ID              string          `json:"ID"`
	Name            string      `json:"Name"`
	Author          string      `json:"Author"`
	Published_at    string      `json:"published_at"`
}

type allBooks []book

var novel = allBooks{
	{
		ID:          "1",
		Name:       "Silmarilion",
		Author: "JRR Tolkein",
		Published_at: "1977-08-15T08:00:00",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//create an book
func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the book title, author and publisher only in order to update")
	}

	json.Unmarshal(reqBody, &newBook)
	novel = append(novel, newBook)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newBook)
}

// get one book
func getOneBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["id"]

	for _, singleBook := range novel {
		if singleBook.ID == bookID {
			json.NewEncoder(w).Encode(singleBook)
		}
	}
}
// get all books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(novel)
}

//update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["id"]
	var updatedBook book

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the book title, author and publisher only in order to update")
	}
	json.Unmarshal(reqBody, &updatedBook)

	for i, singleBook := range novel {
		if singleBook.ID == bookID {
			singleBook.Name = updatedBook.Name
			singleBook.Author = updatedBook.Author
			singleBook.Published_at = updatedBook.Published_at
			novel = append(novel[:i], singleBook)
			json.NewEncoder(w).Encode(singleBook)
		}
	}
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["id"]

	for i, singleBook := range novel {
		if singleBook.ID == bookID {
			novel = append(novel[:i], novel[i+1:]...)
			fmt.Fprintf(w, "The book with ID %v has been deleted successfully", bookID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/book", createBook).Methods("POST")
	router.HandleFunc("/novel", getAllBooks).Methods("GET")
	router.HandleFunc("/novel/{id}", getOneBook).Methods("GET")
	router.HandleFunc("/novel/{id}", updateBook).Methods("PATCH")
	router.HandleFunc("/novel/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}