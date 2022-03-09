package api_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestBadPassword(t *testing.T) {
	form := url.Values{}
	form.Add("username", "toto")
	form.Add("password", "titi")
	body := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Errorf("Should be able to create a request but got %s", err)
	}
	req.Form = form
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)
	if rw.Code != http.StatusUnauthorized {
		t.Errorf("Should receive \"%v\" but got %v", http.StatusUnauthorized, rw.Code)
	}
}
func TestGoodPassword(t *testing.T) {
	form := url.Values{}
	form.Add("username", "samething")
	form.Add("password", "samething")
	body := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Errorf("Should be able to create a request but got %s", err)
	}
	req.Form = form
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)
	if rw.Code != http.StatusOK {
		t.Errorf("Should receive \"%v\" but got %v", http.StatusOK, rw.Code)
	}
}
