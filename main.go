package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//retrieve a new mux router
	router := mux.NewRouter().StrictSlash(true)

	//add a index route to the router
	router.HandleFunc("/", Index)

	//add basic public routes
	router.HandleFunc("/public/models", ModelIndex)
	router.HandleFunc("/public/models/{modelId}", ModelDetail)

	//listen and serve
	//log errors
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index - handle function for index route
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome!")
}

func ModelIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list of all models")
}

func ModelDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modelId := vars["modelId"]
	fmt.Fprintln(w, "model:", modelId)
}
