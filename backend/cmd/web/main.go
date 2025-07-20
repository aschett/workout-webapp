package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/aschett/workout-webapp/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	workouts *models.WorkoutModel
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	addr := flag.String("addr", ":4000", "Http network address")

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	if host == "" || user == "" || dbname == "" || password == "" {
		logger.Error("One or more environment variables are not set")
		os.Exit(1)
	}
	// Build the dsn dynamically
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbname)

	db, err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	logger.Info("Starting Server on", "addr", *addr)

	app := &application{
		logger:   logger,
		workouts: &models.WorkoutModel{DB: db},
	}

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
