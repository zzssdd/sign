package rpc

import (
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"sign/conf"
	"sign/kitex_gen/sign/choose/chooseservice"
	. "sign/pkg/log"
)

var ChooseClient chooseservice.Client

func newChooseClient() chooseservice.Client {
	r, err := etcd.NewEtcdResolver([]string{conf.GlobalConfig.DSN.EtcdDSN})
	if err != nil {
		Log.Panicf("get etcd resolver error:%v\n", err)
		return nil
	}
	client, err := chooseservice.NewClient(
		conf.GlobalConfig.DSN.ChooseServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		Log.Panicf("new client error:%v\n", err)
		return nil
	}
	return client
}

func init() {
	if ChooseClient == nil {
		ChooseClient = newChooseClient()
	}
}
