package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Task struct {
	Start string
	End   string
}

type CompletedTask struct {
	Start string
	End   string
	Steps []string
}

const TaskId = "taskId"

func main() {
	completedTasks := make(map[string]CompletedTask)

	r := mux.NewRouter()
	r.HandleFunc("/createTask", CreateTaskHandler(completedTasks))
	r.HandleFunc("/getResult/{"+TaskId+"}", GetResultHandler(completedTasks))

	log.Println("Starting server on :8000...")
	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)
}

func CreateTaskHandler(completedTasks map[string]CompletedTask) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var task Task

		handleRequest(responseWriter, request, &task)

		if task.End != "" && task.Start != "" {
			var taskResult CompletedTask
			taskResult.Start = task.Start
			taskResult.End = task.End
			taskResult.Steps = append(taskResult.Steps, task.Start, task.End)
			id := fmt.Sprint(time.Now().UnixNano())
			completedTasks[id] = taskResult

			http.Redirect(responseWriter, request, "/getResult/"+id, http.StatusSeeOther)
		} else {
			http.Error(responseWriter, "\"start\" and/or \"end\" URLs not provided", http.StatusUnprocessableEntity)
		}
	}
}

func GetResultHandler(completedTasks map[string]CompletedTask) func(writer http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		completedTask, ok := completedTasks[params[TaskId]]
		if ok {
			responseBody, _ := json.Marshal(completedTask)
			responseWriter.Header().Set("Content-Type", "application/json")
			_, _ = responseWriter.Write(responseBody)
		} else {
			responseWriter.WriteHeader(http.StatusNotFound)
		}
	}
}
