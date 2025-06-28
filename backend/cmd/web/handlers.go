package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Workout struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Date   string  `json:"date"`
	Weight float64 `json:"weight"`
}

func (app *application) viewExercise(w http.ResponseWriter, r *http.Request) {

	currentTime := time.Now().UTC().Format("2006-01-02")

	workouts := []Workout{
		{1, "Chest Press", currentTime, 32.5},
		{2, "Leg Press", currentTime, 100},
		{2, "Leg Press", currentTime, 110},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(workouts)

	if err != nil {
		app.serverError(w, r, err)
	}
}
