package api

import (
	"context"
	"sign/biz/handler/sign/common"
	base "sign/biz/model/sign/base"
	"sign/biz/rpc"
	base2 "sign/kitex_gen/sign/base"
	. "sign/pkg/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AddPrize .
// @router /prize [POST]
func AddPrize(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.PrizeInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base2.BaseResp)
	rpcReq := new(base2.PrizeInfo)

	err = common.BindRPC(req, rpcReq)
	if err != nil {
		Log.Errorf("bind rpcReq error %v\n", err)
		return
	}

	resp, err = rpc.BaseClient.AddPrize(ctx, rpcReq)
	if err != nil {
		Log.Errorf("rpc.BaseClient.AddPrize error:%v\n", err)
		return
	}
	c.JSON(consts.StatusOK, resp)
}
