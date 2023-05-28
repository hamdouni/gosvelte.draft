/*
Pour tester l'api web, on veut vérifier que quoi que la couche métier envoi
c'est bien ce que l'on récupère au bout au format json.
Donc, on va utiliser un business complètement bidon qui renvoit toujours la même
chaine de caractères "Fake Biz" (cf web_test.go).
*/
package api_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"wtk/biz"
	"wtk/biz/user"
	"wtk/ext/secure"
	"wtk/ext/store/ram"
	"wtk/ui/api"
)

func init() {
	fakeStore, _ := ram.New()
	fakeSecure, _ := secure.New()
	biz.Initialize(&fakeStore, fakeSecure)
	user.Add("FakeRealm", "samething", "samething", 1)
	api.Routes(".")
}

// TestPassword s'assure qu'un mauvais mot de passe ou un mauvais realm ne
// permet pas de se connecter
func TestPassword(t *testing.T) {
	tt := []struct {
		name     string
		realm    string
		username string
		password string
		status   int
	}{
		{"Bad Password", "FakeRealm", "toto", "titi", http.StatusUnauthorized},
		{"Good Password wrong realm", "WrongRealm", "samething", "samething", http.StatusUnauthorized},
		{"Good Password", "FakeRealm", "samething", "samething", http.StatusOK},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			payload := fmt.Sprintf(`{"username":"%s","password":"%s"}`, tc.username, tc.password)
			req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
			if err != nil {
				t.Errorf("Should be able to create a request but got %s", err)
			}
			req.Host = tc.realm + ".localhost"
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != tc.status {
				t.Errorf("Should receive %d but got %d", tc.status, rw.Code)
			}
		})
	}
}

// TestAuthEndpointUnauthorized vérifie que les endpoints qui nécessitent une
// authentification renvoient bien StatusUnauthorized
func TestAuthEndpointUnauthorized(t *testing.T) {
	tt := []struct {
		name     string
		endpoint string
	}{
		{"hello", "/hello"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %s", err)
			}
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != http.StatusUnauthorized {
				t.Errorf("Should not be authorized with code %d but got %d", http.StatusUnauthorized, rw.Code)
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
		{"hello", "/hello"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %s", err)
			}
			fakeCookie := base64.StdEncoding.EncodeToString([]byte("badtoken|badip|2006-01-02T15:04:05Z"))
			req.AddCookie(&http.Cookie{Name: "jeton", Value: fakeCookie})
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != http.StatusUnauthorized {
				t.Errorf("when %s should not be authorized with code %d but got %d", tc.endpoint, http.StatusUnauthorized, rw.Code)
			}
		})
	}
}

// TestAuthEndpoint compare le résultat du endpoint avec l'attendu.
// Sachant que l'attendu peut être une string ou une slice de string, on utilise
// une interface vide et on fait du type assertion pour faire les comparaisons
// nécessaires.
// @NOTE: le "name" du test est utilisé comme paramètre du service
func TestAuthEndpoint(t *testing.T) {
	tt := []struct {
		name     string
		endpoint string
		results  interface{}
	}{
		{"hello", "/hello", "Bonjour hello depuis le métier !"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			body := url.Values{"nom": []string{tc.name}}
			req, err := http.NewRequest("POST", tc.endpoint, bytes.NewBuffer([]byte(body.Encode())))
			if err != nil {
				t.Errorf("Should be able to create a request but got %s", err)
			}

			// forge une requête avec un jeton valide
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.RemoteAddr = "1.2.3.4"
			token, err := user.Token("fakejeton", req.RemoteAddr)
			if err != nil {
				t.Errorf("Token should not err but got %s", err)
			}
			fakeCookie := base64.StdEncoding.EncodeToString([]byte(token))
			req.AddCookie(&http.Cookie{Name: "jeton", Value: fakeCookie})

			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code == http.StatusUnauthorized {
				t.Errorf("Should be authorized but got %s", rw.Result().Status)
			}
			resp := tc.results
			if err := json.NewDecoder(rw.Body).Decode(&resp); err != nil {
				t.Errorf("Should decode the response but got %s", err)
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

func TestUnauthResponses(t *testing.T) {
	tt := []struct {
		name     string
		endpoint string
		status   int
		result   string
	}{
		{"login", "/login", http.StatusMethodNotAllowed, "method not allowed"},
		{"logout", "/logout", http.StatusFound, "redirected"},
		{"hello", "/hello", http.StatusUnauthorized, "not authorized"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %s", err)
			}
			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != tc.status {
				t.Errorf("Should receive %d but got %d", tc.status, rw.Code)
			}
			body := rw.Body
			resp := ""
			if err := json.NewDecoder(body).Decode(&resp); err != nil {
				t.Errorf("Should decode the response but got %s", err)
			}
			if resp != tc.result {
				t.Errorf("Should respond %s but got %s", tc.result, resp)
			}
		})
	}
}

func TestBadRequest(t *testing.T) {

	tt := []struct {
		name     string
		endpoint string
	}{
		{"login", "/login"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// format a failing request with empty form body and method POST
			req, err := http.NewRequest("POST", tc.endpoint, nil)
			if err != nil {
				t.Errorf("Should be able to create a request but got %s", err)
			}

			rw := httptest.NewRecorder()
			http.DefaultServeMux.ServeHTTP(rw, req)
			if rw.Code != http.StatusBadRequest {
				t.Errorf("Should receive %d but got %d", http.StatusBadRequest, rw.Code)
			}
		})
	}
}
