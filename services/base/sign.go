package base

import (
	"context"
	"sign/dao/mq/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"time"
)

// Sign implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Sign(ctx context.Context, req *base.SignReq) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	var signInTime time.Time
	var signOutTime time.Time
	if req.SigninTime != nil {
		signInTime, _ = time.Parse("2006-01-02 15:04:05", *req.SigninTime)
	}
	if req.SignoutTime != nil {
		signOutTime, _ = time.Parse("2006-01-02 15:04:05", *req.SignoutTime)
	}
	var place string
	if req.Place != nil {
		place = *req.Place
	}
	msg := &model.Sign{
		Uid:         req.Uid,
		Gid:         req.Gid,
		SignInTime:  signInTime,
		SignOutTime: signOutTime,
		Place:       place,
		PublishTime: time.Now(),
	}
	err = s.mq.PublishSignMsg(msg)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		return
	}
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}

// SignMonth implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) SignMonth(ctx context.Context, req *base.MonthSignReq) (resp *base.MonthSignResp, err error) {
	resp = new(base.MonthSignResp)
	month, err := s.db.Sign.GetSignMonth(req.Gid, req.Uid)
	if err != nil {
		Log.Errorf("s.db.Sign.GetSignMonth error:%v\n", err)
		resp.Base.Code = errmsg.ERROR
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		return
	}
	resp.Base.Code = errmsg.SUCCESS
	resp.Base.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	resp.Data = &month
	return
}
