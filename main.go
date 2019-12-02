package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Task struct {
	Start string
	End   string
}

type TaskResult struct {
	Start string
	End   string
	Steps []string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createTask", CreateTaskHandler)
	r.HandleFunc("/getResult/{taskId}", GetResultHandler)

	log.Println("Starting server on :8000...")
	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)
}

func CreateTaskHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var task Task

	dumpRequest(request)

	handleRequest(responseWriter, request, &task)

	if task.End != "" && task.Start != "" {
		http.Redirect(responseWriter, request, "/GetResultHandler/1", http.StatusSeeOther)
	} else {
		http.Error(responseWriter, "\"start\" and/or \"end\" URLs not provided", http.StatusUnprocessableEntity)
	}
}

func GetResultHandler(writer http.ResponseWriter, request *http.Request) {
	dumpRequest(request)
	var taskResult TaskResult
	taskResult.Start = "url1"
	taskResult.End = "url2"
	taskResult.Steps = append(taskResult.Steps, taskResult.Start, taskResult.End)
	responseBody, _ := json.Marshal(taskResult)
	_, _ = fmt.Fprintf(writer, string(responseBody))
}
