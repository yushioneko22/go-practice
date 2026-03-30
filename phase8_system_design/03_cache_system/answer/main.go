package main

import (
	"fmt"
	"time"
)

// === LRU Cache ===

type Node struct {
	key, val   int
	prev, next *Node
	expiresAt  time.Time
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func NewLRUCache(capacity int) *LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return &LRUCache{capacity: capacity, cache: make(map[int]*Node), head: head, tail: tail}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToFront(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if node, ok := c.cache[key]; ok {
		node.val = value
		c.moveToFront(node)
		return
	}
	node := &Node{key: key, val: value}
	c.cache[key] = node
	c.addToFront(node)
	if len(c.cache) > c.capacity {
		lru := c.tail.prev
		c.removeNode(lru)
		delete(c.cache, lru.key)
	}
}

func (c *LRUCache) addToFront(n *Node)  { n.prev = c.head; n.next = c.head.next; c.head.next.prev = n; c.head.next = n }
func (c *LRUCache) removeNode(n *Node)  { n.prev.next = n.next; n.next.prev = n.prev }
func (c *LRUCache) moveToFront(n *Node) { c.removeNode(n); c.addToFront(n) }

// === TTL Cache ===

type TTLCache struct {
	lru *LRUCache
	ttl time.Duration
}

func NewTTLCache(capacity int, ttl time.Duration) *TTLCache {
	return &TTLCache{lru: NewLRUCache(capacity), ttl: ttl}
}

func (c *TTLCache) Get(key int) int {
	if node, ok := c.lru.cache[key]; ok {
		if time.Now().After(node.expiresAt) {
			c.lru.removeNode(node)
			delete(c.lru.cache, key)
			return -1
		}
		c.lru.moveToFront(node)
		return node.val
	}
	return -1
}

func (c *TTLCache) Put(key, value int) {
	c.lru.Put(key, value)
	c.lru.cache[key].expiresAt = time.Now().Add(c.ttl)
}

func main() {
	fmt.Println("=== 演習1: LRU Cache ===")
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Printf("[LRU] Get(1): %d\n", cache.Get(1))
	cache.Put(3, 3)
	fmt.Printf("[LRU] Get(2): %d (削除済み)\n", cache.Get(2))
	fmt.Printf("[LRU] Get(3): %d\n", cache.Get(3))

	fmt.Println("\n=== 演習2: TTL付きキャッシュ ===")
	ttlCache := NewTTLCache(10, 2*time.Second)
	ttlCache.Put(1, 100)
	fmt.Printf("[TTL] Get(1): %d\n", ttlCache.Get(1))
	fmt.Println("[TTL] 3秒待機...")
	time.Sleep(3 * time.Second)
	fmt.Printf("[TTL] Get(1): %d (期限切れ)\n", ttlCache.Get(1))
}
