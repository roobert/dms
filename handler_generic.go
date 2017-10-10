package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func genericHandler(w http.ResponseWriter, r *http.Request) {
	c := Client{Name: getSiteNameGenericJSON(r), TimeStamp: time.Now()}

	upsertClient(c)
	upsertResult(c)
}

func getSiteNameGenericJSON(r *http.Request) string {
	decoder := json.NewDecoder(r.Body)

	var pdg postDataGeneric

	err := decoder.Decode(&pdg)
	checkErr(err)

	return pdg.Site
}
