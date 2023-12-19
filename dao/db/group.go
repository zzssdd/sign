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

func (g *Group) CreateGroup(info *model.Group) (int64, error) {
	exec, err := commonDB.user.Exec("INSERT INTO groupInfo(name,owner,places,sign_in,sign_out,created_at) VALUES (?,?,?,?,?,?)", info.Name, info.Owner, info.Places, info.Sign_in, info.Sign_out, time.Now())
	if err != nil {
		Log.Errorf("create group error:%v\n", err)
		return -1, err
	}
	if num, _ := exec.RowsAffected(); num <= 0 {
		return -1, fmt.Errorf("创建群组失败")
	}
	return exec.LastInsertId()
}

func (g *Group) GetGroup(gid int64) (*model.Group, error) {
	var group *model.Group
	var err error
	row := commonDB.user.QueryRow("SELECT name,owner,places,sign_in,sign_out,count FROM groupInfo WHERE id=?", gid)
	if err = row.Err(); err != nil {
		Log.Errorf("select groupInfo from group error:%v\n", err)
		return nil, err
	}
	err = row.Scan(&group.Name, &group.Owner, &group.Places, &group.Sign_in, &group.Sign_out, &group.Count)
	if err != nil {
		return nil, err
	}
	return group, nil
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

func (g *Group) GetUserGroups(uid int64) (string, error) {
	var err error
	var groups string
	row := commonDB.user.QueryRow("SELECT groups FROM ? WHERE uid=?", g.getUserGroupTable(uid), uid)
	if err = row.Err(); err != nil {
		fmt.Errorf("get user groups from db error:%v\n", err)
		return "", err
	}
	err = row.Scan(&groups)
	return groups, err
}
func (g *Group) UpdateGroupsPrizes(uid int64, groups string) error {
	exec, err := commonDB.choose.Exec("UPDATE ? SET groups=? WHERE  uid=?", g.getUserGroupTable(uid), groups, uid)
	if err != nil {
		return err
	}
	if num, err := exec.RowsAffected(); num == 0 || err != nil {
		return fmt.Errorf("updated user_group error")
	}
	return nil
}
