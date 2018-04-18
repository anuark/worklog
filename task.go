package main

import (
	"time"
)

// Task struct
type Task struct {
	Description string    `datastore:"description" json:"description"`
	Created     time.Time `datastore:"created" json:"created"`

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
