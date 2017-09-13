package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index - handle function for index route
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome!")
}

// ModelIndex - handle function for public model route
func ModelIndex(w http.ResponseWriter, r *http.Request) {
	//set some header info
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(models); err != nil {
		panic(err)
	}
}

// ModelDetail - handle function for public model detail route
func ModelDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modelID, err := strconv.Atoi(vars["modelId"])
	model, err := Read(modelID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	//set some header info
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(model); err != nil {
		panic(err)
	}
}
