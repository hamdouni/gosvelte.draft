package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var ErrJsonPayloadEmpty = errors.New("Empty Json Payload")

func decodeJson(r *http.Request, v interface{}) error {
	content, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return err
	}
	if len(content) == 0 {
		return ErrJsonPayloadEmpty
	}
	err = json.Unmarshal(content, v)
	if err != nil {
		return err
	}
	return nil
}
