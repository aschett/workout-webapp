package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/workouts", app.AddNewWorkout)
	mux.HandleFunc("GET /api/workouts", app.ViewAllWorkouts)
	mux.HandleFunc("GET /api/workouts/{id}", app.ViewWorkoutByID)
	return mux
}
