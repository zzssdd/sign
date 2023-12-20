package model

type Group struct {
	Id       int64
	Name     string
	Owner    int64
	Places   string
	Sign_in  string
	Sign_out string
	Count    int
	Score    int32
}
