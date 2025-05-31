package main

import (
	"encoding/json"
	"maps"
	"net/http"
)

func (app *application) respondWithJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// MarshalIndent is very expensive compared to normal Marshal
	// need to be used with care
	payload, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')

	// copy all the header to the http.ResponseWriter header map
	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)

	return nil
}
