package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Feature struct {
	FeatureName string `json:"feature_name"`
	IsEnabled   bool   `json:"is_enabled"`
}

const (
	host     = "10.0.1.11"
	port     = 5432
	user     = "webauditor"
	password = "webauditor@1234"
	dbname   = "webauditor"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Successful.......")
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM feature")
	if err != nil {
		log.Fatal(err)
	}

	var features []Feature
	for rows.Next() {
		var feature Feature
		rows.Scan(&feature.FeatureName, &feature.IsEnabled)
		features = append(features, feature)
	}

	for _, feature := range features {
		t := template.Must(template.ParseFiles("./basic.html"))
		if err := t.Execute(w, feature); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	defer rows.Close()
	defer db.Close()
}

func main() {
	http.HandleFunc("/", GETHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
