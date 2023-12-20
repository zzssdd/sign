package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sign/conf"
	"sign/dao/cache/model"
)

type Prize struct {
	cahceTime int
}

const PrizeKey = "prize_%d"

func newPrize(config *conf.Cache) *Prize {
	return &Prize{
		cahceTime: config.Prize,
	}
}

func prizeKey(id int64) string {
	return fmt.Sprintf(PrizeKey, id)
}

func (p *Prize) ExistAndExpirePrize(id int64) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Int(rds.Do("EXPIRE", prizeKey(id), p.cahceTime))
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (p *Prize) GetPrize(id int64) (*model.Prize, error) {
	rds := CachePool.Get()
	defer rds.Close()
	v, err := redis.Values(rds.Do("HGETALL", prizeKey(id)))
	if err != nil {
		return nil, err
	}
	prize := new(model.Prize)
	if err = redis.ScanStruct(v, prize); err != nil {
		return nil, err
	}
	return prize, nil
}

func (p *Prize) StorePrize(id int64, prize *model.Prize) error {
	rds := CachePool.Get()
	defer rds.Close()
	err := rds.Send("HMSET", redis.Args{}.Add(prizeKey(id)).AddFlat(prize)...)
	if err != nil {
		return err
	}
	err = rds.Send("EXPIRE", prizeKey(id), p.cahceTime)
	if err != nil {
		return err
	}
	err = rds.Flush()
	return err
}
