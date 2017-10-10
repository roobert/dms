package main

import (
	"net/http"
)

func setupRoutes() {
	// update API
	http.HandleFunc("/prometheus", func(w http.ResponseWriter, r *http.Request) {
		prometheusHandler(w, r)
	})
	//prometheusHandler(w http.ResponseWriter, r *http.Request, table)

	// FIXME: implement generic handler to handle dms -> dms monitoring
	//http.HandleFunc("/generic", genericHandler)

	// read API - this should be a group
	//http.HandleFunc("/api/all", apiHandlerAll)

	// web UI
	//http.HandleFunc("/", webUIHandler)
}
