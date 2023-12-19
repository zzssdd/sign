package model

import "time"

type Activity struct {
	Gid        int64
	Start_time time.Time
	End_time   time.Time
	Prizes     string
	PrizesTmp  string
	Cost       int64
}
