package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func main() {
	// open connection to psql db
	connDB, err := sql.Open("postgres", "postgresql://db:db@"+os.Getenv("DB_URL")+"/api?sslmode=disable")
	if err != nil {
		log.Println(err)
		return
	}
	db := DB{connDB}
	defer connDB.Close()

	// create api router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// health route
	r.Get("/healthz", db.checkDB)

	http.ListenAndServe(":3000", r)
}

// Check the DB connection by making a sql call
func (db DB) checkDB(w http.ResponseWriter, r *http.Request) {
	err := db.conn.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Database is up")
}
