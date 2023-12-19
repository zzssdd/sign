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
	Signin_time    string
	Signout_time   string
	Signin_places  string
	Signout_places string
}
