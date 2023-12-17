package base

import (
	"context"
	"sign/dao/cache/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"sign/utils"
)

// Register implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Register(ctx context.Context, req *base.UserInfo) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	var rollAffect int64
	var id int64
	rollAffect, id, err = s.db.User.Register(req.GetEmail(), req.GetPassword())
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
	go func() {
		info := &model.User{
			Email:    req.GetEmail(),
			Id:       id,
			Password: req.GetPassword(),
		}
		err2 := s.cache.User.StoreUser(info)
		if err2 != nil {
			Log.Errorf("store userInfo into cache error:%v\n", err2)
		}
	}()
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}

// Login implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Login(ctx context.Context, req *base.UserInfo) (resp *base.LoginResp, err error) {
	resp = new(base.LoginResp)
	var id int64
	if ok, err2 := s.cache.User.ExistAndExpireUser(req.GetEmail()); err2 == nil && ok {
		ok, id, err2 = s.cache.User.CheckLogin(req.GetEmail(), req.GetPassword())
		if !ok {
			resp.Base.Code = errmsg.EmailOrPasswordError
			resp.Base.Msg = errmsg.GetErrMsg(errmsg.EmailOrPasswordError)
			return
		}
	} else {
		var password string
		id, password, err = s.db.User.Login(req.GetEmail())
		go func() {
			info := &model.User{
				Email:    req.GetEmail(),
				Id:       id,
				Password: password,
			}
			err2 := s.cache.User.StoreUser(info)
			if err2 != nil {
				Log.Errorf("store userInfo into cache error:%v\n", err2)
			}
		}()
		if password != req.Password {
			resp.Base.Code = errmsg.EmailOrPasswordError
			resp.Base.Msg = errmsg.GetErrMsg(errmsg.EmailOrPasswordError)
			return
		}
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
