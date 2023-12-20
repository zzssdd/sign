package model

import "time"

type SignMonth struct {
	uid    int64
	gid    int64
	month  time.Time
	bitmap int
}

type SignDate struct {
	Uid            int64
	Gid            int64
	Date           string
	Signin_time    string
	Signout_time   string
	Signin_places  string
	Signout_places string
}
