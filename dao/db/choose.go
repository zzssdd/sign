package db

import (
	"fmt"
	"sign/conf"
	"sign/dao/db/model"
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

func (c *Choose) getChooseTable(id int64) string {
	return fmt.Sprintf("choose_record_%d", c.sliceMap[id%c.mod])
}

func (c *Choose) GenID() int64 {
	return c.snowflow.GenID()
}

func (c *Choose) CreateOrder(record *model.Record) error {
	id := c.snowflow.GenID()
	exec, err := commonDB.choose.Exec("INSERT INTO ?(id,uid,pid,getTime,status) VALUES(?,?,?,?,?)", c.getChooseTable(id), id, record.Uid, record.Pid, record.GetTime, "未发货")
	if err != nil {
		return err
	}
	affected, err := exec.RowsAffected()
	if err != nil || affected <= 0 {
		return fmt.Errorf("create order error")
	}
	return nil
}
