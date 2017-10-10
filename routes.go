package main

import (
	"net/http"
)

func setupRoutes() {
	// update API
	// FIXME: use ?format=prometheus|generic ? - default to generic?
	http.HandleFunc("/ping/prometheus", prometheusHandler)
	http.HandleFunc("/ping/generic", genericHandler)

	// read API - this should be a group

	// return all clients
	//http.HandleFunc("/api/clients/all", allClientsHandler)

	// return a single client
	//http.HandleFunc("/api/client", allClientsHandler)

	// return all results for a client
	// require param ?client=
	//http.HandleFunc("/api/client/results", clientResultsHandler)

	// web UI
	//http.HandleFunc("/", webUIHandler)
}
