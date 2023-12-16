package model

import "time"

type Sign struct {
	Uid         int64
	Gid         int64
	SignInTime  time.Time
	SignOutTime time.Time
	Place       string
	PublishTime time.Time
}
