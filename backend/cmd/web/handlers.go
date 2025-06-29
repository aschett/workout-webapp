package main

import (
	"net/http"
)

func (app *application) ViewAllWorkouts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	workouts, err := app.workouts.GetAllWorkouts()
	if err != nil {
		app.serverError(w, r, err)
	}
	app.writeJSON(w, http.StatusOK, workouts)
}
