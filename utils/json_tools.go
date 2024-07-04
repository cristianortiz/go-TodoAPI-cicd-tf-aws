package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// JsonTools is the type for this package. Create a variable of this type and you have access
// to all the methods with the receiver type *JsonTools.
type JsonTools struct {
}

// JSONResponse is the type used for sending JSON around
type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ReadJSON tries to read the body of a request and converts it into JSON
func (t *JsonTools) ReadJSON(r *http.Request, data any) error {

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}

	return nil
}

// WriteJSON takes a response status code and arbitrary data and writes a json response to the client
func (t *JsonTools) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {

	out := json.NewEncoder(w)
	//check if header was included as the last parameter
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value

		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := out.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

// ErrorJSON takes an error, and optionally a response status code, and generates and sends
// a json error response
func (t *JsonTools) ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return t.WriteJSON(w, statusCode, payload)
}
