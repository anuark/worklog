package main

import (
	"time"

	"cloud.google.com/go/datastore"
)

// Task struct
type Task struct {
	Key         *datastore.Key `datastore:"__key__" json:"id"`
	Description string         `datastore:"description" json:"description"`
	Created     time.Time      `datastore:"created" json:"created"`
	Cursor      string         `datastore:"cursor"`

	Model
}

// NewTask instantiate Task
func NewTask() *Task {
	t := &Task{
		Created: time.Now(),
	}
	t.Kind = "Task"
	return t
}
