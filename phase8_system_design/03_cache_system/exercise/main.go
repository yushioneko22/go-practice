package main

import (
	"fmt"
	"time"
)

func main() {
	// 演習1: LRU Cacheをスクラッチで実装してください
	// Get(key) → O(1)、Put(key, value) → O(1)
	fmt.Println("=== 演習1: LRU Cache ===")
	// cache := NewLRUCache(2)
	// cache.Put(1, 1)
	// cache.Put(2, 2)
	// fmt.Println(cache.Get(1))  // 1
	// cache.Put(3, 3)            // key=2が削除される
	// fmt.Println(cache.Get(2))  // -1

	// 演習2: TTL付きキャッシュ
	// 各エントリに有効期限を設定し、期限切れなら-1を返す
	fmt.Println("\n=== 演習2: TTL付きキャッシュ ===")
	// ttlCache := NewTTLCache(10, 2*time.Second)
	// ttlCache.Put(1, 100)
	// fmt.Println(ttlCache.Get(1)) // 100
	// time.Sleep(3 * time.Second)
	// fmt.Println(ttlCache.Get(1)) // -1 (期限切れ)

	_ = fmt.Println
	_ = time.Now
}

// 演習1: LRUCacheを双方向連結リスト+ハッシュマップで実装してください
// type Node struct { ... }
// type LRUCache struct { ... }
// func NewLRUCache(capacity int) *LRUCache { ... }
// func (c *LRUCache) Get(key int) int { ... }
// func (c *LRUCache) Put(key, value int) { ... }

// 演習2: TTLCache を実装してください
// type TTLCache struct { ... }
// func NewTTLCache(capacity int, ttl time.Duration) *TTLCache { ... }
// func (c *TTLCache) Get(key int) int { ... }
// func (c *TTLCache) Put(key, value int) { ... }
