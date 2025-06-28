package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) viewExercise(w http.ResponseWriter, r *http.Request) {
	workouts := []map[string]interface{}{
		{"id": 1, "title": "Chest Press", "weight": 32.5},
		{"id": 2, "title": "Leg Press", "weight": 100},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(workouts)

	if err != nil {
		app.serverError(w, r, err)
	}
}
