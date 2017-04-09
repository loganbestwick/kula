package api

import (
	"net/http"

	"database/sql"
	"encoding/json"
	"io/ioutil"
)

type API struct {
	db *sql.DB
}

func (a *API) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var wr Workout
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(jsonData), &wr)
	if err != nil {
		panic(err)
	}
	sqlStatement := `
		INSERT INTO workouts (name, series, category, subcategory, duration)
		VALUES ($1, $2, $3, $4, $5)`
	_, err = a.db.Exec(sqlStatement, wr.Name, wr.Series, wr.Category, wr.SubCategory, wr.Duration)
	if err != nil {
		panic(err)
	}
}

func (a *API) Setup(db *sql.DB) {
	a.db = db
}