package main

import (
	//"github.com/roobert/dms/server/db"
	//. "github.com/roobert/dms/error"
	//. "github.com/roobert/dms/server/client"
	. "github.com/roobert/dms/server/db"
	//. "github.com/roobert/dms/server/handler"
	"net/http"
	//"server/handler"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	defer DB.Close()

	// FIXME: create gc routine

	CreateTables()
	SetupRoutes()

	// FIXME: error on bind fail
	// FIXME: parameterize LISTEN address/port
	http.ListenAndServe(":9600", nil)
}
