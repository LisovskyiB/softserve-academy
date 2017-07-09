package server

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type NameStruct struct {
	Name string `json:"name"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	if checkParams(r) {

		returnResult(w, makeJsonString(getFields(r)[`name`][0]))
	}
}

func getFields(r *http.Request) url.Values {

	r.ParseForm()

	data := r.Form

	return data
}

func makeJsonString(name string) string{

	return string(doMarshal(name))
}


func doMarshal(name string) []byte{

	jsonFormat, err := json.Marshal(generateNemeStruct(name))

	if err != nil {

		log.Printf("%+v\n", "Json Marshal error", err.Error())

		return []byte{}
	}

	return jsonFormat
}

func generateNemeStruct(name string) NameStruct {

	return NameStruct{
		Name : name,
	}
}

func returnResult(w http.ResponseWriter, str string) {

	fmt.Printf("%+v\n", str)

	fmt.Fprintf(w, str)
}

func checkParams(r *http.Request) bool{

	if r.Method != "GET" {
		return false
	}

	if r.URL.Path[1:] == `name` {

		return true
	}

	return false
}
