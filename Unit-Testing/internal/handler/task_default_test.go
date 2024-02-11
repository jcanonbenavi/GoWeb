package handler_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/unit-testing/internal"
	"github.com/jcanonbenavi/unit-testing/internal/handler"
	"github.com/jcanonbenavi/unit-testing/internal/repository"
	"github.com/stretchr/testify/require"
)

// req := NewRequest("GET", "/tasks/1", nil, map[string]string{"id": "1"}, nil)
func NewRequest(method, url string, body io.Reader, urlParams map[string]string, urlQuery map[string]string) *http.Request {
	req := httptest.NewRequest(method, url, body)

	if urlParams != nil {
		chiCtx := chi.NewRouteContext()
		for key, value := range urlParams {
			chiCtx.URLParams.Add(key, value)
		}
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	}

	if urlQuery != nil {
		q := req.URL.Query()
		for key, value := range urlQuery {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}
	return req
}

func TestTaskDefault_GetTask(t *testing.T) {
	t.Run("should return the task with the given ID", func(t *testing.T) {
		// create a task
		db := map[int]internal.Task{
			1: {
				ID:          1,
				Name:        "task 1",
				Description: "description 1",
				Completed:   false,
			},
		}
		//initialize the repository with the db
		repository := repository.NewTaskMap(db)
		// create a handler
		hd := handler.NewTaskDefault(repository)
		// create a request
		handlerFunc := hd.GetTask()
		// create a response with:
		//GET METHOD
		//URL /tasks/1
		//BODY nil because is a GET method
		//URLPARAMS map[string]string{"id": "1"}
		//URLQUERY nil because is a GET method
		req := NewRequest("GET", "/tasks/1", nil, map[string]string{"id": "1"}, nil)
		// create a response recorder
		res := httptest.NewRecorder()
		// execute the handler func with the request and response
		handlerFunc(res, req)
		// check the response
		expectedCode := http.StatusOK
		expectedBody := `{"id":1,"name":"task 1","description":"description 1","completed":false}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})

	t.Run("failure 01 - task not found", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewTaskMap(nil)
		// - handler
		hd := handler.NewTaskDefault(rp)
		hdFunc := hd.GetTask()

		// act
		req := NewRequest("GET", "/tasks/1", nil, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusNotFound
		expectedBody := fmt.Sprintf(`{"status":"%s","message":"%s"}`, http.StatusText(expectedCode), "task not found")
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("failure 02 - invalid id", func(t *testing.T) {
		// arrange
		// - repository
		// ...
		// - handler
		hd := handler.NewTaskDefault(nil)
		hdFunc := hd.GetTask()

		// act
		req := httptest.NewRequest("GET", "/tasks/invalid", nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := fmt.Sprintf(`{"status":"%s","message":"%s"}`, http.StatusText(expectedCode), "invalid id")
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

}
