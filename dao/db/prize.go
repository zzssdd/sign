package db

import (
	"fmt"
	"sign/dao/db/model"
	. "sign/pkg/log"
)

type Prize struct {
}

func newPrize() *Prize {
	return &Prize{}
}

func (p *Prize) CreatePrize(prize *model.Prize) error {
	exec, err := commonDB.choose.Exec("INSERT INTO sign_prizes(name,gid) VALUES(?,?)", prize.Name, prize.Gid)
	if count, err := exec.RowsAffected(); err != nil || count <= 0 {
		Log.Errorf("create prize error:%v", err)
		return fmt.Errorf("createPrize error")
	}
	return err
}
