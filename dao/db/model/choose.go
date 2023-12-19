package model

import "time"

type Record struct {
	Uid     int64
	Pid     int64
	GetTime time.Time
}
