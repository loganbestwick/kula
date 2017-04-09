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
)

type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
}

func getDBConfig() DBConfig {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := &DBConfig{}
	err := decoder.Decode(config)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return config
}

func setupDB(config *DBConfig) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.DBName)
}

func main() {
	config := getDBConfig()
	setupDB(config)
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/v1/:workout"), api.Hello)

	http.ListenAndServe("localhost:8000", mux)
}