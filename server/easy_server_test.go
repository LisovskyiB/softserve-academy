package server

import (
	"testing"
	"reflect"
)

var json_string = `{"name":"Bohdan"}`

var bytes_string = []byte{123, 34, 110, 97, 109, 101, 34, 58, 34, 66, 111, 104, 100, 97, 110, 34, 125}


func TestMakeJsonString(t *testing.T) {

	if makeJsonString("Bohdan") != json_string {
		t.Error()
	}
}

func TestDoMarshal(t *testing.T) {

	if !reflect.DeepEqual(doMarshal("Bohdan"), bytes_string) {
		t.Error()
	}
}

func TestGenerateNemeStruct(t *testing.T) {

	name := NameStruct{
		Name:"Bohdan",
	}


	if !reflect.DeepEqual(generateNemeStruct("Bohdan"), name) {
		t.Error()
	}
}
