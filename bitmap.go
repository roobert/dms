package main

import (
	"fmt"
	"time"
)

func newBitmap() []bool {
	var bitmapLength = 4 * 60 * 24
	bitmap := make([]bool, bitmapLength)

	return bitmap
}

func fetchBitmap(c Client) []bool {
	var bitmap []bool

	cond := fmt.Sprintf("name = '%s' AND date = '%s'", c.Name, c.Date)

	if multipleRowsExist("data", cond) == true {
		panic(fmt.Sprintf("detected multiple rows for condition: %s", cond))
	}

	if rowExists("data", cond) {
		//bitmap = selectBitmapFromDB()
		fmt.Println("select bitmap bullshit")
	} else {
		bitmap = newBitmap()
	}

	return bitmap
}

func updateBitmap(t time.Time, bitmap []bool) []bool {
	y, m, d := t.Date()
	midnight := time.Date(y, m, d, 0, 0, 0, 0, t.Location())

	slot := (midnight.Unix() - t.Unix()) / 15

	fmt.Println(slot)

	bitmap[slot] = true

	return bitmap
}
