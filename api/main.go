package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func main() {
	// create connection to DB
	connURL := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_NAME"),
	)
	db, err := newDB(connURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.conn.Close()
	log.Println("Connected to DB")

	// create table for drinks
	query := `
		CREATE TABLE IF NOT EXISTS drinks (
			id SERIAL PRIMARY KEY,
			emoji VARCHAR(100) NOT NULL,
			name VARCHAR(100) NOT NULL
		)`
	_, err = db.conn.Exec(query)
	if err != nil {
		log.Println("The table already exists")
	}
	log.Println("Table drinks is ready")

	// create api router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// health route
	r.Get("/healthz", db.checkDB)

	r.Get("/drinks", db.getDrinks)
	r.Post("/drinks", db.addDrink)

	log.Println("Starting api on port 3000")
	http.ListenAndServe(":3000", r)
}

// Create a new DB connection
func newDB(url string) (*DB, error) {
	var conn *sql.DB
	var err error

	// try 10 times to connect to the DB
	for i := 0; i < 10; i++ {
		conn, _ = sql.Open("postgres", url)
		err = conn.Ping()
		if err == nil {
			break
		}
		log.Println("Failed to connect to DB. Retrying in 5 seconds")
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return &DB{}, fmt.Errorf("Failed to connect to DB after 10 tries")
	}

	return &DB{conn}, nil
}

// Check the DB connection by making a sql call
func (db *DB) checkDB(w http.ResponseWriter, r *http.Request) {
	err := db.conn.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Database is up")
}
