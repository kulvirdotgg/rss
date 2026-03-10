package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576 // 1 MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes)) // limiting the size of request body

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // disallow unknown fields in the request body

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}
