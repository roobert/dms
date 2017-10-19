package main

import (
	_ "github.com/roobert/dms/common"
	. "github.com/roobert/golang-db"
)

func main() {
	defer DB.Close()
}
