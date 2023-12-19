package db

import (
	"fmt"
	"sign/conf"
	"sign/dao/db/model"
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
	return fmt.Sprintf("sign_month_%d", s.sliceMap[id%s.mod])
}

func (s *Sign) getSignUserDateTable(id int64) string {
	return fmt.Sprintf("sign_record_%d", s.sliceMap[id%s.mod])
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

func (s *Sign) GetSignUserDate(uid int64, gid int64, date string) (*model.SignDate, error) {
	row := commonDB.sign.QueryRow("SELECT signin_time,signout_time,signin_places,signout_places FROM ? WHERE uid=? AND gid=? AND date=?", s.getSignUserDateTable(uid), uid, gid, date)
	if err := row.Err(); err != nil {
		return nil, err
	}
	signData := new(model.SignDate)
	err := row.Scan(&signData.Signin_time, &signData.Signout_time, &signData.Signin_time, &signData.Signout_places)
	return signData, err
}

func (s *Sign) StoreSignUserData(data *model.SignDate) error {
	id := s.snowflow.GenID()
	exec, err := commonDB.sign.Exec("INSERT INTO ?(id,uid,gid,date,signin_time,signout_time,signin_places,signout_places) VALUES (?,?,?,?,?,?,?,?)", s.getSignUserDateTable(data.Uid), id, data.Uid, data.Gid, data.Date, data.Signin_time, data.Signout_time, data.Signin_places, data.Signout_places)
	if err != nil {
		return err
	}
	affected, err := exec.RowsAffected()
	if err != nil || affected <= 0 {
		return fmt.Errorf("insert into table error")
	}
	return nil
}

func (s *Sign) UpdateSignoutUserData(data *model.SignDate) error {
	exec, err := commonDB.sign.Exec("UPDATE ? SET signout_time=?,signout_places=? WHERE uid=? AND gid=? AND date=?", s.getSignUserDateTable(data.Uid), data.Signout_time, data.Signout_places, data.Uid, data.Gid, data.Date)
	affected, err := exec.RowsAffected()
	if err != nil || affected <= 0 {
		return fmt.Errorf("insert into table error")
	}
	return nil
}
