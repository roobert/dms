package main

import (
	"reflect"
	"testing"
)

func TestNewBitmapLength(t *testing.T) {
	bitmap := newBitmap()

	if len(bitmap) != 5760 {
		t.Errorf("newBitmap() does not return a bitmap with which is 5760 in length")
	}
}

func TestNewBitmapType(t *testing.T) {
	bitmap := newBitmap()

	if reflect.TypeOf(bitmap).Elem().Kind() != reflect.Bool {
		t.Errorf("newBitmap() does not return a []bool type")
	}
}
