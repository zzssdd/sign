package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"sign/dao/mq/model"
	. "sign/pkg/log"
)

func (r *RabbitConn) PublishSignMsg(signMsg *model.Sign) error {
	q, err := r.Ch.QueueDeclare("sign", true, false, false, false, nil)
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

func (r *RabbitConn) ConsumeSignMsg() <-chan amqp.Delivery {
	q, err := r.Ch.QueueDeclare("sign", true, false, false, false, nil)
	if err != nil {
		Log.Errorf("publish sign msg error:%v\n", err)
		return nil
	}
	msgChan, err := r.Ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		Log.Errorf("consume msg error:%v\n", err)
		return nil
	}
	return msgChan
}
