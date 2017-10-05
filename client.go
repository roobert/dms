package main

import "time"

type Client struct {
	Name      string
	TimeStamp time.Time
	Bitmap    []bool
}

func (c Client) Date() string {
	return c.TimeStamp.Format("2006-01-02")
}
