package main

import (
	"fmt"
)

func newBitmap() []bool {
	var bitmapLength = 4 * 60 * 24
	bitmap := make([]bool, bitmapLength)

	return bitmap
}

func fetchBitmap(c Client) []bool {
	var bitmap []bool

	cond := fmt.Sprintf("name = '' AND date = ''", c.Name, c.Date)

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
