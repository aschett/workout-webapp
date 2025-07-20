package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/aschett/workout-webapp/internal/models"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
)

var testModel *models.WorkoutModel
var testDB *sql.DB
var container testcontainers.Container

func TestMain(m *testing.M) {
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
	testModel = &models.WorkoutModel{DB: db}

	code := m.Run()

	os.Exit(code)
}
