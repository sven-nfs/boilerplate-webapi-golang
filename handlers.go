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
