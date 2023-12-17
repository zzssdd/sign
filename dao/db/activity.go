package db

import (
	"fmt"
	"sign/dao/db/model"
	. "sign/pkg/log"
)

type Activity struct {
}

func newActivity() *Activity {
	return &Activity{}
}

func (a *Activity) CreateActivity(activity *model.Activity) (int64, error) {
	exec, err := commonDB.choose.Exec("INSERT INTO sign_activity(gid,start_time,end_time,prizes,cost) VALUES(?,?,?,?,?)", activity.Gid, activity.Start_time, activity.End_time, activity.Prizes, activity.Cost)
	if err != nil {
		return -1, err
	}
	if count, err := exec.RowsAffected(); err != nil || count <= 0 {
		Log.Errorf("create activity error:%v", err)
		return -1, fmt.Errorf("createActivity error")
	}
	return exec.LastInsertId()
}
