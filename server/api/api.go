package api

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"goji.io/pat"
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
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	id := 0
	err = a.db.QueryRow(sqlStatement, wr.Name, wr.Series, wr.Category, wr.SubCategory, wr.Duration).Scan(&id)
	if err != nil {
		panic(err)
	}
}

func (a *API) GetWorkout(w http.ResponseWriter, r *http.Request) {
	var wr Workout
	id := pat.Param(r, "id")
	sqlStatement := `SELECT * FROM workouts WHERE id=$1;`
	row := a.db.QueryRow(sqlStatement, id)
	err := row.Scan(&wr.Name, &wr.Series, &wr.Category, &wr.SubCategory, &wr.Duration, &wr.ID)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(wr)
	default:
		panic(err)
	}
	jsonOut, _ := json.Marshal(wr)
	fmt.Fprintf(w, string(jsonOut))
}

func (a *API) Setup(db *sql.DB) {
	a.db = db
}
