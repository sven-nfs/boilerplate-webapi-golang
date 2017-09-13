package main

import (
	"fmt"
	"time"
)

var currentID int

var models ModelList

func Initialze() {
	Create(Model{Name: "Franz Oberhauser"})
	Create(Model{Name: "James Bond"})
	Create(Model{Name: "Dr. No"})
	Create(Model{Name: "Der Beißer"})
	Create(Model{Name: "Miss Moneypenny"})
	Create(Model{Name: "Q"})
}

func Create(m Model) error {
	currentID++
	m.ID = currentID
	m.CreationDate = time.Now()
	models = append(models, m)
	return nil
}

func Read(id int) (Model, error) {
	for _, m := range models {
		if m.ID == id {
			return m, nil
		}
	}
	// return empty model if not found
	return Model{}, fmt.Errorf("model not found")
}

func Update(m Model) error {
	for i, modelItem := range models {
		if m.ID == modelItem.ID {
			//stupid swap - fix this
			tmp := append(models[:i], m)
			tmp = append(tmp, models[i+1:]...)
			models = tmp
			return nil
		}
	}
	return fmt.Errorf("Could not find model with id of %d to update", m.ID)
}

func Delete(m Model) error {
	for i, modelItem := range models {
		if modelItem.ID == m.ID {
			models = append(models[:i], models[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find model with id of %d to delete", m.ID)
}
