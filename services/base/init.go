package base

import (
	"sign/conf"
	"sign/dao/cache"
	"sign/dao/db"
	"sign/dao/mq"
)

type BaseServiceImpl struct {
	db    *db.DB
	cache *cache.Cache
	mq    *mq.RabbitConn
	conf  *conf.Config
}

func NewBaseServiceImpl() *BaseServiceImpl {
	config := conf.NewConfig()
	return &BaseServiceImpl{
		db:    db.NewDB(config),
		cache: cache.NewCache(config),
		mq:    mq.NewRabbitConn(config),
		conf:  config,
	}
}
