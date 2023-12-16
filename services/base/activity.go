package base

import (
	"context"
	"sign/dao/db/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"time"
)

// AddActivity implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) AddActivity(ctx context.Context, req *base.ActicityInfo) (resp *base.BaseResp, err error) {
	start_time, _ := time.Parse("2006-01-02 15:04:05", req.StartTime)
	end_time, _ := time.Parse("2006-01-02 15:04:05", req.EndTime)
	info := &model.Activity{
		Gid:        req.Gid,
		Start_time: start_time,
		End_time:   end_time,
		Prizes:     req.Prizes,
		Cost:       req.Cost,
	}
	err = s.db.Activity.CreateActivity(info)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		Log.Errorf("s.db.Activity.CreateActivity error:%v\n", err)
		return
	}
	resp.Code = errmsg.SUCCESS
	resp.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	return
}
