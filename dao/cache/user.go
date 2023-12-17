package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sign/conf"
	"sign/dao/cache/model"
	"strconv"
)

type User struct {
	cahceTime int
}

const (
	UserKey      = "user_%s"
	UserScoreKey = "user_score_%d"
)

func newUser(config *conf.Cache) *User {
	return &User{
		cahceTime: config.User,
	}
}

func userKey(email string) string {
	return fmt.Sprintf(UserKey, email)
}

func userScoreKey(id int64) string {
	return fmt.Sprintf(UserScoreKey, id)
}

func (u *User) ExistAndExpireUser(email string) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Int(rds.Do("EXPIRE", userKey(email), u.cahceTime))
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (u *User) CheckLogin(email string, password string) (bool, int64, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.StringMap(rds.Do("HGETALL", userKey(email), "password"))
	if err != nil {
		return false, -1, err
	}
	if reply["password"] != password {
		return false, -1, fmt.Errorf("password error")
	}
	id, err := strconv.ParseInt(reply["id"], 10, 64)
	return true, id, err
}

func (u *User) StoreUser(userInfo *model.User) error {
	rds := CachePool.Get()
	defer rds.Close()
	err := rds.Send("HMSET", redis.Args{}.Add(userKey(userInfo.Email)).AddFlat(&userInfo)...)
	if err != nil {
		return err
	}
	err = rds.Send("EXPIRE", userKey(userInfo.Email), u.cahceTime)
	if err != nil {
		return err
	}
	err = rds.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ExistAndExpireUserScore(id int64) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Int(rds.Do("EXPIRE", userScoreKey(id), u.cahceTime))
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (u *User) GetUserScore(id int64) (int64, error) {
	rds := CachePool.Get()
	defer rds.Close()
	return redis.Int64(rds.Do("GET", userScoreKey(id)))
}

func (u *User) DeleteUserScore(id int64) error {
	rds := CachePool.Get()
	defer rds.Close()
	_, err := rds.Do("DEL", userScoreKey(id))
	return err
}
