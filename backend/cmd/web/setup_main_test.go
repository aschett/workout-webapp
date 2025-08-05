package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/aschett/workout-webapp/internal/models"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
)

var testModel *models.WorkoutModel
var testDB *sql.DB

func TestMain(m *testing.M) {
	code := runTests(m)
	os.Exit(code)
}

func runTests(m *testing.M) int {
	ctx := context.Background()

	mariadbContainer, err := mariadb.Run(ctx,
		"mariadb:11.0.3",
		mariadb.WithScripts(filepath.Join("..", "..", "testdata", "seed.sql")),
		mariadb.WithDatabase("testing"),
		mariadb.WithUsername("foo"),
		mariadb.WithPassword("bar"),
	)
	if err != nil {
		log.Fatalf("failed to start testMariaDB-container: %s", err)
	}
	defer func() {
		if err := mariadbContainer.Terminate(ctx); err != nil {
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
	testModel = &models.WorkoutModel{DB: db}

	return m.Run()
}
