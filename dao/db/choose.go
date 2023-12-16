package db

import (
	"sign/conf"
	"sign/utils"
	"strconv"
)

type Choose struct {
	sliceMap map[int64]int
	mod      int64
	snowflow *utils.SnowFlow
}

func newChoose(config *conf.Config) *Choose {
	sliceMap := make(map[int64]int)
	for indexString, nums := range config.ChooseSlice.Slice {
		index, _ := strconv.Atoi(indexString)
		for _, v := range nums {
			num, _ := strconv.ParseInt(v, 10, 64)
			sliceMap[num] = index
		}
	}
	return &Choose{
		sliceMap: sliceMap,
		mod:      config.ChooseSlice.Mod,
		snowflow: utils.NewSnowFlow(config.SnowFlow),
	}
}

func (c *Choose) GenID() int64 {
	return c.snowflow.GenID()
}
