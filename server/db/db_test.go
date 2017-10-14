package db

import (
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestCreateDB(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("createDB() panic'd when trying to create DB")
			} else {
				//deleteDB("test.db")
			}
		}()

		createDB("test.db")

		// FIXME: test if db exists..
	}()
}

//func TestDeleteDB(t *testing.T) {
//	func() {
//		defer func() {
//			if r := recover(); r != nil {
//				t.Errorf("DeleteDB() panic'd when trying to delete DB")
//			}
//		}()
//
//		createDB("test.db")
//		deleteDB("test.db")
//		// FIXME: ensure db is gone
//	}()
//}

func TestCreateTable(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("createTable() panic'd when trying to create table")
			} else {
				db.Exec("DELETE TABLE test")
			}
		}()

		createTable("test", "id INTEGER PRIMARY KEY, name TEXT")
	}()
}
