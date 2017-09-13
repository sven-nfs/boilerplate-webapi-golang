package main

import (
	"log"
	"net/http"
)

func main() {

	Initialze()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
