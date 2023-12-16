package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sign/conf"
	. "sign/pkg/log"
)

type DB struct {
	*User
	*Group
	*Sign
	*Choose
	*Prize
	*Activity
}

type mysqlConn struct {
	user   *sql.DB
	sign   *sql.DB
	choose *sql.DB
	order  *sql.DB
	conf   *conf.Config
}

var commonDB *mysqlConn

func NewDB(conf *conf.Config) *DB {
	if commonDB == nil {

		userDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.DSN.UserNameDB, conf.DSN.PassWordDB, conf.DSN.MysqlDSN, conf.DSN.UserDB))
		if err != nil {
			Log.Panicf("connect user db error:%v\n", err)
		}
		signDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.DSN.UserNameDB, conf.DSN.PassWordDB, conf.DSN.MysqlDSN, conf.DSN.SignDB))
		if err != nil {
			Log.Panicf("connect sign db error:%v\n", err)
		}
		chooseDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.DSN.UserNameDB, conf.DSN.PassWordDB, conf.DSN.MysqlDSN, conf.DSN.ChooseDB))
		if err != nil {
			Log.Panicf("connect choose db error:%v\n", err)
		}
		orderDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.DSN.UserNameDB, conf.DSN.PassWordDB, conf.DSN.MysqlDSN, conf.DSN.OrderDB))
		if err != nil {
			Log.Panicf("connect choose db error:%v\n", err)
		}
		commonDB = &mysqlConn{
			user:   userDb,
			sign:   signDb,
			choose: chooseDb,
			order:  orderDb,
			conf:   conf,
		}
	}
	return &DB{
		newUser(commonDB.conf),
		newGroup(commonDB.conf),
		newSign(commonDB.conf),
		newChoose(commonDB.conf),
		newPrize(),
		newActivity(),
	}
}
