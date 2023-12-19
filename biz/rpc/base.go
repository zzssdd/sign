package rpc

import (
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"sign/conf"
	"sign/kitex_gen/sign/base/baseservice"
	. "sign/pkg/log"
)

var BaseClient baseservice.Client

func newBaseClient() baseservice.Client {
	r, err := etcd.NewEtcdResolver([]string{conf.GlobalConfig.DSN.EtcdDSN})
	if err != nil {
		Log.Panicf("get etcd resolver error:%v\n", err)
		return nil
	}
	client, err := baseservice.NewClient(
		conf.GlobalConfig.DSN.BaseServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		Log.Panicf("new client error:%v\n", err)
		return nil
	}
	return client
}

func init() {
	if BaseClient == nil {
		BaseClient = newBaseClient()
	}
}
