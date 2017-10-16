package server

import (
	. "github.com/roobert/dms/db"
)

func init() {
	CreateDB("dms.db")
}
