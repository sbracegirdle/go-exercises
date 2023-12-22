package main

import (
	"container/list"
	"errors"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type pair struct {
	key   int
	value int
}

func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be greater than 0")
	}
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}, nil
}

func (c *LRUCache) Get(key int) (int, error) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(pair).value, nil
	}
	return -1, errors.New("key not found")
}

func (c *LRUCache) Put(key int, value int) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value = pair{key, value}
	} else {
		if c.list.Len() >= c.capacity {
			delete(c.cache, c.list.Back().Value.(pair).key)
			c.list.Remove(c.list.Back())
		}
		c.cache[key] = c.list.PushFront(pair{key, value})
	}
}

func main() {
	fmt.Println("Hello, World!")
}
