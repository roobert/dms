package handler

import (
	"fmt"

	. "github.com/roobert/dms/server/client"
	. "github.com/roobert/golang-db"
	. "github.com/roobert/golang-error"
)

func UpdateClient(c Client) {
	query := fmt.Sprintf("INSERT OR IGNORE INTO clients (name, date) VALUES ('%s', '%s')", c.Name, c.Date())

	fmt.Println(query)
	_, err := DB.Exec(query)
	CheckErr(err)
}

func UpdateResult(c Client) {
	query := fmt.Sprintf("INSERT OR IGNORE INTO results (id, slot) "+
		"SELECT id, %v FROM clients WHERE name = '%s' AND date = '%v'", c.Slot(), c.Name, c.Date())

	fmt.Println(query)
	_, err := DB.Exec(query)
	CheckErr(err)
}
