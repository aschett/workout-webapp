package models

import (
	"database/sql"
)

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Workout struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
}

type WorkoutEntry struct {
	ID         int     `json:"id"`
	WorkoutID  int     `json:"workoutID"`
	ExerciseID int     `json:"exerciseID"`
	Weight     float64 `json:"weight"`
	Sets       int     `json:"sets"`
	Reps       int     `json:"reps"`
}

type WorkoutModel struct {
	DB *sql.DB
}
