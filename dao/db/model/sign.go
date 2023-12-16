package model

import "time"

type SignMonth struct {
	uid    int64
	gid    int64
	month  time.Time
	bitmap int
}
