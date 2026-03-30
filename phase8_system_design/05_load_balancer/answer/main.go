package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// === 演習1: Weighted Round Robin ===

type ServerConfig struct {
	URL    string
	Weight int
}

type WeightedRRLB struct {
	mu      sync.Mutex
	servers []ServerConfig
	weights []int // 現在の重みカウンタ
	index   int
}

func NewWeightedRRLB(configs []ServerConfig) *WeightedRRLB {
	weights := make([]int, len(configs))
	for i, c := range configs {
		weights[i] = c.Weight
	}
	return &WeightedRRLB{servers: configs, weights: weights}
}

// Next はSmooth Weighted Round Robinで次のサーバーを選択
func (lb *WeightedRRLB) Next() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	totalWeight := 0
	for _, s := range lb.servers {
		totalWeight += s.Weight
	}

	// 各サーバーの現在の重みに元の重みを加算
	maxWeight := -1
	bestIdx := 0
	for i := range lb.servers {
		lb.weights[i] += lb.servers[i].Weight
		if lb.weights[i] > maxWeight {
			maxWeight = lb.weights[i]
			bestIdx = i
		}
	}

	// 選択されたサーバーの重みから合計を引く
	lb.weights[bestIdx] -= totalWeight
	return lb.servers[bestIdx].URL
}

// === 演習2: ヘルスチェック付きLB ===

type HealthServer struct {
	URL   string
	Alive bool
}

type HealthCheckLB struct {
	mu       sync.RWMutex
	servers  []*HealthServer
	current  int
	interval time.Duration
}

func NewHealthCheckLB(urls []string, interval time.Duration) *HealthCheckLB {
	servers := make([]*HealthServer, len(urls))
	for i, u := range urls {
		servers[i] = &HealthServer{URL: u, Alive: true}
	}
	return &HealthCheckLB{servers: servers, interval: interval}
}

func (lb *HealthCheckLB) StartHealthCheck() {
	go func() {
		ticker := time.NewTicker(lb.interval)
		defer ticker.Stop()
		for range ticker.C {
			for _, s := range lb.servers {
				go func(server *HealthServer) {
					client := http.Client{Timeout: 2 * time.Second}
					resp, err := client.Get(server.URL + "/health")
					lb.mu.Lock()
					if err != nil || resp.StatusCode != 200 {
						server.Alive = false
						fmt.Printf("[Health] %s DOWN\n", server.URL)
					} else {
						server.Alive = true
						fmt.Printf("[Health] %s UP\n", server.URL)
					}
					lb.mu.Unlock()
					if resp != nil {
						resp.Body.Close()
					}
				}(s)
			}
		}
	}()
}

func main() {
	fmt.Println("=== 演習1: Weighted Round Robin ===")
	wlb := NewWeightedRRLB([]ServerConfig{
		{URL: "http://server1:8081", Weight: 3},
		{URL: "http://server2:8082", Weight: 1},
		{URL: "http://server3:8083", Weight: 2},
	})

	count := map[string]int{}
	for i := 0; i < 12; i++ {
		server := wlb.Next()
		count[server]++
		fmt.Printf("[WRR] リクエスト%2d → %s\n", i+1, server)
	}
	fmt.Println("\n[WRR] 振り分け結果:")
	for url, c := range count {
		fmt.Printf("[WRR] %s: %d回\n", url, c)
	}

	fmt.Println("\n=== 演習2: ヘルスチェック ===")
	fmt.Println("[Health] 実サーバーがないため、シミュレーションのみ表示")
	hlb := NewHealthCheckLB([]string{
		"http://localhost:9001",
		"http://localhost:9002",
	}, 5*time.Second)
	_ = hlb
	fmt.Println("[Health] StartHealthCheck() で定期チェック開始（実サーバー起動時に有効）")
}
