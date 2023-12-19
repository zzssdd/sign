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

// Join .
// @router /join [POST]
func Join(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.JoinReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.BaseResp)
	rpcReq := new(base.JoinReq)
	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	rpcReq.Uid = c.GetInt64("id")
	resp, err = rpc.BaseClient.Join(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.Join error:%v\n", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// CreateGroup .
// @router /group [POST]
func CreateGroup(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GroupInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.BaseResp)
	rpcReq := new(base.GroupInfo)
	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	resp, err = rpc.BaseClient.CreateGroup(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.CreateGroup error:%v\n", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
