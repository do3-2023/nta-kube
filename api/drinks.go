package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Drink struct {
	Emoji string
	Name  string
}

var drinks = []Drink{
	{Emoji: "🍷", Name: "White wine"},
	{Emoji: "🍷", Name: "Red wine"},
	{Emoji: "🍺", Name: "IPA"},
	{Emoji: "🍺", Name: "Pale Ale"},
	{Emoji: "🍺", Name: "Lager"},
	{Emoji: "🍺", Name: "Stout"},
	{Emoji: "🍹", Name: "Pina Colada"},
	{Emoji: "🍹", Name: "Sex on the beach"},
	{Emoji: "🍸", Name: "Martini"},
	{Emoji: "🥤", Name: "Water"},
}

func (db *DB) getDrinks(w http.ResponseWriter, r *http.Request) {
	query := `SELECT emoji, name FROM drinks`
	rows, err := db.conn.Query(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	var drinks []Drink
	for rows.Next() {
		var drink Drink
		if err := rows.Scan(&drink.Emoji, &drink.Name); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		drinks = append(drinks, drink)
	}

	// convert to json
	out, err := json.Marshal(drinks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(out))
}

func (db *DB) addDrink(w http.ResponseWriter, r *http.Request) {
	// get random drink to add
	drink := drinks[rand.Intn(len(drinks))]

	// add to db
	query := `INSERT INTO drinks (emoji, name) VALUES ($1, $2)`
	_, err := db.conn.Exec(
		query,
		drink.Emoji,
		drink.Name,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	log.Println("New drink added")

	// convert to json
	out, err := json.Marshal(drink)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(out))
}
