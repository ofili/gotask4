package main

import (
	"gotask4/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Route Definition
func main() {

	router := mux.NewRouter()
	//Create
	router.HandleFunc("/book", api.CreateBook).Methods("POST")
	//Read
	router.HandleFunc("/book/{book_id}", api.GetBook).Methods("GET")
	//Read all
	router.HandleFunc("/book", api.GetBooks).Methods("GET")
	//Update
	router.HandleFunc("/book/{book_id}", api.UpdateBook).Methods("PUT")
	//Delete
	router.HandleFunc("/book/{book_id}", api.DeleteBook).Methods("DELETE")
	api.InitDB()

	log.Fatal(http.ListenAndServe(":8000", router))
}