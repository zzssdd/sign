package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"sign/dao/mq/model"
	. "sign/pkg/log"
)

func (r *RabbitConn) PublishChooseMsg(signMsg *model.Choose) error {
	q, err := r.Ch.QueueDeclare("choose", true, false, false, false, nil)
	if err != nil {
		Log.Errorf("publish sign msg error:%v\n", err)
		return err
	}
	marshal, err := json.Marshal(signMsg)
	if err != nil {
		Log.Errorf("marshal data error:%v\n", err)
		return err
	}
	msg := amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: 1,
		Body:         marshal,
	}
	err = r.Ch.Publish(
		"",
		q.Name,
		false,
		false,
		msg,
	)
	if err != nil {
		Log.Errorf("publish msg error:%v\n", err)
		return err
	}
	return nil
}

func (r *RabbitConn) ConsumeChooseMsg() amqp.Delivery {
	q, err := r.Ch.QueueDeclare("choose", true, false, false, false, nil)
	if err != nil {
		Log.Errorf("publish sign msg error:%v\n", err)
		return amqp.Delivery{}
	}
	msg, _, err := r.Ch.Get(q.Name, true)
	if err != nil {
		Log.Errorf("consume msg error:%v\n", err)
		return amqp.Delivery{}
	}
	return msg
}
