package main

import (
	"net/http"
	"strconv"
)

func (app *application) ViewAllWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := app.workouts.GetAllWorkouts()
	if err != nil {
		app.serverError(w, r, err)
	}
	app.writeJSON(w, http.StatusOK, workouts)
}

func (app *application) ViewWorkoutEntries(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	exercises, err := app.workouts.GetWorkoutEntries(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusOK, exercises)
}
