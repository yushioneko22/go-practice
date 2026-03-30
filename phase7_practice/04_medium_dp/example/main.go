package main

import "fmt"

func main() {
	// ========================================
	// Coin Change
	// coins = [1,5,10,25] で amount を作る最小枚数
	// ========================================
	fmt.Println("=== [CoinChange] Coin Change ===")
	coins := []int{1, 5, 10, 25}
	amounts := []int{11, 30, 0, 3}
	for _, amount := range amounts {
		fmt.Printf("[CoinChange] coins=%v, amount=%d → %d枚\n",
			coins, amount, coinChange(coins, amount))
	}
}

// coinChange はDP（ボトムアップ）で最小枚数を求める
// dp[i] = 金額iを作るための最小コイン枚数
// 遷移: dp[i] = min(dp[i-coin] + 1) for each coin
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // 「不可能」を表す大きな値
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
