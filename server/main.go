package main

import (
	"net/http"
	"database/sql"
	"encoding/json"

	"goji.io"
	"goji.io/pat"
	"github.com/loganbestwick/kula/server/api"
	_ "github.com/lib/pq"
	"os"
	"fmt"
	"io/ioutil"
)

type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
}

func getDBConfig() *DBConfig {
	config := &DBConfig{}
	file, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, config)
	return config
}

func setupDB()  *sql.DB {
	config := getDBConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connection Successful")
	return db
}

func main() {
	var apiInstance api.API
	db := setupDB()
	defer db.Close()
	apiInstance.Setup(db)
	mux := goji.NewMux()
	mux.HandleFunc(pat.Post("/v1/workout"), apiInstance.CreateWorkout)

	http.ListenAndServe("localhost:8000", mux)
}