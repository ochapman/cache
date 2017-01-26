/*
* File Name:	cache.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2017-01-18
 */

package cache

import (
	"container/list"
	"fmt"
	"time"
)

type entry struct {
	key        interface{}
	value      interface{}
	createTime time.Time
}
type Cache struct {
	maxEntries uint64
	ll         *list.List //list of entry
	expire     time.Duration
	m          map[interface{}]*list.Element //key to list.List
}

func New(maxEntries uint64, expire time.Duration) *Cache {
	return &Cache{
		maxEntries: maxEntries,
		ll:         list.New(),
		expire:     expire,
		m:          make(map[interface{}]*list.Element),
	}
}

func (c *Cache) Add(key interface{}, value interface{}) {
	e := entry{
		key:        key,
		value:      value,
		createTime: time.Now(),
	}

	if v, ok := c.m[key]; ok {
		//update value
		if v.Value.(*entry).value != value {
			v.Value.(*entry).value = value
		}
		//Move to front if exist
		c.ll.MoveToFront(v)
	} else {
		element := c.ll.PushFront(&e)
		c.m[key] = element
	}
}

func (c *Cache) Dump() {
	for en := c.ll.Front(); en != nil; en = en.Next() {
		fmt.Printf("%#v\n", en.Value)
	}
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	if v, ok := c.m[key]; ok {
		e := v.Value.(*entry)
		fmt.Printf("duration: %v, expired: %v\n", time.Since(e.createTime), c.expire)
		if c.expire == 0 {
			return e.value, true
		} else {
			if time.Since(e.createTime) < c.expire {
				return e.value, true
			}
		}
	}
	return nil, false
}
