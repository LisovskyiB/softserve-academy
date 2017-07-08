package main

import (
	"net/http"
	"softserve-academy/my_easy_server/server"
)

func main() {

	http.HandleFunc("/name", server.Handler)

	http.ListenAndServe(":2005", nil)

	select {}
}
