package main

import (
	"net/http"

	_ "github.com/roobert/dms/common"
	. "github.com/roobert/dms/server"
	. "github.com/roobert/golang-db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	defer DB.Close()

	// FIXME: create gc routine

	SetupRoutes()

	// FIXME: error on bind fail
	// FIXME: parameterize LISTEN address/port
	http.ListenAndServe(":9600", nil)
}
