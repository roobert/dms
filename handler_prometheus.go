package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	c := Client{Name: getSiteNamePrometheusJSON(r), TimeStamp: time.Now()}

	upsertClient(c)
	upsertResult(c)
}

func getSiteNamePrometheusJSON(r *http.Request) string {
	decoder := json.NewDecoder(r.Body)

	var pdp postDataPrometheus

	err := decoder.Decode(&pdp)
	checkErr(err)

	return pdp.Site()
}
