package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 演習1: Sliding Window Rate Limiterの実装
	// 過去N秒間のリクエスト数をカウントし、上限を超えたら拒否する
	fmt.Println("=== 演習1: Sliding Window Rate Limiter ===")
	// limiter := NewSlidingWindow(5, 10*time.Second) // 10秒間に5リクエストまで
	// for i := 0; i < 7; i++ {
	//     fmt.Printf("[SlidingWindow] リクエスト%d: %v\n", i+1, limiter.Allow())
	// }

	// 演習2: IPアドレスベースの制限
	// IPアドレスごとに独立したRate Limiterを持つ仕組みを実装
	fmt.Println("\n=== 演習2: IPベースRate Limiter ===")
	// ipLimiter := NewIPRateLimiter(3, 1) // 最大3トークン、1/秒補充
	// fmt.Printf("[IP] 192.168.1.1: %v\n", ipLimiter.Allow("192.168.1.1"))
	// fmt.Printf("[IP] 192.168.1.2: %v\n", ipLimiter.Allow("192.168.1.2"))
}

// 演習1: SlidingWindowLimiter を実装してください
type SlidingWindowLimiter struct {
	mu         sync.Mutex
	timestamps []time.Time // リクエストのタイムスタンプを記録
	maxReqs    int         // ウィンドウ内の最大リクエスト数
	window     time.Duration // ウィンドウサイズ
}

func NewSlidingWindow(maxReqs int, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		maxReqs: maxReqs,
		window:  window,
	}
}

func (s *SlidingWindowLimiter) Allow() bool {
	// ここに実装を書いてください
	// ヒント:
	// 1. ウィンドウ外の古いタイムスタンプを削除
	// 2. 現在のリクエスト数が上限以下なら許可＆タイムスタンプ追加
	return false
}

// 演習2: IPRateLimiter を実装してください
type IPRateLimiter struct {
	mu       sync.Mutex
	limiters map[string]*SlidingWindowLimiter // IP → Limiter
	// ここにフィールドを追加してください
}

// func NewIPRateLimiter(maxReqs int, refillRate float64) *IPRateLimiter { }
// func (ipl *IPRateLimiter) Allow(ip string) bool { }

// ダミーのmain用（コンパイルエラー回避）
var _ = fmt.Println
var _ = sync.Mutex{}
var _ = time.Now
