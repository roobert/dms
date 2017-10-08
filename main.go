package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	defer db.Close()

	table := "data"
	createTable(table, "id INTEGER PRIMARY KEY, name TEXT, date TEXT, bitmap TEXT")
	setupRoutes(table)

	// FIXME: parameterize LISTEN address/port
	http.ListenAndServe(":9600", nil)
}
