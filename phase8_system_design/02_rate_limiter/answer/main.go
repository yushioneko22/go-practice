package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	mu         sync.Mutex
	timestamps []time.Time
	maxReqs    int
	window     time.Duration
}

func NewSlidingWindow(maxReqs int, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		maxReqs: maxReqs,
		window:  window,
	}
}

func (s *SlidingWindowLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-s.window)

	// ウィンドウ外の古いタイムスタンプを削除
	valid := 0
	for _, ts := range s.timestamps {
		if ts.After(cutoff) {
			s.timestamps[valid] = ts
			valid++
		}
	}
	s.timestamps = s.timestamps[:valid]

	// 上限チェック
	if len(s.timestamps) < s.maxReqs {
		s.timestamps = append(s.timestamps, now)
		return true
	}
	return false
}

type IPRateLimiter struct {
	mu       sync.Mutex
	limiters map[string]*SlidingWindowLimiter
	maxReqs  int
	window   time.Duration
}

func NewIPRateLimiter(maxReqs int, window time.Duration) *IPRateLimiter {
	return &IPRateLimiter{
		limiters: make(map[string]*SlidingWindowLimiter),
		maxReqs:  maxReqs,
		window:   window,
	}
}

func (ipl *IPRateLimiter) Allow(ip string) bool {
	ipl.mu.Lock()
	limiter, ok := ipl.limiters[ip]
	if !ok {
		limiter = NewSlidingWindow(ipl.maxReqs, ipl.window)
		ipl.limiters[ip] = limiter
	}
	ipl.mu.Unlock()
	return limiter.Allow()
}

func main() {
	fmt.Println("=== 演習1: Sliding Window Rate Limiter ===")
	limiter := NewSlidingWindow(5, 10*time.Second)
	for i := 0; i < 7; i++ {
		fmt.Printf("[SlidingWindow] リクエスト%d: %v\n", i+1, limiter.Allow())
	}

	fmt.Println("\n=== 演習2: IPベースRate Limiter ===")
	ipLimiter := NewIPRateLimiter(3, 10*time.Second)
	for i := 0; i < 5; i++ {
		fmt.Printf("[IP] 192.168.1.1 リクエスト%d: %v\n", i+1, ipLimiter.Allow("192.168.1.1"))
	}
	fmt.Printf("[IP] 192.168.1.2 リクエスト1: %v\n", ipLimiter.Allow("192.168.1.2"))
}
