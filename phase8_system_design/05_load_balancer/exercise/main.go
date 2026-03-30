package main

import "fmt"

func main() {
	// 演習1: Weighted Round Robinの実装
	// 各サーバーにweight（重み）を設定し、重みに比例してリクエストを振り分ける
	// ヒント: weight=3のサーバーはweight=1のサーバーの3倍リクエストを受ける
	fmt.Println("=== 演習1: Weighted Round Robin ===")
	// wlb := NewWeightedRRLB([]ServerConfig{
	//     {URL: "http://server1:8081", Weight: 3},
	//     {URL: "http://server2:8082", Weight: 1},
	//     {URL: "http://server3:8083", Weight: 2},
	// })
	// for i := 0; i < 12; i++ {
	//     fmt.Printf("[WRR] リクエスト%d → %s\n", i+1, wlb.Next())
	// }

	// 演習2: ヘルスチェック機能の追加
	// 定期的にHTTP GETでサーバーの生死を確認し、ダウンしたサーバーをプールから除外
	// ヒント: goroutineで定期的にhttp.Getを送り、エラーならAlive=false
	fmt.Println("\n=== 演習2: ヘルスチェック ===")
	// lb := NewHealthCheckLB(servers, 5*time.Second)
	// lb.StartHealthCheck()
}

// 演習1: WeightedRoundRobinLB を実装してください
// type ServerConfig struct { URL string; Weight int }
// type WeightedRRLB struct { ... }
// func NewWeightedRRLB(configs []ServerConfig) *WeightedRRLB { ... }
// func (lb *WeightedRRLB) Next() string { ... }

// 演習2: HealthCheckLB を実装してください
// ヒント: goroutineでticker + http.Get でサーバーの生死を定期チェック
