package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DATABASE"))
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
