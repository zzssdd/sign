package db

import (
	"fmt"
	"sign/conf"
	. "sign/pkg/log"
	"sign/utils"
	"strconv"
)

type User struct {
	sliceMap map[int64]int
	mod      int64
	snowflow *utils.SnowFlow
}

func newUser(config *conf.Config) *User {
	sliceMap := make(map[int64]int)
	for indexString, nums := range config.UserSlice.Slice {
		index, _ := strconv.Atoi(indexString)
		for _, v := range nums {
			num, _ := strconv.ParseInt(v, 10, 64)
			sliceMap[num] = index
		}
	}
	return &User{
		sliceMap: sliceMap,
		mod:      config.UserSlice.Mod,
		snowflow: utils.NewSnowFlow(config.SnowFlow),
	}
}

func (u *User) getUserTable(id int64) string {
	return fmt.Sprintf("user_%s", u.sliceMap[id%u.mod])
}

func (u *User) Register(email, password string) (int64, error) {
	var count int64
	tx, err := commonDB.user.Begin()
	if err != nil {
		Log.Errorf("commonDB.user.Begin() error:%v\n", err)
		return -1, err
	}
	row := tx.QueryRow("SELECT count(*) FROM email_id WHERE email=?", email)
	if row.Err() == nil && row.Scan(&count) == nil && count >= 0 {
		return 0, fmt.Errorf("邮箱已被注册")
	}
	id := u.snowflow.GenID()
	result, err := commonDB.user.Exec("INSERT INTO ? VALUES (?,?,?)", u.getUserTable(id), id, email, password)
	if err != nil {
		Log.Errorf("insert into user table err:%v\n", err)
		return 0, err
	}
	return result.RowsAffected()
}

func (u *User) Login(email, password string) (int64, error) {
	var id int64
	var count int64
	var err error
	row := commonDB.user.QueryRow("SELECT id FROM email_id WHERE email=?", email)
	if err = row.Scan(&id); err != nil {
		Log.Errorf("query from user error:%v\n", err)
		return -1, err
	}
	row2 := commonDB.user.QueryRow("SELECT count(*) FROM ? WHERE email=? AND password=?", u.getUserTable(id), email, password)
	if err = row2.Scan(&count); err != nil {
		Log.Errorf("row2.Scan error:%v\n", err)
		return -1, err
	}
	if count <= 0 {
		return -1, nil
	}
	return id, nil
}
