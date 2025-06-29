package main

import (
	"net/http"
)

func (app *application) ViewAllWorkouts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	workouts := app.
		app.writeJSON(w, http.StatusOK, workouts)

	if err != nil {
		app.serverError(w, r, err)
	}
}
