package server

import (
	. "github.com/roobert/golang-db"
)

func init() {
	CreateDB("dms.db")
}
