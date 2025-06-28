package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *application) viewExercise(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	workouts := []map[string]interface{}{
		{"id": 1, "title": "Chest Press", "weight": 32.5},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(workouts)

	if err != nil {
		app.serverError(w, r, err)
	}
}
