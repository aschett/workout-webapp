package models

import (
	"database/sql"
)

type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Workout struct {
	ID       int            `json:"id"`
	Date     string         `json:"date"`
	Workouts []WorkoutEntry `json:"workouts"`
}

type WorkoutEntry struct {
	ID           int     `json:"id"`
	WorkoutID    int     `json:"workoutID"`
	ExerciseID   int     `json:"exerciseID"`
	ExerciseName string  `json:"exerciseName"`
	Weight       float64 `json:"weight"`
	Sets         int     `json:"sets"`
	Reps         int     `json:"reps"`
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

func (m *WorkoutModel) GetAllWorkouts() ([]Workout, error) {
	stmt := `SELECT id, date FROM workouts`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []Workout

	for rows.Next() {
		var w Workout
		err := rows.Scan(&w.ID, &w.Date)
		if err != nil {
			return nil, err
		}

		entries, err := m.GetWorkoutEntries(w.ID)
		if err != nil {
			return nil, err
		}
		w.Workouts = entries

		workouts = append(workouts, w)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workouts, nil
}

func (m *WorkoutModel) GetWorkoutEntries(workoutID int) ([]WorkoutEntry, error) {
	stmt := `SELECT we.id, we.workoutID, we.exerciseID, e.name AS exerciseName, we.weight, we.sets, we.reps
             FROM workoutEntries we JOIN exercises e on we.exerciseID = e.id
             WHERE we.workoutID = ?`

	rows, err := m.DB.Query(stmt, workoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []WorkoutEntry

	for rows.Next() {
		var e WorkoutEntry
		err := rows.Scan(
			&e.ID,
			&e.WorkoutID,
			&e.ExerciseID,
			&e.ExerciseName,
			&e.Weight,
			&e.Sets,
			&e.Reps,
		)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
