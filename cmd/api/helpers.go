package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

// This gets the ID parameter from an http request
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// This takes the destination http.ResponseWriter, the HTTP status code to send,
// the data to encode to JSON, and a header map containing any additional HTTP
// headers we want to include in the response.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, val := range headers {
		w.Header()[key] = val
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
