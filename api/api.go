package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Book represents the model for a book
//Default table name will be `books`
type Book struct {
	//gorm.Model
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

var books []Book

//Initial DB setup and config
var db *gorm.DB

//Init
func InitDB() {
	var err error
	db, err = gorm.Open("mysql", "root:+*P@ssw0rd)@tcp(localhost:3306)/library_db?parseTime=True")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	log.Println("Connection established")

	//Creat the tables. This is a  one-time step
	//Comment out if running multiple times - You may see an error otherwise
	//db.CreateTable(&Book{})
	db.HasTable(&Book{})
	db.DropTableIfExists(&Book{})

	//Migration to create tables for library_db schema
	db.AutoMigrate(&Book{})
}

//Create api
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	//Create new books by inserting records in the books table
	db.Create(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

//Read-all api
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

//Read api
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputBookID := params["id"]
	var book Book
	db.First(&book, inputBookID)
	json.NewEncoder(w).Encode(books)
}

//Update api
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook Book
	json.NewDecoder(r.Body).Decode(&updatedBook)
	db.Save(&updatedBook)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

//Delete api
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := r.URL.Query()["id"]
	db.Where("id = ?", param).Delete(&books)
	json.NewEncoder(w).Encode(books)
}
