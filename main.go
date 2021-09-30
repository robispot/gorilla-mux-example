package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Golang pointers", Author: "Robbie", Year: "2011"},
		Book{ID: 2, Title: "Golang pointers2", Author: "Robbie2", Year: "2011"},
		Book{ID: 3, Title: "Golang pointers3", Author: "Robbie3", Year: "2011"},
		Book{ID: 4, Title: "Golang pointers4", Author: "Robbie4", Year: "2011"})
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params)

	//convert string to integer and return
	i, _ := strconv.Atoi(params["id"])
	log.Println("check id type: ", reflect.TypeOf(i))

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)
	log.Println(books)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove book")
}
