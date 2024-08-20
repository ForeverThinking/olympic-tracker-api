package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Env struct {
	db *sql.DB
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{db: db}

	http.HandleFunc("/", env.countriesIndex)
	http.ListenAndServe(":8080", nil)
}

func (env *Env) countriesIndex(w http.ResponseWriter, r *http.Request) {
	countries, err := allCountries(env.db)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

    json.NewEncoder(w).Encode(countries)
}
