package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewBitmapLength(t *testing.T) {
	bitmap := newBitmap()

	if len(bitmap.Value) != 5760 {
		t.Errorf("incorrect bitmap length")
	}
}

func TestNewBitmapType(t *testing.T) {
	bitmap := newBitmap()

	if reflect.TypeOf(bitmap.Value).Elem().Kind() != reflect.Bool {
		t.Errorf("bitmap value is not []bool type")
	}
}

// FIXME: split tests into separate functions
func TestUpdateBitmap(t *testing.T) {
	bitmap := newBitmap()
	timeStamp := time.Now()

	bitmap = updateBitmap(timeStamp, bitmap)

	// test bitmap contains one instance of true
	var count int

	for _, bit := range bitmap.Value {
		if bit == true {
			count += 1
		}
	}

	if count != 1 {
		t.Errorf("incorrect number of true values in updated bitmap")
	}

	// test updated bitmap slot for timestamp contains "true"
	y, m, d := timeStamp.Date()
	midnight := time.Date(y, m, d, 0, 0, 0, 0, timeStamp.Location())
	slot := (timeStamp.Unix() - midnight.Unix()) / 15

	if bitmap.Value[slot] != true {
		t.Errorf("failed to update bitmap, slot value does not equal 'true'")
	}
}
