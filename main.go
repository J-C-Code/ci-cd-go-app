package main

import (
	"log"
	"net/http"

	"github.com/Jomma52637/ci-cd-go-app/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", api.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", api.CreateTask).Methods("POST")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
