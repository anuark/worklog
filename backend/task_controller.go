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
		k, err := t.Next(&task)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		task.ID = k.ID

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
	task := NewTask()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if err != nil {
		fmt.Println(err)
	}

	user, _ := UserFromContext(r.Context())
	fmt.Println(user)

	task.Save(task)
	json, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(json))
}

// TaskUpdate .
func TaskUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["taskId"])
	t, err := findTask(int64(id), w)
	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(t)
	if err != nil {
		fmt.Println(err)
	}

	json, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}

	t.Save(t)

	fmt.Fprintln(w, string(json))
}

// TaskDelete .
func TaskDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["taskId"])
	t, err := findTask(int64(id), w)
	if err != nil {
		fmt.Println(err)
	}
	t.Delete()
}

// TaskView .
func TaskView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["taskId"])
	t, err := findTask(int64(id), w)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	json, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(json))
}

func findTask(id int64, w http.ResponseWriter) (*Task, error) {
	t := NewTask()
	k := datastore.IDKey(t.Kind, id, nil)
	if err := dsClient.Get(dsCtx, k, t); err != nil {
		fmt.Println(err)
	}
	t.ID = id

	var err error
	if t.Key == nil {
		http.Error(w, "No task with id: "+string(id), http.StatusNotFound)
		err = errors.New("No task with id: " + string(id))
		return t, err
	}

	return t, err
}
