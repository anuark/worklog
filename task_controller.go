package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

// TaskList .
func TaskList(w http.ResponseWriter, r *http.Request) {
	order := r.FormValue("_order")
	end, _ := strconv.Atoi(r.FormValue("_end"))
	start, _ := strconv.Atoi(r.FormValue("_start"))
	limit := end - start
	if limit < 0 {
		limit = -limit
	}
	offset := start

	orderSign := ""
	if order == "DESC" {
		orderSign = "-"
	}
	sort := r.FormValue("_sort")

	orderStr := orderSign + sort
	if sort == "id" {
		orderStr = ""
	}

	q := datastore.NewQuery("Task").
		Order(orderStr).
		Limit(limit).
		Offset(offset)

	var tasks []Task
	for t := dsClient.Run(r.Context(), q); ; {
		var task Task
		_, err := t.Next(&task)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		tasks = append(tasks, task)
	}

	count := fmt.Sprintf("%d", len(tasks))
	w.Header().Add("X-Total-Count", count)

	j, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	if string(j) == "null" {
		j = []byte("[]")
	}

	fmt.Fprint(w, string(j))
}

// TaskCreate action for new task.
func TaskCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	task := NewTask()

	var t Task
	err := decorder.Decode(&t)
	if err != nil {
		fmt.Println(err)
	}

	if len(val) > 1 {
		task.Description = r.PostFormValue("desc")
		task.Save(task)
	}
}

// TaskUpdate .
func TaskUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["taskId"])
	t, err := findTask(int64(id), w)
	if err != nil {
		return
	}
	fmt.Println(t)
	fmt.Fprintf(w, "{\"message\": \"success\"}")
}

// TaskDelete .
func TaskDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["taskId"])
	t, err := findTask(int64(id), w)
	if err != nil {
		log.Fatal(err)
	}
	t.Delete()
}

func findTask(id int64, w http.ResponseWriter) (*Task, error) {
	t := NewTask()
	t.Get(id, t)
	var err error
	if t.Key == nil {
		http.Error(w, "No task with id: "+string(id), http.StatusNotFound)
		err = errors.New("No task with id: " + string(id))
		return t, err
	}

	return t, err
}