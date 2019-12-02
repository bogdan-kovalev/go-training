package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTaskHandlerPositive(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]string{
		"start": "url1",
		"end":   "url2",
	})

	req, _ := http.NewRequest("POST", "/createTask", bytes.NewBuffer(requestBody))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}
}

func TestCreateTaskHandlerNegative(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]string{})

	req, _ := http.NewRequest("POST", "/createTask", bytes.NewBuffer(requestBody))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}
}

func TestGetTaskHandlerPositive(t *testing.T) {
	// TODO: mock task with id 1

	req, _ := http.NewRequest("GET", "/getTask/1", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetResultHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedBody := TestResponse{Start: "url1", End: "url2", Steps: []string{"url1", "url2"}}

	expectedString, _ := json.Marshal(expectedBody)

	if rr.Body.String() != string(expectedString) {
		t.Errorf("handler returned wrong response body: got %v want %v", rr.Body.String(), string(expectedString))
	}
}

type TestResponse struct {
	Start string   `json:"Start"`
	End   string   `json:"End"`
	Steps []string `json:"Steps"`
}
