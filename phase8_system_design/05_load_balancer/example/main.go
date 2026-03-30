package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

// Server はバックエンドサーバーの情報
type Server struct {
	URL     *url.URL
	Alive   bool
	Weight  int
	ConnNum int32
}

// RoundRobinLB はラウンドロビンロードバランサー
type RoundRobinLB struct {
	mu      sync.RWMutex
	servers []*Server
	current uint64
}

func NewRoundRobinLB(urls []string) *RoundRobinLB {
	var servers []*Server
	for _, rawURL := range urls {
		u, _ := url.Parse(rawURL)
		servers = append(servers, &Server{URL: u, Alive: true})
	}
	return &RoundRobinLB{servers: servers}
}

// NextServer は次のサーバーをラウンドロビンで選択する
func (lb *RoundRobinLB) NextServer() *Server {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	// 生存サーバーのみ対象
	alive := make([]*Server, 0)
	for _, s := range lb.servers {
		if s.Alive {
			alive = append(alive, s)
		}
	}
	if len(alive) == 0 {
		return nil
	}

	idx := atomic.AddUint64(&lb.current, 1) % uint64(len(alive))
	return alive[idx]
}

// SetAlive はサーバーの生死状態を設定する
func (lb *RoundRobinLB) SetAlive(serverURL string, alive bool) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	for _, s := range lb.servers {
		if s.URL.String() == serverURL {
			s.Alive = alive
			break
		}
	}
}

func main() {
	fmt.Println("=== [LB] ロードバランサー (Round Robin) ===")

	// デモ: サーバー選択のシミュレーション
	lb := NewRoundRobinLB([]string{
		"http://server1:8081",
		"http://server2:8082",
		"http://server3:8083",
	})

	fmt.Println("[LB] 10リクエストの振り分け:")
	for i := 1; i <= 10; i++ {
		server := lb.NextServer()
		fmt.Printf("[LB] リクエスト%d → %s\n", i, server.URL)
	}

	// サーバー2をダウンさせる
	fmt.Println("\n[LB] server2 をダウンに設定")
	lb.SetAlive("http://server2:8082", false)

	fmt.Println("[LB] ダウン後の5リクエスト:")
	for i := 1; i <= 5; i++ {
		server := lb.NextServer()
		fmt.Printf("[LB] リクエスト%d → %s\n", i, server.URL)
	}

	// リバースプロキシのセットアップ例（実際にはサーバーが必要）
	fmt.Println("\n[LB] リバースプロキシのセットアップ例:")
	fmt.Println("[LB] http.HandleFunc で NextServer().URL にプロキシ転送")

	_ = httputil.NewSingleHostReverseProxy
	_ = http.ListenAndServe
}
