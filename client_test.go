package main

import (
	"testing"
	"time"
)

func TestClientDateMethod(t *testing.T) {
	c := Client{Name: "test", TimeStamp: time.Now(), Bitmap: newBitmap()}

	date := time.Now().Format("2006-01-02")

	if c.Date() != date {
		t.Errorf("Client Date method does not return date in expected format")
	}
}
