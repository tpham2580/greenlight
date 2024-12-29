package main

import (
	"fmt"
	"net/http"
)

// generic helper for logging an error message
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// generic helper for sending JSON-formatted error messages to the client
// with a given status code
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// sends a 500 Internal Server Error status code and JSON response to the client
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encoutered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// sends a 400 Bad Request status code and JSON response to client
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// sends a 404 Not Found status code and JSON response to client
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// sends a 405 Method Not Allowed status code and JSON respone to client
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
