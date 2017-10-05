package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	defer db.Close()
	createTable("data", "id INTEGER PRIMARY KEY, name TEXT, date DATE, bitmap INTEGER")
	setupRoutes()
	http.ListenAndServe(":9600", nil)
}
