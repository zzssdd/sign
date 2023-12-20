package model

import "time"

type Sign struct {
	Uid         int64
	Gid         int64
	SignInTime  string
	SignOutTime string
	Place       string
	PublishTime time.Time
	Flag        int8
}
