package sign

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
	"sign/kitex_gen/sign/sign/signservice"
	. "sign/pkg/log"
)

type SignServiceImpl struct {
	db    *db.DB
	cache *cache.Cache
	mq    *mq.RabbitConn
	conf  *conf.Config
}

func SignServiceStart(config *conf.Config) {
	signServer := &SignServiceImpl{
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
	addr, err := net.ResolveTCPAddr("tcp", config.DSN.SignService)
	if err != nil {
		Log.Errorf("registry base tcp addr err:", err)
	}
	svc := signservice.NewServer(signServer,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.DSN.SignServiceName}),
	)
	go func() {
		_, _ = signServer.Sign(context.Background(), nil)
	}()
	if err = svc.Run(); err != nil {
		Log.Fatalf("sign Service run error:%v\n", err)
	}
}
