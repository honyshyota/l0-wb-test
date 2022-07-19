package cache

import (
	"sync"

	"github.com/honyshyota/l0-wb-test/internal/model"
)

type cache struct {
	sync.Mutex
	cache    map[int]*model.Order
	badCache map[int][]byte
}

func NewCache() Cache {
	return &cache{
		cache:    make(map[int]*model.Order),
		badCache: make(map[int][]byte),
	}
}

func (c *cache) Set(id int, value *model.Order) {
	c.Lock()
	c.cache[id] = value
	c.Unlock()
}

func (c *cache) SetBadCache(id int, value []byte) {
	c.Lock()
	c.badCache[id] = value
	c.Unlock()
}

func (c *cache) Get(id int) *model.Order {
	return c.cache[id]
}

func (c *cache) GetBadCache(id int) []byte {
	return c.badCache[id]
}
