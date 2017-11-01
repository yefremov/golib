package golib

import (
	"container/list"
	"sync"
)

// Cache ...
type Cache interface {
	Get(int64) interface{}
	Put(int64, interface{}) bool
	Reset()
}

// Options ...
type options struct {
	size int
}

// cache ...
type cache struct {
	options *options
	mutex   sync.Mutex
	list    *list.List
	items   map[int64]*list.Element
}

// record ...
type record struct {
	key   int64
	value interface{}
}

// NewCache ...
func NewCache(size int) Cache {
	c := &cache{
		options: &options{
			size: size,
		},
	}
	c.Reset()
	return c
}

// Get ...
func (c *cache) Get(key int64) interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e, ok := c.items[key]; ok {
		c.list.MoveToFront(e)
		return e.Value.(record).value
	}

	return nil
}

// Put ...
func (c *cache) Put(key int64, value interface{}) bool {
	if value == nil {
		return false
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e, ok := c.items[key]; ok {
		c.list.MoveToFront(e)
		return true
	}

	item := c.list.PushFront(record{
		key:   key,
		value: value,
	})
	c.items[key] = item

	if c.list.Len() > c.options.size {
		e := c.list.Back()
		if e != nil {
			c.list.Remove(e)
			k := e.Value.(record).key
			delete(c.items, k)
		}
	}
	return true
}

// Reset ...
func (c *cache) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = make(map[int64]*list.Element)
	c.list = list.New()
}
