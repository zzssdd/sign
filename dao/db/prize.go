package db

import (
	"fmt"
	"sign/dao/db/model"
	. "sign/pkg/log"
	"time"
)

type Prize struct {
}

func newPrize() *Prize {
	return &Prize{}
}

func (p *Prize) CreatePrize(prize *model.Prize) (int64, error) {
	exec, err := commonDB.choose.Exec("INSERT INTO sign_prizes(name,gid,created_at) VALUES(?,?,?)", prize.Name, prize.Gid, time.Now())
	if err != nil {
		return -1, err
	}
	if count, err := exec.RowsAffected(); err != nil || count <= 0 {
		Log.Errorf("create prize error:%v", err)
		return -1, fmt.Errorf("createPrize error")
	}
	return exec.LastInsertId()
}

func (p *Prize) GetPrize(id int64) (*model.Prize, error) {
	row := commonDB.choose.QueryRow("SELECT name,gid FROM sign_prizes WHERE id=?", id)
	if err := row.Err(); err != nil {
		return nil, err
	}
	prize := new(model.Prize)
	err := row.Scan(&prize.Name, prize.Gid)
	return prize, err
}
