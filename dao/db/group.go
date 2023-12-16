package db

import (
	"fmt"
	"sign/conf"
	"sign/dao/db/model"
	. "sign/pkg/log"
	"sign/utils"
	"strconv"
	"time"
)

type Group struct {
	sliceMap map[int64]int
	mod      int64
	snowflow *utils.SnowFlow
}

func (g *Group) getUserGroupTable(id int64) string {
	return fmt.Sprintf("user_group_%s", g.sliceMap[id%g.mod])
}

func newGroup(config *conf.Config) *Group {
	sliceMap := make(map[int64]int)
	for indexString, nums := range config.GroupSlice.Slice {
		index, _ := strconv.Atoi(indexString)
		for _, v := range nums {
			num, _ := strconv.ParseInt(v, 10, 64)
			sliceMap[num] = index
		}
	}
	return &Group{
		sliceMap: sliceMap,
		mod:      config.GroupSlice.Mod,
		snowflow: utils.NewSnowFlow(config.SnowFlow),
	}
}

func (g *Group) CreateGroup(info *model.Group) error {
	exec, err := commonDB.user.Exec("INSERT INTO groupInfo(name,owner,places,sign_in,sign_out,created_at) VALUES (?,?,?,?,?,?)", info.Name, info.Owner, info.Places, info.Sign_in, info.Sign_out, time.Now())
	if err != nil {
		Log.Errorf("create group error:%v\n", err)
		return err
	}
	if num, _ := exec.RowsAffected(); num <= 0 {
		return fmt.Errorf("创建群组失败")
	}
	return err
}

func (g *Group) JoinGroup(uid, gid int64) (int64, error) {
	id := g.snowflow.GenID()
	var groups string
	tx, err := commonDB.user.Begin()
	if err != nil {
		return -1, err
	}
	row := commonDB.user.QueryRow("SELECT groups FROM ? WHERE uid=? AND gid=?", g.getUserGroupTable(uid), uid, gid)
	if err := row.Scan(&groups); err != nil {
		Log.Errorf("query groups error:%v\n", err)
		tx.Rollback()
		return -1, err
	}
	if utils.StringContainInt64(groups, gid) {
		tx.Rollback()
		return -1, fmt.Errorf("用户已经加入该群组")
	}
	newGroups := utils.AddInt64ToString(groups, gid)
	exec, err := tx.Exec("INSERT INTO ?(id,uid,group,) VALUES (?,?)", g.getUserGroupTable(uid), id, uid, newGroups)
	if err != nil {
		Log.Errorf("insert into groups error:%v\n", err)
		tx.Rollback()
	}
	tx.Commit()
	return exec.RowsAffected()
}
