package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	defer db.Close()

	createTables()
	setupRoutes()

	// FIXME: parameterize LISTEN address/port
	http.ListenAndServe(":9600", nil)
}
