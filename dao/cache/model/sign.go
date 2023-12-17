package model

type SignPos struct {
	Gid        int64
	Name       string
	Latitle    float64
	Longtitude float64
}

type SignMonth struct {
	Uid   int64
	Month string
	Day   int
}

type Sign struct {
	signin_time    string
	signout_time   string
	signin_places  string
	signout_places string
}
