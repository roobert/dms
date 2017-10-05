package main

import (
	"errors"
	"testing"
)

func TestCheckErrWithNil(t *testing.T) {
	checkErr(nil)
}

func TestCheckErrWithError(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				// panic rescued
			} else {
				t.Errorf("checkErr did not panic when passed an error")
			}
		}()

		checkErr(errors.New("test error"))
	}()

}
