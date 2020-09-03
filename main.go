package main

import (
	"fmt"
	"log"
	"net/http"
	"ramonmoraes/httq/infra"

	"github.com/gorilla/mux"
)

type HTTQ interface {
	GetMessage(w http.ResponseWriter, r *http.Request)
	WriteMessage(w http.ResponseWriter, r *http.Request)
	GetPrefix() string
}

func main() {
	router := mux.NewRouter()
	infra := &infra.KafkaHTTQ{}

	path := fmt.Sprintf("/%s/{key}", infra.GetPrefix())
	router.HandleFunc(path, infra.GetMessage).Methods("GET")
	router.HandleFunc(path, infra.WriteMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
