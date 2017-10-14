package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/roobert/dms/error"
	"os"
)

var DB *sql.DB

func CreateDB(d string) {
	var err error
	DB, err = sql.Open("sqlite3", d)
	CheckErr(err)
}

func DeleteDB(d string) {
	err := os.Remove(d)
	CheckErr(err)
}

func CreateTable(table, fields string) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, fields)
	_, err := DB.Exec(query)
	CheckErr(err)
}
