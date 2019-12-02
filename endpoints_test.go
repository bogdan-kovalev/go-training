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
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}
}

func TestCreateTaskHandlerNegative(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]string{})

	req, err := http.NewRequest("POST", "/createTask", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}
}
