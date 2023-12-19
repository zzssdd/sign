package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sign/conf"
	"sign/dao/cache/model"
)

type Activity struct {
	cahceTime int
}

const ActivityKey = "activity_%d"

func newActivity(config *conf.Cache) *Activity {
	return &Activity{
		cahceTime: config.Activity,
	}
}

func activityKey(id int64) string {
	return fmt.Sprintf(ActivityKey, id)
}

func (a *Activity) ExistAndExpireActivity(id int64) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Int(rds.Do("EXPIRE", activityKey(id), a.cahceTime))
	if err != nil {
		return false, err
	}
	return reply == 1, err
}

func (a *Activity) GetActivity(id int64) (*model.Activity, error) {
	rds := CachePool.Get()
	defer rds.Close()
	v, err := redis.Values(rds.Do("HGETALL", activityKey(id)))
	if err != nil {
		return nil, err
	}
	activity := new(model.Activity)
	if err = redis.ScanStruct(v, activity); err != nil {
		return nil, err
	}
	return activity, nil
}
func (a *Activity) StoreActivity(id int64, activity *model.Activity) error {
	rds := CachePool.Get()
	defer rds.Close()
	err := rds.Send("HMSET", redis.Args{}.Add(groupUserKey(id)).AddFlat(&activity)...)
	if err != nil {
		return err
	}
	err = rds.Send("EXPIRE", groupUserKey(id), a.cahceTime)
	if err != nil {
		return err
	}
	err = rds.Flush()
	return err
}

func (a *Activity) UpdatePrizes(id int64, prizes string) error {
	rds := CachePool.Get()
	defer rds.Close()
	_, err := rds.Do("HSET", activityKey(id), "Prizes", prizes)
	return err
}
