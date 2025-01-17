package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func Add(value1 int, value2 int) int {
	return value1 + value2
}

func Substract(value1 int, value2 int) int {
	return value1 - value2
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200);
	response.Write([]byte("Hello world"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
