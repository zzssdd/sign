package rpc

import (
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"sign/conf"
	"sign/kitex_gen/sign/sign/signservice"
	. "sign/pkg/log"
)

var SignClient signservice.Client

func newSignClient() signservice.Client {
	r, err := etcd.NewEtcdResolver([]string{conf.GlobalConfig.DSN.EtcdDSN})
	if err != nil {
		Log.Panicf("get etcd resolver error:%v\n", err)
		return nil
	}
	client, err := signservice.NewClient(
		conf.GlobalConfig.DSN.SignServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		Log.Panicf("new client error:%v\n", err)
		return nil
	}
	return client
}

func init() {
	if SignClient == nil {
		SignClient = newSignClient()
	}
}
