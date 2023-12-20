package base

import (
	"context"
	"math/rand"
	model2 "sign/dao/cache/model"
	model3 "sign/dao/db/model"
	"sign/dao/mq"
	"sign/dao/mq/model"
	base "sign/kitex_gen/sign/base"
	"sign/pkg/errmsg"
	. "sign/pkg/log"
	"sign/pkg/state"
	"sign/utils"
	"sync"
	"time"
)

// Choose implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Choose(ctx context.Context, req *base.ChooseReq) (resp *base.ChooseResp, err error) {
	resp = new(base.ChooseResp)
	id := s.db.Choose.GenID()
	info := &model2.Order{
		Uid: req.GetUid(),
		Aid: req.GetId(),
	}
	err = s.cache.CreateOrder(id, info)
	if err != nil {
		Log.Errorf("create order error:%v\n", err)
		resp.Base.Code = errmsg.ERROR
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
		return
	}
	resp.Base.Code = errmsg.SUCCESS
	resp.Base.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	resp.Id = &id
	return
}

// ChooseSubmit implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChooseSubmit(ctx context.Context, req *base.ChooseSubmitReq) (resp *base.ChooseSubmitResp, err error) {
	resp = new(base.ChooseSubmitResp)
	var ok bool
	if ok, err = s.cache.ExistSignAndExpireOrder(req.GetId()); !ok || err != nil {
		resp.Base.Code = errmsg.OrderExpired
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.OrderExpired)
		return
	}
	var order *model2.Order
	order, err = s.cache.GetOrder(req.GetId())
	if err != nil {
		resp.Base.Code = errmsg.OrderExpired
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.OrderExpired)
		return
	}
	if order.Status != state.OrderCreated {
		return
	}
	var score int64
	var activity *model3.Activity
	var groups string
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		var ok bool
		ok, err = s.cache.ExistAndExpireUserScore(req.GetUid())
		if err != nil {
			Log.Errorf("exist user score error,err:%v\n", err)
			return
		}
		if !ok {
			score, err = s.db.User.GetUserScore(req.GetUid())
			if err != nil {
				Log.Errorf("get user score from db error,err:%v\n", err)
				return
			}
		} else {
			score, err = s.cache.GetUserScore(req.GetUid())
			if err != nil {
				Log.Errorf("get user score from cache error,err:%v\n", err)
				return
			}
		}
		wg.Done()
	}()
	go func() {
		var ok bool
		ok, err = s.cache.ExistAndExpireActivity(order.Aid)
		if err != nil {
			Log.Errorf("exist activity error:%v\n", err)
			return
		}
		if !ok {
			activity, err = s.db.GetActivity(order.Aid)
			if err != nil {
				Log.Errorf("get activity info from db error:%v\n", err)
				return
			}
			info := &model2.Activity{
				Gid:        activity.Gid,
				Start_time: activity.Start_time.Format("2006-01-02 15:04:05"),
				End_time:   activity.End_time.Format("2006-01-02 15:04:05"),
				Prizes:     activity.Prizes,
				Cost:       activity.Cost,
			}
			_ = s.cache.StoreActivity(order.Aid, info)
			return
		} else {
			var info *model2.Activity
			info, err = s.cache.GetActivity(order.Aid)
			if err != nil {
				Log.Errorf("get activity from cache error:%v\n", err)
				return
			}
			start_time, _ := time.Parse("2006-01-02 15:04:05", info.Start_time)
			end_time, _ := time.Parse("2006-01-02 15:04:05", info.End_time)
			activity = &model3.Activity{
				Gid:        info.Gid,
				Start_time: start_time,
				End_time:   end_time,
				Prizes:     "",
				Cost:       0,
			}
		}
		wg.Done()
	}()
	go func() {
		var ok bool
		ok, err = s.cache.ExistAndExpireUserGroups(req.GetUid())
		if err != nil {
			return
		}
		if !ok {
			groups, err = s.db.Group.GetUserGroups(req.GetUid())
			_ = s.cache.StoreUserGroupsInfo(req.GetUid(), groups)
		} else {
			groups, err = s.cache.GetUserGroupsInfo(req.GetUid())
		}

	}()
	wg.Wait()
	now := time.Now()
	if !utils.StringContainInt64(groups, activity.Gid) {
		resp.Base.Code = errmsg.UserNotInGroup
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.UserNotInGroup)
		return
	}
	if now.Before(activity.Start_time) || now.After(activity.End_time) {
		resp.Base.Code = errmsg.NotInTime
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.NotInTime)
		return
	}
	if score < activity.Cost {
		resp.Base.Code = errmsg.ScoreNotEniough
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.ScoreNotEniough)
		return
	}
	prizes := utils.ParsePrizeString(activity.Prizes)
	var num int64
	for _, v := range prizes {
		num += v.Num
	}
	if num <= 0 {
		resp.Base.Code = errmsg.Thanks
		resp.Base.Msg = errmsg.GetErrMsg(errmsg.Thanks)
		return
	}
	wg.Add(1)
	var prizeInfo *base.PrizeInfo
	go func() {
		randNum := rand.Int63n(num)
		var pid int64
		for _, v := range prizes {
			randNum -= v.Num
			if randNum <= 0 {
				pid = v.Pid
				v.Num--
				break
			}
		}
		prizeString := utils.PackPrizeString(prizes)
		_ = s.cache.UpdatePrizes(order.Aid, prizeString)
		msg := &model.Choose{
			Uid: req.GetUid(),
			Id:  req.GetId(),
			Pid: pid,
		}
		chooseMq := mq.NewRabbitConn(s.conf)
		defer chooseMq.Close()
		err = chooseMq.PublishChooseMsg(msg)
		if err != nil {
			resp.Base.Code = errmsg.ERROR
			resp.Base.Msg = errmsg.GetErrMsg(errmsg.ERROR)
			return
		}
		for {
			_, _ = s.cache.ExistSignAndExpireOrder(req.GetId())
			order, _ = s.cache.GetOrder(req.GetId())
			if order.Status == state.OrderFinished {
				break
			} else {
				time.Sleep(time.Second)
			}
		}
		if pid == -1 {
			prizeInfo = &base.PrizeInfo{
				Id:   &pid,
				Name: "很遗憾，您未能中奖",
				Gid:  -1,
			}
		} else {
			if ok, err := s.cache.Prize.ExistAndExpirePrize(order.Pid); ok && err == nil {
				prize, err := s.cache.GetPrize(order.Pid)
				if err != nil {
					Log.Errorf("get prize from cache error:%v\n", err)
					return
				}
				id := order.Pid
				prizeInfo = &base.PrizeInfo{
					Id:   &id,
					Name: prize.Name,
					Gid:  prize.Gid,
				}
			} else {
				prize, err := s.db.Prize.GetPrize(order.Pid)
				if err != nil {
					Log.Errorf("get prize from db error,err:%v\n", err)
					return
				}
				prizeInfo = &base.PrizeInfo{
					Id:   &pid,
					Name: prize.Name,
					Gid:  prize.Gid,
				}
				info := &model2.Prize{
					Name: prize.Name,
					Gid:  prize.Gid,
				}
				_ = s.cache.StorePrize(pid, info)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	resp.Base.Code = errmsg.SUCCESS
	resp.Base.Msg = errmsg.GetErrMsg(errmsg.SUCCESS)
	resp.Info = prizeInfo
	return
}
