package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sign/conf"
	"sign/dao/cache/model"
)

type Group struct {
	cahceTime int
}

func newGroup(config *conf.Cache) *Group {
	return &Group{
		cahceTime: config.Group,
	}
}

const (
	UserGroupsKey = "user_groups_%d"
	GroupKey      = "group_%d"
)

func groupKey(id int64) string {
	return fmt.Sprintf(GroupKey, id)
}

func groupUserKey(id int64) string {
	return fmt.Sprintf(UserGroupsKey, id)
}

func (g *Group) ExistAndExpireGroup(id int64) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Int(rds.Do("EXPIRE", groupKey(id), g.cahceTime))
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (g *Group) GetGroup(id int64) (*model.Group, error) {
	rds := CachePool.Get()
	defer rds.Close()
	v, err := redis.Values(rds.Do("HGETALL", groupKey(id)))
	if err != nil {
		return nil, err
	}
	group := new(model.Group)
	if err = redis.ScanStruct(v, group); err != nil {
		return nil, err
	}
	return group, nil
}

func (g *Group) StoreGroup(id int64, group *model.Group) error {
	rds := CachePool.Get()
	defer rds.Close()
	err := rds.Send("HMSET", redis.Args{}.Add(groupKey(id)).AddFlat(group)...)
	if err != nil {
		return err
	}
	err = rds.Send("EXPIRE", groupKey(id), g.cahceTime)
	if err != nil {
		return err
	}
	err = rds.Flush()
	return err
}

func (g *Group) IncrGroupCount(id int64) error {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := rds.Do("HINCR", groupKey(id), "count")
	if err != nil {
		return err
	}
	if reply == 0 {
		return fmt.Errorf("incr group count error")
	}
	return nil
}

func (g *Group) ExistAndExpireUserGroups(uid int64) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Int(rds.Do("EXPIRE", groupUserKey(uid), g.cahceTime))
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (g *Group) GetUserGroupsInfo(id int64) (string, error) {
	rds := CachePool.Get()
	defer rds.Close()
	v, err := redis.String(rds.Do("GET", groupUserKey(id)))
	if err != nil {
		return "", err
	}
	return v, nil
}

func (g *Group) DelUserGroupsInfo(id int64) error {
	rds := CachePool.Get()
	defer rds.Close()
	_, err := rds.Do("DEL", groupKey(id))
	return err
}

func (g *Group) StoreUserGroupsInfo(id int64, groups string) error {
	rds := CachePool.Get()
	defer rds.Close()
	_, err := rds.Do("SETEX", groupUserKey(id), g.cahceTime, groups)
	return err
}
