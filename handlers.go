package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Index - handle function for index route
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome!")
}

// ModelIndex - handle function for public model route
func ModelIndex(w http.ResponseWriter, r *http.Request) {

	//grab this data from a db or whatever
	ml := ModelList{
		Model{ID: "0815-4223", Name: "random model 1", CreationDate: time.Now()},
		Model{ID: "007-BOND", Name: "random model 2", CreationDate: time.Now()},
	}

	//set some header info
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ml); err != nil {
		panic(err)
	}
}

// ModelDetail - handle function for public model detail route
func ModelDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modelID := vars["modelId"]
	fmt.Fprintln(w, "model:", modelID)
}
