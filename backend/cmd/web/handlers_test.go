package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aschett/workout-webapp/internal/models"
)

func TestViewAllWorkouts(t *testing.T) {
	var app = &application{
		workouts: testModel,
	}
	req := httptest.NewRequest(http.MethodGet, "/api/workouts", nil)
	rr := httptest.NewRecorder()

	app.ViewAllWorkouts(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, received %d", rr.Code)
	}

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Failed to read rsponse body: %v", err)
	}

	var workouts []models.Workout
	if err := json.Unmarshal(body, &workouts); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if len(workouts) != 10 {
		t.Errorf("Expected 10 workouts, received %d", len(workouts))
	}

	if workouts[0].Date == "" {
		t.Errorf("Expected date for first workout, received empty")
	}

	if workouts[0].Date != "2024-06-20" {
		t.Errorf("Expected 2024-06-20, received %s", workouts[0].Date)
	}
}

func TestViewWorkoutEntries(t *testing.T) {
	var app = &application{
		workouts: testModel,
	}
	req := httptest.NewRequest(http.MethodGet, "/api/workouts/1", nil)
	rr := httptest.NewRecorder()

	app.routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, received %d", rr.Code)
	}

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var workoutEntries []models.WorkoutEntry
	if err := json.Unmarshal(body, &workoutEntries); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if len(workoutEntries) != 3 {
		t.Errorf("Expected 3 WorkoutEntries, received %d", len(workoutEntries))
	}

	if workoutEntries[0].ExerciseName == "" {
		t.Errorf("Expected name for first WorkoutEntry, received empty")
	}

	if workoutEntries[0].ExerciseName != "Bench Press" {
		t.Errorf("Expected Bench Press, received %s", workoutEntries[0].ExerciseName)
	}
}

func TestViewWorkoutEntries_InvalidID(t *testing.T) {
	app := &application{workouts: testModel}

	req := httptest.NewRequest(http.MethodGet, "/api/workouts/abc", nil)
	rr := httptest.NewRecorder()

	app.routes().ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("Expected 404 for invalid ID, received %d", rr.Code)
	}
}

func TestWorkoutRoutes(t *testing.T) {
	app := &application{workouts: testModel}

	tests := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{"All Workouts", "/api/workouts", http.StatusOK},
		{"Valid ID", "/api/workouts/1", http.StatusOK},
		{"Invalid ID", "/api/workouts/abc", http.StatusNotFound},
		{"Negative ID", "/api/workouts/-1", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rr := httptest.NewRecorder()

			app.routes().ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected %d for %s, received %d", tt.expectedStatus, tt.url, rr.Code)
			}
		})
	}
}
