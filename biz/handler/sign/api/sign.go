package api

import (
	"context"
	"sign/biz/handler/sign/common"
	"sign/biz/model/sign/api"
	"sign/biz/rpc"
	base2 "sign/kitex_gen/sign/base"
	. "sign/pkg/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Sign .
// @router /join [POST]
func Sign(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SignReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base2.BaseResp)
	rpcReq := new(base2.SignReq)
	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	rpcReq.Uid = c.GetInt64("id")
	resp, err = rpc.BaseClient.Sign(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.Sign %v\n", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// SignMonth .
// @router /signMonth [GET]
func SignMonth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.MonthSignReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base2.MonthSignResp)
	rpcReq := new(base2.MonthSignReq)
	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	rpcReq.Uid = c.GetInt64("id")
	resp, err = rpc.BaseClient.SignMonth(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.SignMonth error:%v\n", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
