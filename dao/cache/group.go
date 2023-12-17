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
	v, err := redis.Values(rds.Do("HGETALL", groupUserKey(id)))
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
	err := rds.Send("HMSET", redis.Args{}.Add(groupUserKey(id)).AddFlat(&group)...)
	if err != nil {
		return err
	}
	err = rds.Send("EXPIRE", groupKey(id), g.cahceTime)
	if err != nil {
		return err
	}
	return rds.Flush()
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

func (g *Group) GetUserGroupsInfo(id int64) (*model.UserGroups, error) {
	rds := CachePool.Get()
	defer rds.Close()
	v, err := redis.Values(rds.Do("HGETALL", groupUserKey(id)))
	if err != nil {
		return nil, err
	}
	userGroups := new(model.UserGroups)
	if err = redis.ScanStruct(v, userGroups); err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (g *Group) StoreUserGroupsInfo(id int64, userGroup *model.UserGroups) error {
	rds := CachePool.Get()
	defer rds.Close()
	err := rds.Send("HMSET", redis.Args{}.Add(groupUserKey(id)).AddFlat(&userGroup)...)
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
