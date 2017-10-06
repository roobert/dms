package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	c := Client{Name: getSiteName(r), TimeStamp: time.Now()}
	c.Bitmap = fetchBitmap(c)

	upsertClientData(c)
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
