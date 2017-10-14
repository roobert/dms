package main

import (
	. "github.com/roobert/dms/server/db"
)

func init() {
	CreateDB("dms.db")
}
