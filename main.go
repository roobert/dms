package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	defer db.Close()

	// FIXME: create gc routine
	createTables()
	setupRoutes()

	// FIXME: error on bind fail
	// FIXME: parameterize LISTEN address/port
	http.ListenAndServe(":9600", nil)
}
