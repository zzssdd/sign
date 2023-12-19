package base

import (
	"context"
	model2 "sign/dao/cache/model"
	"sign/dao/mq/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"sign/utils"
	"time"
)

// Sign implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Sign(ctx context.Context, req *base.SignReq) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	var signInTime time.Time
	var signOutTime time.Time
	if req.SigninTime != nil {
		signInTime, _ = time.Parse("2006-01-02 15:04:05", req.GetSigninTime())
	}
	if req.SignoutTime != nil {
		signOutTime, _ = time.Parse("2006-01-02 15:04:05", req.GetSignoutTime())
	}
	go func() {
		if ok, err := s.cache.ExistAndExpireGroup(req.GetGid()); !ok || err != nil {
			group, err := s.db.GetGroup(req.GetGid())
			if err != nil {
				Log.Errorf("get group from db error,err:%v\n", err)
				return
			}
			info := &model2.Group{
				Name:    group.Name,
				Owner:   group.Owner,
				Places:  group.Places,
				SignIn:  group.Sign_in.Format("2006-01-02 15:04:05"),
				SignOut: group.Sign_out.Format("2006-01-02 15:04:05"),
				Count:   0,
			}
			_ = s.cache.Group.StoreGroup(req.GetGid(), info)
			places := utils.ParsePlacesString(group.Places)
			for _, p := range places {
				pos := &model2.SignPos{
					Gid:        req.GetGid(),
					Name:       p.Name,
					Latitle:    p.Latitude,
					Longtitude: p.Longtitude,
				}
				_ = s.cache.AddSignPos(pos)
			}
		}
	}()
	msg := &model.Sign{
		Uid:         req.GetUid(),
		Gid:         req.GetGid(),
		SignInTime:  signInTime,
		SignOutTime: signOutTime,
		Place:       req.GetPlace(),
		PublishTime: time.Now(),
		Flag:        req.GetFlag(),
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
	info := &model2.SignMonth{
		Uid:   req.GetUid(),
		Month: req.GetMonth(),
	}
	var bits int32
	var ok bool
	ok, err = s.cache.ExistAndExpireMonth(info)
	if err == nil && ok {
		bits, err = s.cache.GetSignMonth(info)
		if err != nil {
			Log.Errorf("s.cache.GetSignMonth error:%v\n", err)
			resp.Base.Code = errmsg.ERROR
			resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
			return
		}
	} else {
		bits, err = s.db.Sign.GetSignMonth(req.GetGid(), req.GetUid())
		if err != nil {
			Log.Errorf("s.db.Sign.GetSignMonth error:%v\n", err)
			resp.Base.Code = errmsg.ERROR
			resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
			return
		}
		go func() {
			_ = s.cache.SetSignMonth(info, bits)
		}()
	}
	resp.Base.Code = errmsg.SUCCESS
	resp.Base.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	resp.Data = &bits
	return
}
