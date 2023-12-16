package base

import (
	"context"
	"sign/dao/db/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"time"
)

// Join implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Join(ctx context.Context, req *base.JoinReq) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	num, err := s.db.Group.JoinGroup(req.Uid, req.Gid)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		Log.Errorf("s.db.Group.JoinGroup error:%v", err)
		return
	}
	if num <= 0 {
		resp.Code = errmsg.UserAlreadyInGroup
		resp.Msg = errmsg.GetErrMsg(errmsg.UserAlreadyInGroup)
		return
	}
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}

// CreateGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) CreateGroup(ctx context.Context, req *base.GroupInfo) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	sign_in, _ := time.Parse("2006-01-02 15:04:05", req.SignIn)
	sign_out, _ := time.Parse("2006-01-02 15:04:05", req.SignOut)
	info := &model.Group{
		Name:     req.Name,
		Owner:    req.Owner,
		Places:   req.Places,
		Sign_in:  sign_in,
		Sign_out: sign_out,
	}
	err = s.db.Group.CreateGroup(info)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		Log.Errorf("s.db.Group.CreateGroup error:%v\n", err)
		return
	}
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}
