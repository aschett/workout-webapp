package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/aschett/workout-webapp/internal/models"
	"golang.org/x/term"

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
	user := flag.String("dbuser", "web", "MySQL User")
	dbname := flag.String("dbname", "testworkoutdb", "MySQL Database Name")

	flag.Parse()

	fmt.Printf("Enter your MySQL password for User %s for Database %s:", *user, *dbname)
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	password := string(passwordBytes)

	//Building the dsn dynamically
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", *user, password, *dbname)

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
