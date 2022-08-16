package cache

import (
	"github.com/dingowd/WB/L0/internal/logger"
	storage2 "github.com/dingowd/WB/L0/internal/storage"
	"github.com/dingowd/WB/L0/model"
	"time"
)

type CacheInterface interface {
	Init()
	ReadFromCache(id string) (*model.CacheOrder, error)
	WriteToCache(o model.Order) int
}

type Cache struct {
	Log    logger.Logger
	Stor   storage2.Storage
	Amount int
	Body   model.CacheOrderList
}

func NewCache(log logger.Logger, stor storage2.Storage, a int) *Cache {
	return &Cache{
		Log:    log,
		Stor:   stor,
		Amount: a,
		Body:   make(model.CacheOrderList, 0),
	}
}

func (c *Cache) Init() {
	c.Body, _ = c.Stor.GetOrdersByLimit(c.Amount)
}

func (c *Cache) ReadFromCache(id string) (*model.CacheOrder, error) {
	for i, val := range c.Body {
		if id == val.Order.OrderUid {
			c.Body[i].TimeStamp = time.Now().UnixNano()
			return &c.Body[i], nil
		}
	}
	if !c.Stor.IsOrderExist(id) {
		return nil, storage2.ErrorOrderNotExist
	}
	o, err := c.Stor.GetOrder(id)
	if err != nil {
		return nil, err
	}
	k := c.WriteToCache(o)
	return &c.Body[k], nil
}

func (c *Cache) WriteToCache(o model.Order) int {
	if len(c.Body) < c.Amount {
		var b model.CacheOrder
		b.Order = o
		b.TimeStamp = time.Now().UnixNano()
		c.Body = append(c.Body, b)
		return len(c.Body) - 1
	}
	key := 0
	min := c.Body[0].TimeStamp
	for i := 1; i < len(c.Body); i++ {
		if c.Body[i].TimeStamp < min {
			min = c.Body[i].TimeStamp
			key = i
		}
	}
	c.Body[key].Order = o
	c.Body[key].TimeStamp = time.Now().UnixNano()
	return key
}
