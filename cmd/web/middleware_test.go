package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)
	switch v := h.(type) {
	case http.Handler:
		// do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)
	switch v := h.(type) {
	case http.Handler:
		// do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestWriteToConsole(t *testing.T) {

	called := false
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	tests := []struct {
		name    string
		handler http.Handler
	}{
		{
			name:    "Test 1",
			handler: mockHandler,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrappedHandler := WriteToConsole(tt.handler)

			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			wrappedHandler.ServeHTTP(rec, req)

			if !called {
				t.Error("The handler function was not called")
			}

		})
	}
}
