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

// FIXME: potential race condition
func upsertRow(table string, c Client) {
	cond := fmt.Sprintf("name = '%s' AND date = '%s'", c.Name, c.Date())

	fmt.Println(rowExists(table, cond))

	if rowExists(table, cond) == true {
		fmt.Println("updating")
		updateRow(table, c)
	} else {
		fmt.Println("inserting")
		insertRow(table, c)
	}
}

func updateRow(table string, c Client) {
	query := fmt.Sprintf("UPDATE %s SET bitmap = '%#v' WHERE name = '%s' AND date = '%s'", table, c.Bitmap, c.Name, c.Date())
	_, err := db.Exec(query)
	checkErr(err)
}

func insertRow(table string, c Client) {
	query := fmt.Sprintf("INSERT INTO %s (name, date, bitmap) VALUES ('%s', '%s', '%#v')", table, c.Name, c.Date(), c.Bitmap)
	_, err := db.Exec(query)
	checkErr(err)
}
