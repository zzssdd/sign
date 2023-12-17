package cache

import (
	"github.com/gomodule/redigo/redis"
	"sign/conf"
	"time"
)

type Cache struct {
	*User
	*Sign
	*Prize
	*Group
	*Order
	*Activity
}

var CachePool *redis.Pool

func NewCache(conf *conf.Config) *Cache {
	if CachePool == nil {
		CachePool = &redis.Pool{
			MaxIdle:     100,
			MaxActive:   12000,
			IdleTimeout: time.Duration(180),
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", conf.DSN.RedisDSN)
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return &Cache{
		newUser(conf.Cache),
		newSign(conf.Cache),
		newPrize(conf.Cache),
		newGroup(conf.Cache),
		newOrder(conf.Cache),
		newActivity(conf.Cache),
	}
}
