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

// AddActivity .
// @router /activity [POST]
func AddActivity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.ActicityInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base2.BaseResp)
	rpcReq := new(base2.ActicityInfo)

	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}
	resp, err = rpc.BaseClient.AddActivity(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.AddActivity err:%v\n", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
