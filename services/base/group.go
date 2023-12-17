package base

import (
	"context"
	model2 "sign/dao/cache/model"
	"sign/dao/db/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"sign/utils"
	"strconv"
	"time"
)

// Join implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Join(ctx context.Context, req *base.JoinReq) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	num, err := s.db.Group.JoinGroup(req.GetUid(), req.GetGid())
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
	go func() {
		info := &model2.UserGroups{
			Uid:    req.GetUid(),
			Groups: strconv.FormatInt(req.GetGid(), 10),
		}
		_ = s.cache.StoreUserGroupsInfo(req.GetUid(), info)
		_ = s.cache.IncrGroupCount(req.GetGid())
	}()
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}

// CreateGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) CreateGroup(ctx context.Context, req *base.GroupInfo) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	sign_in, _ := time.Parse("2006-01-02 15:04:05", req.GetSignIn())
	sign_out, _ := time.Parse("2006-01-02 15:04:05", req.GetSignOut())
	info := &model.Group{
		Name:     req.GetName(),
		Owner:    req.GetOwner(),
		Places:   req.GetPlaces(),
		Sign_in:  sign_in,
		Sign_out: sign_out,
	}
	var id int64
	id, err = s.db.Group.CreateGroup(info)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		Log.Errorf("s.db.Group.CreateGroup error:%v\n", err)
		return
	}
	go func() {
		info := &model2.Group{
			Name:    req.GetName(),
			Owner:   req.GetOwner(),
			Places:  req.GetPlaces(),
			SignIn:  req.GetSignIn(),
			SignOut: req.GetSignOut(),
			Count:   0,
		}
		_ = s.cache.Group.StoreGroup(id, info)
		places := utils.ParsePlacesString(req.GetPlaces())
		for _, p := range places {
			pos := &model2.SignPos{
				Gid:        id,
				Name:       p.Name,
				Latitle:    p.Latitude,
				Longtitude: p.Longtitude,
			}
			_ = s.cache.AddSignPos(pos)
		}
	}()
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}
