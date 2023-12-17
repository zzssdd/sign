package utils

import (
	"strconv"
	"strings"
)

type Place struct {
	Name       string
	Latitude   float64
	Longtitude float64
}

func ParsePlacesString(s string) []*Place {
	places := strings.Split(s, " ")
	ret := []*Place{}
	for _, p := range places {
		sliceP := strings.Split(p, ",")
		latitude, _ := strconv.ParseFloat(sliceP[1], 10)
		longtitude, _ := strconv.ParseFloat(sliceP[2], 10)
		p := &Place{
			Name:       sliceP[0],
			Latitude:   latitude,
			Longtitude: longtitude,
		}
		ret = append(ret, p)
	}
	return ret
}
