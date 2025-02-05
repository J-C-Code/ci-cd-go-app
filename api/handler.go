package api

import (
	"encoding/json"
	"net/http"
)

var tasks []Task = []Task{
	{ID: 1, Title: "First task!", Description: "This is the first task, this is awesome!", Done: true},
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
