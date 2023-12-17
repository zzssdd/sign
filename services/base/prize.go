package base

import (
	"context"
	model2 "sign/dao/cache/model"
	"sign/dao/db/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
)

// AddPrize implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) AddPrize(ctx context.Context, req *base.PrizeInfo) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	info := &model.Prize{
		Name: req.GetName(),
		Gid:  req.GetGid(),
	}
	var id int64
	id, err = s.db.Prize.CreatePrize(info)
	if err != nil {
		Log.Errorf("add prize error:%v\n", err)
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		return
	}
	go func() {
		info := &model2.Prize{
			Name: req.GetName(),
			Gid:  req.GetGid(),
		}
		_ = s.cache.Prize.StorePrize(id, info)
	}()
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}
