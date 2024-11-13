package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/shubhangNarain/go-bms/pkg/models"
	"github.com/shubhangNarain/go-bms/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err!= nil {
        fmt.Println("Error parsing book")
    }
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err!= nil {
        fmt.Println("Error parsing book")
    }
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	vars updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
    vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("Error parsing book")
	}
	bookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if UpdateBook.Author!= ""{
        bookDetails.Author = updateBook.Author
    }
	if UpdateBook.Publication!= ""{
        bookDetails.Publication = updateBook.Publication
    }
	db.Save(&bookDetails)
	db, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}