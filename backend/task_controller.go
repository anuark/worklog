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

	_ = orderStr
	_ = limit
	_ = offset

	user, _ := UserFromContext(r.Context())
	q := datastore.NewQuery("Task").
		Ancestor(user.Key).
		Order(orderStr).
		Limit(limit).
		Offset(offset)

	var tasks []Task
	for t := dsClient.Run(r.Context(), q); ; {
		task := Task{}
		k, err := t.Next(&task)
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
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
		panic(err)
	}

	user, _ := UserFromContext(r.Context())
	task.AncestorKey = user.Key

	task.Save(task)
	json, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(json))
}

// TaskUpdate .
func TaskUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["taskId"])
	t, err := findTask(int64(id), w, r)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(t)
	if err != nil {
		panic(err)
	}

	json, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	t.Save(t)

	fmt.Fprintln(w, string(json))
}

// TaskDelete .
func TaskDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["taskId"])
	t, err := findTask(int64(id), w, r)
	if err != nil {
		panic(err)
	}
	t.Delete()
}

// TaskView .
func TaskView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["taskId"])
	t, err := findTask(int64(id), w, r)
	if err != nil {
		panic(err)
	}

	json, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(json))
}

func findTask(id int64, w http.ResponseWriter, r *http.Request) (*Task, error) {
	t := NewTask()
	u, _ := UserFromContext(r.Context())
	k := datastore.IDKey(t.Kind, id, u.Key)
	if err := dsClient.Get(dsCtx, k, t); err != nil {
		panic(err)
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
