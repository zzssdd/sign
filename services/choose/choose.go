package choose

import (
	"context"
	"encoding/json"
	"fmt"
	"sign/dao/db/model"
	model2 "sign/dao/mq/model"
	"sign/kitex_gen/sign/choose"
	. "sign/pkg/log"
	"sign/pkg/state"
	"time"
)

// Choose implements the ChooseServiceImpl interface.
func (s *ChooseServiceImpl) Choose(ctx context.Context, req *choose.Empty) (resp *choose.Empty, err error) {
	resp = new(choose.Empty)
	msgChan := s.mq.ConsumeChooseMsg()
	var chooseInfo *model2.Choose
	for msg := range msgChan {
		err = json.Unmarshal(msg.Body, &chooseInfo)
		if err != nil {
			Log.Errorf("unmarshal choose msg error:%v\n", err)
			continue
		}
		err = s.HandleChoose(chooseInfo)
		if err != nil {
			_ = msg.Ack(false)
		} else {
			_ = msg.Ack(true)
		}
	}
	return
}

func (s *ChooseServiceImpl) HandleChoose(chooseInfo *model2.Choose) error {
	if ok, err := s.cache.ExistSignAndExpireOrder(chooseInfo.Id); !ok || err == nil {
		return fmt.Errorf("order not exist")
	}
	order, err := s.cache.GetOrder(chooseInfo.Id)
	if err != nil {
		Log.Errorf("get order from cache error:%v\n", err)
		return err
	}
	if order.Status != state.OrderDoing {
		return fmt.Errorf("order status not correct")
	}
	var activity *model.Activity
	var ok bool
	activity, err = s.db.GetActivity(order.Aid)
	if err != nil {
		Log.Errorf("get group from db error,err:%v\n", err)
		return err
	}
	ok, err = s.db.Activity.CheckoutAndUpdateTryNum(order.Aid, chooseInfo.Pid)
	if err != nil {
		_, _ = s.db.Activity.CancelNum(order.Aid, chooseInfo.Pid)
		Log.Errorf("s.db.Activity.CheckoutAndUpdateTryNum error,err:%v\n", err)
		return err
	}
	if !ok {
		chooseInfo.Pid = -1
	}
	ok, err = s.db.User.CheckoutAndUpdateTryScore(chooseInfo.Uid, activity.Cost)
	if err != nil {
		_, _ = s.db.Activity.CancelNum(order.Aid, chooseInfo.Pid)
		_ = s.db.User.CancelScore(chooseInfo.Uid, activity.Cost)
		Log.Errorf("get s.db.User.CheckoutAndUpdateTryScore error:%v\n", err)
		return err
	}
	if !ok {
		_, _ = s.db.Activity.CancelNum(order.Aid, chooseInfo.Pid)
		_ = s.db.User.CancelScore(chooseInfo.Uid, activity.Cost)
		Log.Errorf("score not enought\n")
		return nil
	}
	_ = s.db.Activity.CommitNum(order.Aid, chooseInfo.Pid)
	_ = s.db.User.CommitScore(chooseInfo.Uid, activity.Cost)

	_ = s.cache.DeleteOrderPrize(chooseInfo.Id)
	_ = s.cache.DeleteUserScore(chooseInfo.Uid)

	record := &model.Record{
		Uid:     chooseInfo.Uid,
		Pid:     chooseInfo.Pid,
		GetTime: time.Now(),
	}
	err = s.db.Choose.CreateOrder(record)
	if err != nil {
		Log.Errorf("create order error")
		return err
	}

	err = s.cache.UpdateOrderPrize(chooseInfo.Id, chooseInfo.Pid)
	if err != nil {
		Log.Errorf("update order prize error:%v\n", err)
		return err
	}
	err = s.cache.UpdateOrder(chooseInfo.Id, state.OrderFinished)
	if err != nil {
		Log.Errorf("update order state error:%v\n", err)
		return err
	}
	return nil
}
