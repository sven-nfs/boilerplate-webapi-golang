package main

import "time"

// Model - type for storing data
type Model struct {
	ID           int    	`json:"id"`
	Name         string    	`json:"name"`
	CreationDate time.Time 	`json:"created"`
}

// ModelList - slice of model data
type ModelList []Model
