package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.healthCheckHandler)

	apiHandler := http.StripPrefix("/api/v1", mux)

	return apiHandler
}
