package main

import (
	"fmt"
	"net/http"
	"rss/internal/data"
	"rss/internal/validator"
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

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name: input.Name,
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.validationFailedResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}
