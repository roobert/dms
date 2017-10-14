package client

import "time"

type Client struct {
	Name      string
	TimeStamp time.Time
}

func (c Client) Date() string {
	return c.TimeStamp.Format("2006-01-02")
}

func (c Client) Slot() int64 {
	y, m, d := c.TimeStamp.Date()
	midnight := time.Date(y, m, d, 0, 0, 0, 0, c.TimeStamp.Location())
	return (c.TimeStamp.Unix() - midnight.Unix()) / 15
}
