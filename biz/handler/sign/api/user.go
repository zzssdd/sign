package api

import (
	"context"
	"sign/biz/handler/sign/common"
	"sign/biz/model/sign/base"
	"sign/biz/rpc"
	base2 "sign/kitex_gen/sign/base"
	. "sign/pkg/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.UserInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	rpcReq := new(base2.UserInfo)
	resp := new(base2.BaseResp)

	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	resp, err = rpc.BaseClient.Register(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.Register error:%v\n", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.UserInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base2.LoginResp)
	rpcReq := new(base2.UserInfo)

	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	resp, err = rpc.BaseClient.Login(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.Login error:%v\n", err)
	}
	c.JSON(consts.StatusOK, resp)
}
