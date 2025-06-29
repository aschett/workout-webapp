package models

import (
	"database/sql"
)

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Workout struct {
	ID     int    `json:"id"`
	UserID int    `json:"userID"`
	Date   string `json:"date"`
}

type WorkoutEntry struct {
	ID         int     `json:"id"`
	UserID     int     `json:"userID"`
	WorkoutID  int     `json:"workoutID"`
	ExerciseID int     `json:"exerciseID"`
	Weight     float64 `json:"weight"`
	Sets       int     `json:"sets"`
	Reps       int     `json:"reps"`
}

type WorkoutModel struct {
	DB *sql.DB
}

func (m *WorkoutModel) AddWorkout(date string) (int, error) {
	stmt := `INSERT INTO workouts (date) VALUES (?)`

	result, err := m.DB.Exec(stmt, date)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *WorkoutModel) AddWorkOutEntry(workoutID int, exerciseID int, weight float64, sets int, reps int) (int, error) {
	stmt := `INSERT INTO workoutEntries (workoutID, exerciseID, weight, sets, reps)
	VALUES (?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, workoutID, exerciseID, weight, sets, reps)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
