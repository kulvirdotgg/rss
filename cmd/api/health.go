package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	payload := JSON{
		"message": "uWu oniii chan!!!!!",
		"info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.respondWithJSON(w, http.StatusOK, payload, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
