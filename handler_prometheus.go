package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	c := Client{Name: getSiteName(r), TimeStamp: time.Now()}

	upsertClient(c)
	upsertResult(c)
}

func getSiteName(r *http.Request) string {
	decoder := json.NewDecoder(r.Body)

	var pdp postDataPrometheus

	err := decoder.Decode(&pdp)
	checkErr(err)

	return pdp.Site()
}

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
