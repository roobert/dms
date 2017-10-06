package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	// decode the JSON and set some Client properties with the result
	decoder := json.NewDecoder(r.Body)

	var pdp postDataPrometheus

	err := decoder.Decode(&pdp)
	checkErr(err)

	c := Client{Name: pdp.Alerts[0].Annotations.Site, TimeStamp: time.Now()}

	// calculate the bitmap and set it as a property on the Client
	c.Bitmap = fetchBitmap(c)
	checkErr(err)

	// insert or update the result
	//err = upsertClientData(c)
	//checkErr(err)
}
