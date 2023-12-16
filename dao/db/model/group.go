package model

import "time"

type Group struct {
	Id       int64
	Name     string
	Owner    int64
	Places   string
	Sign_in  time.Time
	Sign_out time.Time
	Count    int
}
