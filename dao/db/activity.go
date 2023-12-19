package db

import (
	"fmt"
	"sign/dao/db/model"
	. "sign/pkg/log"
	"sign/utils"
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

func (a *Activity) GetActivity(aid int64) (*model.Activity, error) {
	row := commonDB.choose.QueryRow("SELECT gid,start_time,end_time,prizes,prizesTmp,cost FROM sign_activity WHERE id=?", aid)
	if err := row.Err(); err != nil {
		return nil, err
	}
	activity := new(model.Activity)
	err := row.Scan(&activity.Gid, &activity.Start_time, &activity.End_time, &activity.Prizes, &activity.PrizesTmp, &activity.Cost)
	return activity, err
}

func (a *Activity) UpdateActivityPrizes(aid int64, prizes string) error {
	exec, err := commonDB.choose.Exec("UPDATE sign_activity SET prize=? WHERE  id=?", prizes, aid)
	if err != nil {
		return err
	}
	if num, err := exec.RowsAffected(); num == 0 || err != nil {
		return fmt.Errorf("updated sign_activity error")
	}
	return nil
}

func (a *Activity) UpdateActivityTmpPrizes(aid int64, prizes string) error {
	exec, err := commonDB.choose.Exec("UPDATE sign_activity SET prizeTmp=? WHERE  id=?", prizes, aid)
	if err != nil {
		return err
	}
	if num, err := exec.RowsAffected(); num == 0 || err != nil {
		return fmt.Errorf("updated sign_activity error")
	}
	return nil
}

func (a *Activity) CheckoutAndUpdateTryNum(aid int64, pid int64) (bool, error) {
	var prize string
	tx, err := commonDB.choose.Begin()
	if err != nil {
		return false, err
	}
	row := tx.QueryRow("SELECT prizeTmp FROM sign_activity WHERE id=?", aid)
	err = row.Err()
	if err != nil {
		tx.Rollback()
		return false, err
	}
	if row.Scan(&prize) != nil {
		return false, err
	}
	subPrize, is_ok := utils.PrizeStringSubAndCheck(prize, pid)
	if !is_ok {
		tx.Rollback()
		return false, nil
	}
	exec, err := tx.Exec("UPDATE sign_activity SET prizesTmp=? WHERE id=?", subPrize, aid)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	affected, err := exec.RowsAffected()
	if err != nil || affected <= 1 {
		tx.Rollback()
		return false, fmt.Errorf("update sign_activity error")
	}
	return true, tx.Commit()
}

func (a *Activity) CommitNum(aid int64, pid int64) error {
	var prizes string
	tx, err := commonDB.choose.Begin()
	if err != nil {
		return err
	}
	row := tx.QueryRow("SELECT prize FROM sign_activity WHERE id=?", aid)
	if err = row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	err = row.Scan(&prizes)
	if err != nil {
		tx.Rollback()
		return err
	}
	subPrizes, is_ok := utils.PrizeStringSubAndCheck(prizes, pid)
	if !is_ok {
		tx.Rollback()
		return fmt.Errorf("sub and check prizes error")
	}
	exec, err := commonDB.choose.Exec("UPDATE sign_activity SET prize=? WHERE id=?", subPrizes, aid)
	if err != nil {
		return err
	}
	affected, err := exec.RowsAffected()
	if err != nil || affected <= 1 {
		return fmt.Errorf("submit tcc num error")
	}
	return nil
}

func (a *Activity) CancelNum(aid int64, pid int64) (bool, error) {
	var prize string
	tx, err := commonDB.choose.Begin()
	if err != nil {
		return false, err
	}
	row := tx.QueryRow("SELECT prizeTmp FROM sign_activity WHERE id=?", aid)
	err = row.Err()
	if err != nil {
		tx.Rollback()
		return false, err
	}
	if row.Scan(&prize) != nil {
		return false, err
	}
	subPrize := utils.PrizeStringAdd(prize, pid)
	exec, err := tx.Exec("UPDATE sign_activity SET prizesTmp=? WHERE id=?", subPrize, aid)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	affected, err := exec.RowsAffected()
	if err != nil || affected <= 1 {
		tx.Rollback()
		return false, fmt.Errorf("update sign_activity error")
	}
	return true, tx.Commit()
}
