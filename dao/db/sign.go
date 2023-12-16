package db

import (
	"fmt"
	"sign/conf"
	. "sign/pkg/log"
	"sign/utils"
	"strconv"
)

type Sign struct {
	sliceMap map[int64]int
	mod      int64
	snowflow *utils.SnowFlow
}

func newSign(config *conf.Config) *Sign {
	sliceMap := make(map[int64]int)
	for indexString, nums := range config.SignSlice.Slice {
		index, _ := strconv.Atoi(indexString)
		for _, v := range nums {
			num, _ := strconv.ParseInt(v, 10, 64)
			sliceMap[num] = index
		}
	}
	return &Sign{
		sliceMap: sliceMap,
		mod:      config.SignSlice.Mod,
		snowflow: utils.NewSnowFlow(config.SnowFlow),
	}
}

func (s *Sign) getSignMonthTable(id int64) string {
	return fmt.Sprintf("sign_month_%s", s.sliceMap[id%s.mod])
}

func (s *Sign) GetSignMonth(gid int64, uid int64) (int32, error) {
	row := commonDB.sign.QueryRow("SELECT bitmap FROM ? WHERE uid=? AND gid=?", s.getSignMonthTable(uid), uid, gid)
	if err := row.Err(); err != nil {
		Log.Errorf("select bitmap error:%v\n", err)
		return 0, err
	}
	var bit int32
	err := row.Scan(&bit)
	return bit, err
}
