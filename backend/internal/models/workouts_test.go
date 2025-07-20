package models

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
)

var testModel *WorkoutModel
var testDB *sql.DB
var container testcontainers.Container

/*
The TestMain function in go is to ensure that every test will be ran out of this specific function and ensures that it will be ran first
Therefore ensuring that our spun up Database Container exists during our testsuite
*/

func TestMain(m *testing.M) {
	ctx := context.Background()

	mariadbContainer, err := mariadb.Run(ctx,
		"mariadb:11.0.3",
		mariadb.WithScripts(filepath.Join("testdata", "seed.sql")),
		mariadb.WithDatabase("testing"),
		mariadb.WithUsername("foo"),
		mariadb.WithPassword("bar"),
	)
	if err != nil {
		log.Fatalf("failed to start testMariaDB-container: %s", err)
	}
	container = mariadbContainer

	defer func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Printf("failed to terminate testMariaDB-container: %s", err)
		}
	}()

	dsn, err := mariadbContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("failed to get DSN: %v", err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to test DB: %v", err)
	}

	testDB = db
	testModel = &WorkoutModel{DB: db}

	code := m.Run()

	os.Exit(code)
}

// All Values are taken out of the ./testdata/seed.sql file
func TestGetWorkoutEntries(t *testing.T) {
	entries, err := testModel.GetWorkoutEntries(1)
	if err != nil {
		t.Fatalf("Functioncall GetWorkoutEntries failed: %v", err)
	}

	expected := []WorkoutEntry{
		{ExerciseID: 1, ExerciseName: "Bench Press", Weight: 100, Sets: 3, Reps: 10},
		{ExerciseID: 2, ExerciseName: "Squat", Weight: 120, Sets: 3, Reps: 8},
		{ExerciseID: 5, ExerciseName: "Barbell Row", Weight: 90, Sets: 3, Reps: 10},
	}
	if len(entries) != len(expected) {
		t.Fatalf("Expected %d entries, received %d", len(expected), len(entries))
	}

	for i, received := range entries {
		expVal := expected[i]

		if received.ExerciseID != expVal.ExerciseID {
			t.Errorf("Entry %d: expected ExerciseID %d, received %d", i, expVal.ExerciseID, received.ExerciseID)
		}
		if received.ExerciseName != expVal.ExerciseName {
			t.Errorf("Entry %d: expected ExerciseName %q, received %q", i, expVal.ExerciseName, received.ExerciseName)
		}
		if received.Weight != expVal.Weight {
			t.Errorf("Entry %d: expected Weight %.1f, received %.1f", i, expVal.Weight, received.Weight)
		}
		if received.Sets != expVal.Sets {
			t.Errorf("Entry %d: expected Sets %d, received %d", i, expVal.Sets, received.Sets)
		}
		if received.Reps != expVal.Reps {
			t.Errorf("Entry %d: expected Reps %d, received %d", i, expVal.Reps, received.Reps)
		}
	}
}

func TestGetAllWorkouts(t *testing.T) {
	workouts, err := testModel.GetAllWorkouts()
	if err != nil {
		t.Fatalf("GetAllWorkouts failed: %v", err)
	}

	expected := []Workout{
		{ID: 1, Date: "2024-06-20"},
		{ID: 2, Date: "2024-06-22"},
		{ID: 3, Date: "2024-06-24"},
		{ID: 4, Date: "2024-06-26"},
		{ID: 5, Date: "2024-06-28"},
		{ID: 6, Date: "2024-06-30"},
		{ID: 7, Date: "2024-07-02"},
		{ID: 8, Date: "2024-07-04"},
		{ID: 9, Date: "2024-07-06"},
		{ID: 10, Date: "2024-07-08"},
	}

	if len(workouts) != len(expected) {
		t.Fatalf("Expected %d workouts, received %d", len(expected), len(workouts))
	}

	for i, received := range workouts {
		want := expected[i]
		if received.ID != want.ID {
			t.Errorf("Workout %d: expected ID %d, received %d", i, want.ID, received.ID)
		}
		if received.Date != want.Date {
			t.Errorf("Workout %d: expected Date %s, received %s", i, want.Date, received.Date)
		}
	}
}

func TestAddWorkout(t *testing.T) {
	testDate := "2025-10-10"

	id, err := testModel.AddWorkout(testDate)
	if err != nil {
		t.Fatalf("AddWorkout failed: %v", err)
	}

	var receivedDate string
	err = testDB.QueryRow("SELECT date FROM workouts WHERE id = ?", id).Scan(&receivedDate)
	if err != nil {
		t.Fatalf("Workout not existent after insert: %v", err)
	}
	if receivedDate != testDate {
		t.Errorf("Expected date %s, received %s", testDate, receivedDate)
	}
}

func TestAddWorkoutEntry(t *testing.T) {
	workoutID := 1
	exerciseID := 3
	weight := 145.0
	sets := 3
	reps := 6

	id, err := testModel.AddWorkOutEntry(workoutID, exerciseID, weight, sets, reps)
	if err != nil {
		t.Fatalf("AddworkoutEntry failed: %v", err)
	}

	var receivedWorkoutID, receivedExerciseID int
	var receivedWeight float64
	var receivedSets, receivedReps int

	err = testDB.QueryRow(`
		SELECT workoutID, exerciseID, weight, sets, reps
		FROM workoutEntries
		WHERE id = ?`, id).Scan(&receivedWorkoutID, &receivedExerciseID, &receivedWeight, &receivedSets, &receivedReps)

	if err != nil {
		t.Fatalf("Inserted workout entry not found: %v", err)
	}
	if receivedWorkoutID != workoutID {
		t.Errorf("Expected workoutID %d, received %d", workoutID, receivedWorkoutID)
	}
	if receivedExerciseID != exerciseID {
		t.Errorf("Expected exerciseID %d, received %d", exerciseID, receivedExerciseID)
	}
	if receivedWeight != weight {
		t.Errorf("Expected weight %.1f, received %.1f", weight, receivedWeight)
	}
	if receivedSets != sets {
		t.Errorf("Expected sets %d, received %d", sets, receivedSets)
	}
	if receivedReps != reps {
		t.Errorf("Expected reps %d, received %d", reps, receivedReps)
	}
}
