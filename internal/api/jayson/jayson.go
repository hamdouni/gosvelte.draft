package jayson

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// Respond retourne une reponse json avec le statut et le contenu passés en paramètre
func Respond(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write([]byte(response)); err != nil {
		log.Printf("respondJSON error %s", err)
	}
}

var ErrJsonPayloadEmpty = errors.New("empty Json Payload")

func DecodeJson(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return ErrJsonPayloadEmpty
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if len(content) == 0 {
		return ErrJsonPayloadEmpty
	}
	err = json.Unmarshal(content, v)
	if err != nil {
		return err
	}
	return nil
}
