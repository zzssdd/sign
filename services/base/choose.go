package base

import (
	"context"
	"sign/dao/mq/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	"time"
)

// Choose implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Choose(ctx context.Context, req *base.ChooseReq) (resp *base.ChooseResp, err error) {
	resp = new(base.ChooseResp)
	id := s.db.Choose.GenID()
	resp.Base.Code = errmsg.SUCCESS
	resp.Base.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	resp.Id = &id
	return
}

// ChooseSubmit implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChooseSubmit(ctx context.Context, req *base.ChooseSubmitReq) (resp *base.ChooseSubmitResp, err error) {
	resp = new(base.ChooseSubmitResp)
	msg := &model.Choose{
		Uid:         req.Uid,
		Id:          req.Id,
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
