package main

import (
	"fmt"
)

// ToDoItem :: an item that needs to be done
type ToDoItem struct {
	ID          int    `json="id"`
	Name        string `json="name"`
	Done        bool   `json="done"`
	CreatedOn   int64  `json="created"`
	CompletedOn int64  `json="completed"`
}

// ToMap :: returns map[string]interface{} of ToDoItem
func (item ToDoItem) ToMap() map[string]interface{} {
	value := map[string]interface{}{
		"id":          item.ID,
		"name":        item.Name,
		"done":        item.Done,
		"createdOn":   item.CreatedOn,
		"completedOn": item.CompletedOn,
	}

	return value
}
