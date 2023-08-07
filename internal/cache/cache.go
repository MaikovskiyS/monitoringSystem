package cache

import (
	"diploma/internal/transport/http/dto"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Cache struct {
	sync.RWMutex
	l                 *logrus.Logger
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	item              Item
}
type Item struct {
	Value      dto.ResultT
	Created    time.Time
	Expiration int64
}

func New(l *logrus.Logger) *Cache {
	return &Cache{
		l:                 l,
		defaultExpiration: 30 * time.Second,
		cleanupInterval:   30 * time.Second,
		item: Item{
			Expiration: 0,
		},
	}
}
func (c *Cache) Set(value dto.ResultT) {

	expiration := time.Now().Add(c.defaultExpiration).UnixNano()
	c.RLock()
	defer c.RUnlock()
	c.item = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}

}
func (c *Cache) Get() (dto.ResultT, bool) {
	c.RLock()
	defer c.RUnlock()

	item := c.item
	if time.Now().UnixNano() > item.Expiration {
		c.Delete()
		return dto.ResultT{}, false
	}

	return item.Value, true
}
func (c *Cache) Delete() error {
	c.RLock()
	defer c.RUnlock()
	c.item = Item{
		Expiration: 0,
	}

	return nil
}
