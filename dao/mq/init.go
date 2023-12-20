package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"sign/conf"
	. "sign/pkg/log"
)

type RabbitConn struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewRabbitConn(conf *conf.Config) *RabbitConn {
	dsn := fmt.Sprintf("amqp://%s:%s@%s/%s", conf.DSN.UserNameRabbit, conf.DSN.PassWordRabbit, conf.DSN.RabbitDSN, conf.DSN.RabbitVhost)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		Log.Panicf("connect to rabbitmq error:%v", err)
		panic(err)
	}
	ch, err := conn.Channel()
	return &RabbitConn{
		Conn: conn,
		Ch:   ch,
	}
}

func (r *RabbitConn) Close() {
	r.Ch.Close()
	r.Conn.Close()
}
