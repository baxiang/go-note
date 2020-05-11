package main

import (
	"fmt"
	"container/list"
	"sync"
)

type Node struct {
	key   int
	value int
}
type LRUcache struct {
	mt   sync.Mutex
	m    map[int]*list.Element
	c    int
	list *list.List
}

func NewLRUcache(c int) *LRUcache {
	return &LRUcache{
		m:    make(map[int]*list.Element),
		c:    c,
		list: list.New(),
	}
}

func (l *LRUcache) Get(key int) int {
	l.mt.Lock()
	defer l.mt.Unlock()
	if v, ok := l.m[key]; ok {
		l.list.MoveToBack(v)
		n := v.Value.(*Node)
		return n.value
	}
	return -1
}

func (l *LRUcache) Put(key, value int) {
	l.mt.Lock()
	defer l.mt.Unlock()
	if v, ok := l.m[key]; ok {
		n := v.Value.(*Node)
		n.value = value
		l.list.MoveToFront(v)
	} else {
		if l.c == len(l.m) {
			e := l.list.Back()
			n := e.Value.(*Node)
			delete(l.m, n.key)
			l.list.Remove(e)
		}
		n := &Node{key, value}
		e := l.list.PushFront(n)
		l.m[key] = e

	}
}

func main() {
	cache := NewLRUcache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	cache.Put(4, 4)
}
