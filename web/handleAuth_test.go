/*
Pour tester l'api web, on veut vérifier que quoi que la couche métier envoi
c'est bien ce que l'on récupère au bout au format json.
Donc, on va utiliser un business complètement bidon qui renvoit toujours la même
chaine de caractères "Fake Biz" (cf web_test.go).
*/
package web_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

// TestAuthEndpointUnauthorized vérifie que les endpoints qui nécessitent une
// authentification renvoient bien StatusUnauthorized
func TestAuthEndpointUnauthorized(t *testing.T) {
	tt := []struct {
		name     string
		endpoint string
	}{
		{"lower", "/lower"},
		{"upper", "/upper"},
		{"historic", "/historic"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %v", err)
			}
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != http.StatusUnauthorized {
				t.Errorf("Should not be authorized with code \"%v\" but got %v", http.StatusUnauthorized, rw.Code)
			}
		})
	}
}

// TestAuthEndpointBadToken vérifie qu'on rejette si on tente d'utiliser
// un service avec un faux jeton
func TestAuthEndpointBadToken(t *testing.T) {
	tt := []struct {
		name     string
		endpoint string
	}{
		{"lower", "/lower"},
		{"upper", "/upper"},
		{"historic", "/historic"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %v", err)
			}
			fakeCookie := base64.StdEncoding.EncodeToString([]byte("badtoken|badip|badtimestamp"))
			req.AddCookie(&http.Cookie{Name: "jeton", Value: fakeCookie})
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != http.StatusUnauthorized {
				t.Errorf("when %v should not be authorized with code \"%v\" but got %v", tc.endpoint, http.StatusUnauthorized, rw.Code)
			}
		})
	}
}

// TestAuthEndpoint compare le résultat du endpoint avec l'attendu.
// Sachant que l'attendu peut être une string ou une slice de string, on utilise
// une interface vide et on fait du type assertion pour faire les comparaisons
// nécessaires.
func TestAuthEndpoint(t *testing.T) {
	tt := []struct {
		name     string
		endpoint string
		results  interface{}
	}{
		{"lower", "/lower", "lower"},
		{"upper", "/upper", "UPPER"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			body := url.Values{"nom": []string{tc.name}}
			req, err := http.NewRequest("POST", tc.endpoint, bytes.NewBuffer([]byte(body.Encode())))
			if err != nil {
				t.Errorf("Should be able to create a request but got %v", err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.RemoteAddr = "1.2.3.4"
			timestamp := time.Now().Format(time.RFC3339)
			fakeCookie := base64.StdEncoding.EncodeToString([]byte("fake jeton|" + req.RemoteAddr + "|" + timestamp))
			req.AddCookie(&http.Cookie{Name: "jeton", Value: fakeCookie})
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code == http.StatusUnauthorized {
				t.Errorf("Should be authorized but got %v", rw.Code)
			}
			resp := tc.results
			if err := json.NewDecoder(rw.Body).Decode(&resp); err != nil {
				t.Errorf("Should decode the response but got %v", err)
			}
			different := false
			switch resp.(type) {
			case string:
				different = resp != tc.results
			case []string:
				different = resp.([]string)[0] != tc.results.([]string)[0]
			}
			if different {
				t.Errorf("Should respond %t but got %t", tc.results, resp)
			}
		})
	}
}

// Test le bad request pour les endponts /lower et /upper.
// Le endpoint /historique n'a pas de form input.
// Remarque : une bad request peut être juste une method POST et un body vide.
func TestAuthEndpointBadRequest(t *testing.T) {
	tt := []struct {
		name     string
		endpoint string
	}{
		{"lower", "/lower"},
		{"upper", "/upper"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %v", err)
			}
			req.RemoteAddr = "1.2.3.4"
			timestamp := time.Now().Format(time.RFC3339)
			fakeCookie := base64.StdEncoding.EncodeToString([]byte("fake jeton|" + req.RemoteAddr + "|" + timestamp))
			req.AddCookie(&http.Cookie{Name: "jeton", Value: fakeCookie})
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != http.StatusBadRequest {
				t.Errorf("Should receive \"%v\" but got %v", http.StatusBadRequest, rw.Code)
			}
		})
	}
}

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
