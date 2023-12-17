package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sign/conf"
	"sign/dao/cache/model"
)

type Order struct {
	cahceTime int
}

func newOrder(config *conf.Cache) *Order {
	return &Order{
		cahceTime: config.Order,
	}
}

const (
	UserOrderKey = "order_%d"
)

func orderKey(id int64) string {
	return fmt.Sprintf(UserOrderKey, id)
}

func (o *Order) CreateOrder(id int64, order *model.Order) error {
	rds := CachePool.Get()
	defer rds.Close()
	order.Status = "Created"
	err := rds.Send("HMSET", redis.Args{}.Add(orderKey(id)).AddFlat(&order)...)
	if err != nil {
		return err
	}
	err = rds.Send("EXPIRE", orderKey(id), o.cahceTime)
	if err != nil {
		return err
	}
	return rds.Flush()
}

func (o *Order) ExistSignAndExpireOrder(id int64) (bool, error) {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := redis.Int(rds.Do("EXPIRE", orderKey(id), o.cahceTime))
	if err != nil {
		return false, err
	}
	return reply == 1, nil
}

func (o *Order) GetOrder(id int64) (*model.Order, error) {
	rds := CachePool.Get()
	defer rds.Close()
	order := new(model.Order)
	v, err := redis.Values(rds.Do("HGETALL", orderKey(id)))
	if err != nil {
		return nil, err
	}
	err = redis.ScanStruct(v, order)
	return order, err
}

func (o *Order) UpdateOrder(id int64, state string) error {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := rds.Do("HSET", orderKey(id), "State", state)
	if err != nil {
		return err
	}
	if reply != 1 {
		return fmt.Errorf("update order state error")
	}
	return nil
}

func (o *Order) UpdateOrderPrize(id int64, pid int64) error {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := rds.Do("HSET", orderKey(id), "Pid", pid)
	if err != nil {
		return err
	}
	if reply != 1 {
		return fmt.Errorf("update order pid error")
	}
	return nil
}

func (o *Order) DeleteOrder(id int64) error {
	rds := CachePool.Get()
	defer rds.Close()
	reply, err := rds.Do("DEL", orderKey(id))
	if err != nil {
		return err
	}
	if reply != 1 {
		return fmt.Errorf("delete order error")
	}
	return nil
}
