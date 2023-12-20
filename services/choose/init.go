package choose

import (
	"context"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"sign/conf"
	"sign/dao/cache"
	"sign/dao/db"
	"sign/dao/mq"
	"sign/kitex_gen/sign/choose/chooseservice"
	. "sign/pkg/log"
)

type ChooseServiceImpl struct {
	db    *db.DB
	cache *cache.Cache
	mq    *mq.RabbitConn
	conf  *conf.Config
}

func ChooseServiceStart(config *conf.Config) {
	chooseServer := &ChooseServiceImpl{
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
	addr, err := net.ResolveTCPAddr("tcp", config.DSN.ChooseService)
	if err != nil {
		Log.Errorf("registry base tcp addr err:", err)
	}
	svc := chooseservice.NewServer(chooseServer,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.DSN.ChooseServiceName}),
	)
	go func() {
		_, _ = chooseServer.Choose(context.Background(), nil)
	}()
	if err = svc.Run(); err != nil {
		Log.Fatalf("choose Service run error:%v\n", err)
	}
}
