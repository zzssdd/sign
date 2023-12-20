package base

import (
	"context"
	model2 "sign/dao/cache/model"
	"sign/dao/db/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"sign/utils"
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
				SignIn:  group.Sign_in,
				SignOut: group.Sign_out,
				Count:   0,
			}
			err = s.cache.Group.StoreGroup(req.GetGid(), info)
			if err != nil {
				Log.Errorf("s.cache.Group.StoreGroup error:%v\n", err)
			}
			places := utils.ParsePlacesString(group.Places)
			for _, p := range places {
				pos := &model2.SignPos{
					Gid:        req.GetGid(),
					Name:       p.Name,
					Latitle:    p.Latitude,
					Longtitude: p.Longtitude,
				}
				err = s.cache.AddSignPos(pos)
				if err != nil {
					Log.Errorf("s.cache.AddSignPos error:%v\n", err)
				}
			}
		}
		_ = s.cache.IncrGroupCount(req.GetGid())
	}()
	go func() {
		var ok bool
		ok, err = s.cache.ExistAndExpireUserGroups(req.GetUid())
		if err != nil {
			Log.Errorf("expire user groups error:%v\n", err)
			return
		}
		if ok {
			_ = s.cache.DelUserGroupsInfo(req.Gid)
		}
	}()
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}

// CreateGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) CreateGroup(ctx context.Context, req *base.GroupInfo) (resp *base.BaseResp, err error) {
	resp = new(base.BaseResp)
	info := &model.Group{
		Name:     req.GetName(),
		Owner:    req.GetOwner(),
		Places:   req.GetPlaces(),
		Sign_in:  req.GetSignIn(),
		Sign_out: req.GetSignOut(),
		Score:    req.GetScore(),
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
			Score:   req.GetScore(),
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
