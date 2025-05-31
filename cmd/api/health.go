package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{
		"message":     "uWu oniii chan!!!!!",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.respondWithJSON(w, http.StatusOK, payload, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Server could not process the request", http.StatusInternalServerError)
	}
}
