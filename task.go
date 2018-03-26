package main

import (
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/datastore"
)

// Task struct
type Task struct {
	Description string    `datastore:"description"`
	Created     time.Time `datastore:"created"`
}

// NewTask instantiate Task
func NewTask() *Task {
	return &Task{
		Created: time.Now(),
	}
}

// Save Task entity.
func (e *Task) Save() {
	k := datastore.IncompleteKey("Task", nil)

	if _, err := dsClient.Put(dsCtx, k, e); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New Task with description %q\n", e.Description)
}
