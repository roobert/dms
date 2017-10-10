package main

import "fmt"

func upsertClient(c Client) {
	query := fmt.Sprintf("INSERT OR IGNORE INTO clients (name, date) VALUES ('%s', '%s')", c.Name, c.Date())

	fmt.Println(query)
	_, err := db.Exec(query)
	checkErr(err)
}

func upsertResult(c Client) {
	query := fmt.Sprintf("INSERT OR IGNORE INTO results (id, slot) "+
		"SELECT id, %v FROM clients WHERE name = '%s' AND date = '%v'", c.Slot(), c.Name, c.Date())

	fmt.Println(query)
	_, err := db.Exec(query)
	checkErr(err)
}
