package main

import (
	"net/http"
	"rss/internal/data"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.getIntParam(r, "id")
	if err != nil {
		app.notFoundResponse(w, r, err)
		return
	}

	user := data.User{
		ID:   id,
		Name: "random Guy",
	}

	err = app.respondWithJSON(w, http.StatusOK, JSON{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
