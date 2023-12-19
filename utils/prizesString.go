package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type PrizePair struct {
	Pid int64
	Num int64
}

func ParsePrizeString(s string) []*PrizePair {
	prizesString := strings.Split(s, " ")
	prizesPair := []*PrizePair{}
	for _, v := range prizesString {
		prizePair := strings.Split(v, ",")
		Pid, _ := strconv.ParseInt(prizePair[0], 10, 64)
		Num, _ := strconv.ParseInt(prizePair[1], 10, 64)
		prizesPair = append(prizesPair, &PrizePair{
			Pid: Pid,
			Num: Num,
		})
	}
	return prizesPair
}

func PackPrizeString(prizes []*PrizePair) string {
	var prizesString strings.Builder
	for i := 0; i < len(prizes); i++ {
		if i != 0 {
			prizesString.WriteString(" ")
		}
		prizesString.WriteString(fmt.Sprintf("%d,%d", prizes[i].Pid, prizes[i].Num))
	}
	return prizesString.String()
}

func PrizeStringSubAndCheck(s string, pid int64) (string, bool) {
	prizes := ParsePrizeString(s)
	for _, v := range prizes {
		if v.Pid == pid {
			v.Num--
			if v.Num <= 0 {
				return "", false
			}
			break
		}
	}
	return PackPrizeString(prizes), true
}

func PrizeStringAdd(s string, pid int64) string {
	prizes := ParsePrizeString(s)
	for _, v := range prizes {
		if v.Pid == pid {
			v.Num++
			break
		}
	}
	return PackPrizeString(prizes)
}
