package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	//REST API routes

	mux.HandleFunc("GET /api/workouts", app.viewExercise)
	return mux
}
