package main

import (
	"net/http"
)

func setupRoutes() {
	// update API
	http.HandleFunc("/prometheus", prometheusHandler)
	// FIXME: implement generic handler to handle dms -> dms monitoring
	//http.HandleFunc("/generic", genericHandler)

	// read API - this should be a group
	//http.HandleFunc("/api/all", apiHandlerAll)

	// web UI
	//http.HandleFunc("/", webUIHandler)
}
