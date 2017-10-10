package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// configuration: a list of clients which we expect to hear from (optional) - default to just checking for clients which have already checked in

// setup: none

// tidy-up: purge old entries for clients (3 months or a year or whatever?)

// just write down the slot id for each incoming result - that way we dont end up with entries which are wrong, like if were to populate 'false' for a new client - it couldve been online with no results

func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	c := Client{Name: getSiteName(r), TimeStamp: time.Now()}

	slot := timeStampSlot(c.TimeStamp)

	// create row in clients table if one doesn't exist
	query := fmt.Sprintf("INSERT OR IGNORE INTO clients (name, date) VALUES ('%s', '%s')", c.Name, c.Date())
	fmt.Println(query)
	_, err := db.Exec(query)
	checkErr(err)

	result := true

	// insert row into results table for slot, if one doesn't exist
	query = fmt.Sprintf("INSERT OR IGNORE INTO results (id, slot, result) SELECT id, %v, '%v' FROM clients WHERE name = '%s' AND date = '%v'", slot, result, c.Name, c.Date())
	fmt.Println(query)
	_, err = db.Exec(query)
	checkErr(err)
}

func timeStampSlot(t time.Time) int64 {
	y, m, d := t.Date()
	midnight := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	return (t.Unix() - midnight.Unix()) / 15
}

func getSiteName(r *http.Request) string {
	decoder := json.NewDecoder(r.Body)

	var pdp postDataPrometheus

	err := decoder.Decode(&pdp)
	checkErr(err)

	return pdp.Site()
}

type postDataPrometheus struct {
	Alerts []struct {
		Annotations struct {
			Site string `json:"site`
		} `json:"annotations"`
	} `json:"alerts"`
}

func (pdp postDataPrometheus) Site() string {
	return pdp.Alerts[0].Annotations.Site
}
