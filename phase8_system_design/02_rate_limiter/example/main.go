package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket はトークンバケットアルゴリズムの実装
type TokenBucket struct {
	mu         sync.Mutex
	tokens     float64       // 現在のトークン数
	maxTokens  float64       // 最大トークン数
	refillRate float64       // 1秒あたりの補充速度
	lastRefill time.Time     // 最後にトークンを補充した時刻
}

func NewTokenBucket(maxTokens, refillRate float64) *TokenBucket {
	return &TokenBucket{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow はリクエストを許可するかどうかを判定する
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// トークンを補充
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.maxTokens {
		tb.tokens = tb.maxTokens
	}
	tb.lastRefill = now

	// トークンを消費
	if tb.tokens >= 1 {
		tb.tokens--
		return true
	}
	return false
}

func main() {
	fmt.Println("=== [RateLimiter] Token Bucket ===")

	// 最大5トークン、1秒あたり2トークン補充
	limiter := NewTokenBucket(5, 2)

	// 連続10リクエスト
	fmt.Println("[RateLimiter] 連続10リクエスト:")
	for i := 1; i <= 10; i++ {
		allowed := limiter.Allow()
		status := "ALLOWED"
		if !allowed {
			status = "REJECTED"
		}
		fmt.Printf("[RateLimiter] リクエスト%d: %s\n", i, status)
	}

	// 2秒待ってトークン補充
	fmt.Println("\n[RateLimiter] 2秒待機（トークン補充）...")
	time.Sleep(2 * time.Second)

	fmt.Println("[RateLimiter] 待機後の5リクエスト:")
	for i := 1; i <= 5; i++ {
		allowed := limiter.Allow()
		status := "ALLOWED"
		if !allowed {
			status = "REJECTED"
		}
		fmt.Printf("[RateLimiter] リクエスト%d: %s\n", i, status)
	}
}
