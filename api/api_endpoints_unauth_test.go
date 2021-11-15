package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnauthResponses(t *testing.T) {
	tt := []struct {
		name     string
		endpoint string
		status   int
		result   string
	}{
		{"hello", "/hello", http.StatusOK, "Bonjour  depuis le business !"},
		{"login", "/login", http.StatusMethodNotAllowed, "method not allowed"},
		{"logout", "/logout", http.StatusFound, "redirected"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %v", err)
			}
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != tc.status {
				t.Errorf("Should receive \"%v\" but got %v", tc.status, rw.Code)
			}
			body := rw.Body
			resp := ""
			if err := json.NewDecoder(body).Decode(&resp); err != nil {
				t.Errorf("Should decode the response but got %v", err)
			}
			if resp != tc.result {
				t.Errorf("Should respond %s but got %s", tc.result, resp)
			}
		})
	}
}

func TestUnauthBadRequest(t *testing.T) {

	tt := []struct {
		name     string
		endpoint string
	}{
		{"hello", "/hello"},
		{"login", "/login"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// format a failing request with empty form body and method POST
			req, err := http.NewRequest("POST", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %v", err)
			}

			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != http.StatusBadRequest {
				t.Errorf("Should receive \"%v\" but got %v", http.StatusBadRequest, rw.Code)
			}
		})
	}
}
