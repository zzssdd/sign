package base

import (
	"context"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"sign/utils"
)

// Register implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Register(ctx context.Context, req *base.UserInfo) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	var rollAffect int64
	rollAffect, err = s.db.User.Register(req.Email, req.Password)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		Log.Errorf("s.db.User.Register error:%v\n", err)
		return
	}
	if rollAffect == 0 {
		resp.Code = errmsg.EmialExist
		resp.Msg = errmsg.GetErrMsg(errmsg.EmialExist)
		return
	}
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}

// Login implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Login(ctx context.Context, req *base.UserInfo) (resp *base.LoginResp, err error) {
	resp = new(base.LoginResp)
	id, err := s.db.User.Login(req.Email, req.Password)
	if err != nil {
		resp.Base.Code = errmsg.ERROR
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		Log.Errorf("s.db.User.Login error:%v", err)
		return
	}
	if id == -1 {
		resp.Base.Code = errmsg.EmailOrPasswordError
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.EmailOrPasswordError)
		return
	}
	token, err := utils.GenToken(req.Email, id, s.conf.JwtSecret)
	if err != nil {
		resp.Base.Code = errmsg.ERROR
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		Log.Errorf("utils.GenToken error:%v\n", err)
		return
	}
	resp.Base.Code = errmsg.SUCCESS
	resp.Base.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	resp.Token = &token
	return
}
