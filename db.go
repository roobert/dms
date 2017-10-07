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

func rowExists(table, cond string) bool {
	var exists bool

	query := fmt.Sprintf("SELECT exists (SELECT * FROM %s WHERE %s)", table, cond)
	err := db.QueryRow(query).Scan(&exists)
	checkErr(err)

	return exists
}

func multipleRowsExist(table, cond string) bool {
	var count int

	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE %s", table, cond)
	rows := db.QueryRow(query)
	err := rows.Scan(&count)
	checkErr(err)

	if count > 1 {
		return true
	} else {
		return false
	}
}

// FIXME: do this using a transaction?
func upsertRow(table string, cond string, bitmap []bool) {
	if rowExists(table, cond) == true {
		updateRow(table, cond, bitmap)
	} else {
		insertRow(table, cond, bitmap)
	}
}

func updateRow(t string, cond string, bitmap []bool) {
	fmt.Println(t, cond, bitmap)
	fmt.Println("updating")
}

func insertRow(t string, cond string, bitmap []bool) {
	fmt.Println(t, cond, bitmap)
	fmt.Println("inserting")
}
