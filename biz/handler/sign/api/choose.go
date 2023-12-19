package api

import (
	"context"
	"sign/biz/handler/sign/common"
	"sign/biz/model/sign/api"
	"sign/biz/rpc"
	"sign/kitex_gen/sign/base"
	. "sign/pkg/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Choose .
// @router /choose [GET]
func Choose(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ChooseReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.ChooseResp)
	rpcReq := new(base.ChooseReq)
	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	rpcReq.Uid = c.GetInt64("id")
	resp, err = rpc.BaseClient.Choose(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.Choose error:%v\n", err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// ChooseSubmit .
// @router /choose [POST]
func ChooseSubmit(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ChooseSubmitReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.ChooseSubmitResp)
	rpcReq := new(base.ChooseSubmitReq)
	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	rpcReq.Uid = c.GetInt64("id")
	resp, err = rpc.BaseClient.ChooseSubmit(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.ChooseSubmit error:%s\n", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
