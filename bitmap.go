package main

//import (
//	"bytes"
//	"fmt"
//	"time"
//)

//type Bitmap struct {
//	Value []bool
//}
//
//func (b Bitmap) GoString() string {
//	var buf bytes.Buffer
//
//	for _, bit := range b.Value {
//		if bit == true {
//			buf.WriteString("1")
//		} else {
//			buf.WriteString("0")
//		}
//	}
//
//	return buf.String()
//}
//
//func newBitmap() Bitmap {
//	var bitmapLength = 4 * 60 * 24
//	bitmap := make([]bool, bitmapLength)
//
//	return Bitmap{Value: bitmap}
//}

//func fetchBitmap(table string, c Client) Bitmap {
//	var bitmapString string
//	var bitmap Bitmap
//
//	cond := fmt.Sprintf("name = '%s' AND date = '%s'", c.Name, c.Date())
//
//	if multipleRowsExist(table, cond) == true {
//		panic(fmt.Sprintf("detected multiple rows for condition: %s", cond))
//	}
//
//	// FIXME: move bitmap into a separate table rather than converting to/from string
//	// FIXME: use transaction?
//	if rowExists(table, cond) == true {
//		// FIXME: split out to getRow()?
//		query := fmt.Sprintf("SELECT bitmap FROM %s WHERE name = '%s' AND date = '%s'", table, c.Name, c.Date())
//		err := db.QueryRow(query).Scan(&bitmapString)
//		checkErr(err)
//
//		b := newBitmap()
//
//		for i, bit := range bitmapString {
//			if string(bit) == "1" {
//				b.Value[i] = true
//			} else {
//				b.Value[i] = false
//			}
//		}
//
//		bitmap = b
//
//		fmt.Println("fetched bitmap")
//	} else {
//		bitmap = newBitmap()
//		fmt.Println("created bitmap")
//	}
//
//	return bitmap
//}
//
//func updateBitmap(t time.Time, b Bitmap) Bitmap {
//	y, m, d := t.Date()
//	midnight := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
//	slot := (t.Unix() - midnight.Unix()) / 15
//	b.Value[slot] = true
//
//	return b
//}
