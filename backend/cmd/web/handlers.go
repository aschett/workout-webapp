package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) ViewAllWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := app.workouts.GetAllWorkouts()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusOK, workouts)
}

func (app *application) ViewWorkoutByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	exercises, err := app.workouts.GetWorkoutByID(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusOK, exercises)
}

func (app *application) AddNewWorkout(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	date := (r.PostForm.Get("date"))
	id, err := app.workouts.AddWorkout(date)
	if err != nil {
		app.serverError(w, r, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/api/workouts/%d", id), http.StatusSeeOther)
}
