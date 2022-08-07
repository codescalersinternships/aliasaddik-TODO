package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aliasaddik/todo-project/models"
)

func TestGetTask(t *testing.T) {

	t.Run("nothing to return", func(t *testing.T) {

		request, err := http.NewRequest("GET", "/", nil)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		actualResponse := httptest.NewRecorder()

		server.ServeHTTP(actualResponse, request)

		gotCode := actualResponse.Code
		wantCode := 502
		if gotCode != wantCode {
			t.Errorf("expected %v but got %v", wantCode, gotCode)
		}

		gotBody := actualResponse.Body.String()
		wantBody := "documents not found"

		if gotBody != wantBody {
			t.Errorf("expected %v but got %v", wantBody, gotBody[1:len(gotBody)-1])
		}

	})

	t.Run("return all entries", func(t *testing.T) {

		t1 := models.Task{Title: "Task1", Done: false}
		tasks := []string{}
		taskc.InsertOne(ctx, &t1)
		j1, _ := json.Marshal(t1)

		t2 := models.Task{Title: "Task2", Done: true}
		taskc.InsertOne(ctx, &t2)
		j2, err := json.Marshal(t2)

		tasks = append(tasks, string(j1)+","+string(j2))

		request, err := http.NewRequest("GET", "/", nil)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		actualResponse := httptest.NewRecorder()

		server.ServeHTTP(actualResponse, request)

		gotCode := actualResponse.Code
		wantCode := 200
		if gotCode != wantCode {
			t.Errorf("expected %v but got %v", wantCode, gotCode)
		}

		gotBody := actualResponse.Body.String()
		wantBody := strings.Join(tasks, " ")

		if gotBody != wantBody {
			t.Errorf("expected %v but got %v", wantBody, gotBody[1:len(gotBody)-1])
		}
	})

}
func TestPostTask(t *testing.T) {
	requestBody := []byte(`{"done": false,"title": "ll task"} `)
	request, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}

	actualResponse := httptest.NewRecorder()

	server.ServeHTTP(actualResponse, request)

	gotCode := actualResponse.Code
	wantCode := 201
	if gotCode != wantCode {
		t.Errorf("expected %v but got %v", wantCode, gotCode)
	}

}

func TestEditTask(t *testing.T) {
	requestBody := []byte(`{"done": true,"title": "edited task", "_id": "62e657f2e98f72e7d54edf78"} `)
	request, err := http.NewRequest("PUT", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}

	actualResponse := httptest.NewRecorder()

	server.ServeHTTP(actualResponse, request)

	gotCode := actualResponse.Code
	wantCode := 202
	if gotCode != wantCode {
		t.Errorf("expected %v but got %v", wantCode, gotCode)
	}

}
func TestDeleteTask(t *testing.T) {

	request, err := http.NewRequest("DELETE", "/62e4898b83762e64f0215f04", bytes.NewBuffer(nil))
	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}

	actualResponse := httptest.NewRecorder()

	server.ServeHTTP(actualResponse, request)

	gotCode := actualResponse.Code
	wantCode := 204
	if gotCode != wantCode {
		t.Errorf("expected %v but got %v", wantCode, gotCode)
	}

}
