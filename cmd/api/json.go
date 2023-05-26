package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var ErrJsonPayloadEmpty = errors.New("Empty Json Payload")

func decodeJson(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return ErrJsonPayloadEmpty
	}

	content, err := ioutil.ReadAll(r.Body)
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
