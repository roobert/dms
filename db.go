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

func createTable(table, fields string) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, fields)
	_, err := db.Exec(query)
	checkErr(err)
}
