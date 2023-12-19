package base

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"sign/conf"
	"sign/dao/cache"
	"sign/dao/db"
	"sign/dao/mq"
	"sign/kitex_gen/sign/base/baseservice"
	. "sign/pkg/log"
)

type BaseServiceImpl struct {
	db    *db.DB
	cache *cache.Cache
	mq    *mq.RabbitConn
	conf  *conf.Config
}

func BaseServiceStart(config *conf.Config) {
	baseService := &BaseServiceImpl{
		db:    db.NewDB(config),
		cache: cache.NewCache(config),
		mq:    mq.NewRabbitConn(config),
		conf:  config,
	}
	r, err := etcd.NewEtcdRegistry([]string{config.DSN.EtcdDSN})
	if err != nil {
		Log.Errorf("get etcd registry error:%v\n", err)
		return
	}
	addr, err := net.ResolveTCPAddr("tcp", config.DSN.BaseService)
	if err != nil {
		Log.Errorf("registry base tcp addr err:", err)
	}
	svc := baseservice.NewServer(baseService,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.DSN.BaseServiceName}),
	)
	if err = svc.Run(); err != nil {
		Log.Fatalf("base Service run error:%v\n", err)
	}
}
