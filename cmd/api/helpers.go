package main

import (
	"encoding/json"
	"errors"
	"maps"
	"net/http"
	"strconv"
)

type JSON map[string]any

func (app *application) respondWithJSON(w http.ResponseWriter, status int, payload JSON, headers http.Header) error {
	// MarshalIndent is very expensive compared to normal Marshal
	// need to be used with care
	res, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}
	res = append(res, '\n')

	// copy all the header to the http.ResponseWriter header map
	maps.Copy(w.Header(), headers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)

	return nil
}

func (app *application) getIntParam(r *http.Request, name string) (int64, error) {
	param := r.PathValue(name)
	val, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}
	return val, nil
}
