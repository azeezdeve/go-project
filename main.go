package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-project/controller"
	"go-project/model"
	"net/http"
)

func main() {
	controller.InitStorage()
	r := mux.NewRouter()
	r.HandleFunc("/", CreateBookHandler).Methods("POST")
	r.HandleFunc("/", UpdateBookHandler).Methods("PATCH")
	r.HandleFunc("/{id}", GetBookHandler).Methods("GET")
	r.HandleFunc("/", GetBooksHandler).Methods("GET")
	r.HandleFunc("/{id}", DeleteBookHandler).Methods("DELETE")
	fmt.Println("server starting")
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("server running")
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	//decode to the struct
	books := controller.NewBooks()
	var body model.Books
	_ = json.NewDecoder(r.Body).Decode(&body)
	err := books.CreateBook(context.Background(), body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	books := controller.NewBooks()
	var body model.Books
	_ = json.NewDecoder(r.Body).Decode(&body)
	err := books.UpdateBook(context.Background(), body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusAccepted)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	books := controller.NewBooks()
	vars := mux.Vars(r)
	book, err := books.GetBookByID(context.Background(), vars["id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	bookBytes, _ := json.Marshal(book)
	w.Write(bookBytes)
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := controller.NewBooks()
	book, err := books.GetListOfBook(context.Background())
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	bookBytes, _ := json.Marshal(book)
	w.Write(bookBytes)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	books := controller.NewBooks()
	vars := mux.Vars(r)
	err := books.DeleteBook(context.Background(), vars["id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusAccepted)
}
