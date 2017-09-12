package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

// ModelIndex - handle function for public model route
func ModelIndex(w http.ResponseWriter, r *http.Request) {
	ml := ModelList{
		Model{ID: "0815-4223", Name: "random model 1", CreationDate: time.Now()},
		Model{ID: "007-BOND", Name: "random model 2", CreationDate: time.Now()},
	}

	json.NewEncoder(w).Encode(ml)
}

// ModelDetail - handle function for public model detail route
func ModelDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modelID := vars["modelId"]
	fmt.Fprintln(w, "model:", modelID)
}
