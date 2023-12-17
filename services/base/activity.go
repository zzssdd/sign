package base

import (
	"context"
	model2 "sign/dao/cache/model"
	"sign/dao/db/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"sign/utils"
	"time"
)

// AddActivity implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) AddActivity(ctx context.Context, req *base.ActicityInfo) (resp *base.BaseResp, err error) {
	start_time, _ := time.Parse("2006-01-02 15:04:05", req.GetStartTime())
	end_time, _ := time.Parse("2006-01-02 15:04:05", req.GetEndTime())
	info := &model.Activity{
		Gid:        req.GetGid(),
		Start_time: start_time,
		End_time:   end_time,
		Prizes:     req.GetPrizes(),
		Cost:       req.GetCost(),
	}
	var id int64
	id, err = s.db.Activity.CreateActivity(info)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		Log.Errorf("s.db.Activity.CreateActivity error:%v\n", err)
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
	go func() {
		info := &model2.Activity{
			Gid:        req.GetGid(),
			Start_time: req.GetStartTime(),
			End_time:   req.GetEndTime(),
			Prizes:     "",
			Cost:       0,
		}
		_ = s.cache.StoreActivity(id, info)
	}()
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}
