package main

import "fmt"

func main() {
	// ========================================
	// 階段の登り方（Climbing Stairs）
	// ========================================
	fmt.Println("=== [Stairs] 階段の登り方 ===")
	for _, n := range []int{2, 3, 5, 10} {
		fmt.Printf("[Stairs] %d段の登り方: %d通り\n", n, climbStairs(n))
	}

	// ========================================
	// DPテーブルの構築過程を可視化
	// ========================================
	fmt.Println("\n=== [Stairs] 計算過程の可視化 (n=6) ===")
	climbStairsVisualize(6)
}

// climbStairs はボトムアップDPで階段の登り方を求める
// dp[i] = dp[i-1] + dp[i-2]
// 「i段目に到達する方法 = i-1段目から1段登る + i-2段目から2段登る」
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// 空間最適化: 前の2つだけ保持
	prev2, prev1 := 1, 2
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}
	return prev1
}

func climbStairsVisualize(n int) {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	fmt.Printf("[Stairs] dp[1] = %d (ベースケース)\n", dp[1])
	fmt.Printf("[Stairs] dp[2] = %d (ベースケース)\n", dp[2])

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		fmt.Printf("[Stairs] dp[%d] = dp[%d] + dp[%d] = %d + %d = %d\n",
			i, i-1, i-2, dp[i-1], dp[i-2], dp[i])
	}
	fmt.Printf("[Stairs] 答え: %d通り\n", dp[n])
}
