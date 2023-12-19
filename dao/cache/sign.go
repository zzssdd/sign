package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sign/conf"
	"sign/dao/cache/model"
)

type Sign struct {
	cahceTime int
}

const (
	SignKey   = "sign_%d_%d_%s"
	SignPos   = "sign_pos_%d"
	SignMonth = "sign_month_%s_%d"
)

func newSign(config *conf.Cache) *Sign {
	return &Sign{
		cahceTime: config.Sign,
	}
}

func signKey(uid int64, gid int64, date string) string {
	return fmt.Sprintf(SignKey, uid, gid, date)
}

func signPos(id int64) string {
	return fmt.Sprintf(SignPos, id)
}

func signMonth(uid int64, month string) string {
	return fmt.Sprintf(SignMonth, uid, month)
}

func (s *Sign) ExistSignPos(gid int64) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := rds.Do("EXPIRE", signPos(gid), s.cahceTime)
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (s *Sign) AddSignPos(pos *model.SignPos) error {
	rds := CachePool.Get()
	defer rds.Close()
	err := rds.Send("GEOADD", signPos(pos.Gid), pos.Latitle, pos.Longtitude, pos.Name)
	if err != nil {
		return err
	}
	err = rds.Send("EXPIRE", signPos(pos.Gid), s.cahceTime)
	if err != nil {
		return err
	}
	err = rds.Flush()
	return err
}

func (s *Sign) JudgeSignPos(gid int64, latitude float64, longtitude float64) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Values(rds.Do("GEORADIUS", signPos(gid), latitude, longtitude, 1, "km", "withcoord"))
	if err != nil {
		return false, err
	}
	return len(reply) > 0, nil
}

func (s *Sign) ExistAndExpireMonth(sign *model.SignMonth) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := rds.Do("EXPIRE", signMonth(sign.Uid, sign.Month), s.cahceTime)
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (s *Sign) SetSignMonth(sign *model.SignMonth, bit int32) error {
	rds := CachePool.Get()
	defer rds.Close()
	_, err := rds.Do("SET", signMonth(sign.Uid, sign.Month), bit)
	return err
}

func (s *Sign) AddSignMonth(sign *model.SignMonth) error {
	rds := CachePool.Get()
	defer rds.Close()
	_, err := rds.Do("SETBIT", signMonth(sign.Uid, sign.Month), sign.Day, 1)
	return err
}

func (s *Sign) GetSignMonth(sign *model.SignMonth) (int32, error) {
	rds := CachePool.Get()
	defer rds.Close()
	bits, err := redis.Int(rds.Do("GET", signMonth(sign.Uid, sign.Month)))
	return int32(bits), err
}

func (s *Sign) ExistAndExpireSign(uid, gid int64, date string) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Int(rds.Do("EXPIRE", signKey(uid, gid, date), s.cahceTime))
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (s *Sign) GetUserSign(uid, gid int64, date string) (*model.Sign, error) {
	rds := CachePool.Get()
	defer rds.Close()
	v, err := redis.Values(rds.Do("HGETALL", signKey(uid, gid, date)))
	if err != nil {
		return nil, err
	}
	sign := new(model.Sign)
	_, err = redis.Scan(v, &sign)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

func (s *Sign) UserSign(uid, gid int64, date string, sign *model.Sign) error {
	rds := CachePool.Get()
	defer rds.Close()
	err := rds.Send("HMSET", redis.Args{}.Add(signKey(uid, gid, date)).AddFlat(&sign)...)
	if err != nil {
		return err
	}
	err = rds.Send("EXPIRE", signKey(uid, gid, date), s.cahceTime)
	if err != nil {
		return err
	}
	err = rds.Flush()
	return err
}
