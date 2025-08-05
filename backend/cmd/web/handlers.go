package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

func (app *application) ViewAllWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := app.workouts.GetAllWorkouts()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusOK, workouts)
}

func (app *application) ViewWorkoutByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	exercises, err := app.workouts.GetWorkoutByID(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.writeJSON(w, http.StatusOK, exercises)
}

/*This function sends inserts the date and returns statuscode 409 if the selected date already exists.*/
func (app *application) AddNewWorkout(w http.ResponseWriter, r *http.Request) {
	var mySqlError *mysql.MySQLError
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	date := (r.PostForm.Get("date"))
	id, err := app.workouts.AddWorkout(date)
	if err != nil {
		if errors.As(err, &mySqlError) && mySqlError.Number == 1062 {
			app.clientError(w, http.StatusConflict)
			return
		}
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/api/workouts/%d", id), http.StatusSeeOther)
}
