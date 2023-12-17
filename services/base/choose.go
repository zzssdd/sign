package base

import (
	"context"
	model2 "sign/dao/cache/model"
	"sign/dao/mq/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"time"
)

// Choose implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Choose(ctx context.Context, req *base.ChooseReq) (resp *base.ChooseResp, err error) {
	resp = new(base.ChooseResp)
	id := s.db.Choose.GenID()
	info := &model2.Order{
		Uid: req.GetUid(),
		Pid: req.GetId(),
	}
	err = s.cache.CreateOrder(id, info)
	if err != nil {
		Log.Errorf("create order error:%v\n", err)
		resp.Base.Code = errmsg.ERROR
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		return
	}
	resp.Base.Code = errmsg.SUCCESS
	resp.Base.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	resp.Id = &id
	return
}

// ChooseSubmit implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChooseSubmit(ctx context.Context, req *base.ChooseSubmitReq) (resp *base.ChooseSubmitResp, err error) {
	resp = new(base.ChooseSubmitResp)
	msg := &model.Choose{
		Uid:         req.GetUid(),
		Id:          req.GetId(),
		PublishTime: time.Now(),
	}
	err = s.mq.PublishChooseMsg(msg)
	if err != nil {
		resp.Base.Code = errmsg.ERROR
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		return
	}
	resp.Base.Code = errmsg.SUCCESS
	resp.Base.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}
