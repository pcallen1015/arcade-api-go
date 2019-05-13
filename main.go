package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"pcallen1015/arcade-api-go/database"
	"pcallen1015/arcade-api-go/wins"
)

const (
	apiPrefix = "/api/v1"
)

func main() {
	database.Connect()

	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ping")
		w.Write([]byte("Ping"))
	}).Methods("GET")

	router.HandleFunc("/wins", wins.ListHandler).Methods("GET")
	router.HandleFunc("/wins", wins.CreateHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
