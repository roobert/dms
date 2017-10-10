package main

//import (
//	"encoding/json"
//	"net/http"
//	"time"
//)
//
//func pingHandler(w http.ResponseWriter, r *http.Request) {
//
//	c = prometheusHandler(w, r)
//	c = genericHandler(w, r)
//
//	updateClient(c)
//	updateResult(c)
//}
//
//func getSiteNamePrometheusJSON(r *http.Request) string {
//	decoder := json.NewDecoder(r.Body)
//
//	var pdp postDataPrometheus
//
//	err := decoder.Decode(&pdp)
//	checkErr(err)
//
//	return pdp.Site()
//}
