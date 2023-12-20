package db

import (
	"fmt"
	"sign/conf"
	. "sign/pkg/log"
	"sign/utils"
	"strconv"
	"time"
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
	return fmt.Sprintf("user_%d", u.sliceMap[id%u.mod])
}

func (u *User) getScoreTable(id int64) string {
	return fmt.Sprintf("user_score_%d", u.sliceMap[id%u.mod])
}

func (u *User) Register(email, password string) (int64, int64, error) {
	var count int64
	tx, err := commonDB.user.Begin()
	if err != nil {
		Log.Errorf("commonDB.user.Begin() error:%v\n", err)
		return -1, -1, err
	}
	row := tx.QueryRow("SELECT count(*) FROM email_id WHERE email=?", email)
	if row.Err() == nil && row.Scan(&count) == nil && count > 0 {
		return 0, -1, fmt.Errorf("邮箱已被注册")
	}
	id := u.snowflow.GenID()
	result, err := tx.Exec(fmt.Sprintf("INSERT INTO %s(id,email,password,created_at) VALUES (?,?,?,?)", u.getUserTable(id)), id, email, password, time.Now())
	if err != nil {
		tx.Rollback()
		Log.Errorf("insert into user table err:%v\n", err)
		return 0, -1, err
	}
	affect, err := result.RowsAffected()
	_, err = tx.Exec("INSERT INTO email_id VALUES (?,?)", id, email)
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	_, err = tx.Exec(fmt.Sprintf("INSERT INTO %s(id) VALUES (?)", u.getScoreTable(id)), id)
	err = tx.Commit()
	return id, affect, err
}

func (u *User) Login(email string) (int64, string, error) {
	var id int64
	var realPassword string
	var err error
	row := commonDB.user.QueryRow("SELECT id FROM email_id WHERE email=?", email)
	if err = row.Scan(&id); err != nil {
		Log.Errorf("query from user error:%v\n", err)
		return id, "", err
	}
	row2 := commonDB.user.QueryRow(fmt.Sprintf("SELECT password FROM %s WHERE email=?", u.getUserTable(id)), email)
	if err = row2.Scan(&realPassword); err != nil {
		Log.Errorf("row2.Scan error:%v\n", err)
		return id, "", err
	}
	return id, realPassword, nil
}

func (u *User) GetUserScore(id int64) (int64, error) {
	var score int64
	var err error
	row := commonDB.user.QueryRow(fmt.Sprintf("SELECT score FROM %s WHERE id=?", u.getUserTable(id)), id)
	if err = row.Err(); err != nil {
		return 0, err
	}
	err = row.Scan(&score)
	return score, err
}

func (u *User) AddUserScore(id int64, incr int32) error {
	exec, err := commonDB.user.Exec(fmt.Sprintf("UPDATE %s SET score=score+? WHERE id=?", u.getScoreTable(id)), incr, id)
	if err != nil {
		return err
	}
	affected, err := exec.RowsAffected()
	if err != nil || affected <= 0 {
		return fmt.Errorf("update score error:%v\n", err)
	}
	return nil
}

func (u *User) CheckoutAndUpdateTryScore(id int64, needScore int64) (bool, error) {
	var freezeSub int64
	var score int64
	var err error
	tx, err := commonDB.user.Begin()
	if err != nil {
		return false, err
	}
	row := tx.QueryRow("SELECT score,freezeSub FROM ? WHERE id=?", u.getScoreTable(id), id)
	if err = row.Err(); err != nil {
		return false, err
	}
	err = row.Scan(&score, &freezeSub)
	if score-freezeSub < needScore {
		tx.Rollback()
		return false, fmt.Errorf("score is not enough")
	}
	exec, err := tx.Exec(fmt.Sprintf("UPDATE %s SET freezeSub=? WHERE id=?", u.getScoreTable(id)), freezeSub+needScore, id)
	if err != nil {
		tx.Rollback()
		return false, fmt.Errorf("update freezeSub error")
	}
	affected, err := exec.RowsAffected()
	if err != nil || affected <= 0 {
		return false, fmt.Errorf("update freezeSub error")
	}
	tx.Commit()
	return true, nil
}

func (u *User) CommitScore(id int64, cost int64) error {
	var err error
	tx, err := commonDB.user.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(fmt.Sprintf("UPDATE %s SET score=score-? WHERE id=?", u.getScoreTable(id)), cost)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(fmt.Sprintf("UPDATE %s SET freezeSub=freezeSub-? WHERE id=?", u.getScoreTable(id)), cost)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u *User) CancelScore(id int64, cost int64) error {
	var err error
	_, err = commonDB.user.Exec(fmt.Sprintf("UPDATE %s SET freezeSub=freezeSub-? WHERE id=?", u.getScoreTable(id)), cost)
	if err != nil {
		return err
	}
	return nil
}
