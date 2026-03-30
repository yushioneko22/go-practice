package main

import "fmt"

// Node は双方向連結リストのノード
type Node struct {
	key, val   int
	prev, next *Node
}

// LRUCache はハッシュマップ + 双方向連結リスト
type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node // ダミーノード
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// Get はキャッシュから値を取得（ノードを先頭に移動）
func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToFront(node)
		return node.val
	}
	return -1
}

// Put はキャッシュに値を追加/更新（容量超過時は末尾を削除）
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

func (c *LRUCache) addToFront(node *Node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) moveToFront(node *Node) {
	c.removeNode(node)
	c.addToFront(node)
}

func main() {
	fmt.Println("=== [LRUCache] LRU Cache ===")
	cache := NewLRUCache(3)

	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300)
	fmt.Printf("[LRUCache] Get(1): %d\n", cache.Get(1)) // 100
	fmt.Printf("[LRUCache] Get(2): %d\n", cache.Get(2)) // 200

	// 容量超過 → key=3が一番古いので削除される（key=1は先にアクセスした）
	cache.Put(4, 400)
	fmt.Printf("[LRUCache] Get(3): %d (削除済み)\n", cache.Get(3)) // -1
	fmt.Printf("[LRUCache] Get(4): %d\n", cache.Get(4))             // 400
	fmt.Printf("[LRUCache] Get(1): %d\n", cache.Get(1))             // 100
}
