package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var db *sql.DB

func createDB(d string) {
	var err error
	db, err = sql.Open("sqlite3", d)
	checkErr(err)
}

func deleteDB(d string) {
	err := os.Remove(d)
	checkErr(err)
}

func createTables() {
	clientsSchema := `
		id    INTEGER PRIMARY KEY,
		name  TEXT    NOT NULL,
		date  TEXT    NOT NULL,

		UNIQUE (date, name),
		UNIQUE (id, name)
	`

	createTable("clients", clientsSchema)

	resultsSchema := `
		id     INT  NOT NULL,
		slot   INT  NOT NULL,
		result BOOL NOT NULL,

		UNIQUE (id, slot),
		UNIQUE (slot, result)

		FOREIGN KEY (id) REFERENCES client(id)
	`

	createTable("results", resultsSchema)
}

func createTable(table, fields string) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, fields)
	_, err := db.Exec(query)
	checkErr(err)
}
