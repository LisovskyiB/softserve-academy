package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"softserve-academy/my_easy_server/server"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/name", server.Handler).Methods("GET")

	log.Fatal(http.ListenAndServe(":2006", r))
}
