package main

import "fmt"

func main() {
	fmt.Println("=== [Knapsack] 0/1ナップサック問題 ===")
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	capacity := 7
	fmt.Printf("[Knapsack] 最大価値: %d\n", knapsack(weights, values, capacity))

	fmt.Println("\n=== [LCS] 最長共通部分列 ===")
	fmt.Printf("[LCS] \"abcde\" と \"ace\" → 長さ: %d\n", lcs("abcde", "ace"))
	fmt.Printf("[LCS] \"abc\" と \"def\" → 長さ: %d\n", lcs("abc", "def"))
	fmt.Printf("[LCS] \"abcdef\" と \"acbcf\" → 長さ: %d\n", lcs("abcdef", "acbcf"))
}

// knapsack は1D DP配列で0/1ナップサック問題を解く
// 空間最適化: 2D → 1Dに圧縮（逆順に走査することで前の行の値を保持）
func knapsack(weights, values []int, capacity int) int {
	dp := make([]int, capacity+1)

	for i := 0; i < len(weights); i++ {
		// 逆順に走査（同じアイテムを2回使わないため）
		for w := capacity; w >= weights[i]; w-- {
			newVal := dp[w-weights[i]] + values[i]
			if newVal > dp[w] {
				dp[w] = newVal
			}
		}
	}
	return dp[capacity]
}

// lcs は2D DPで最長共通部分列の長さを求める
func lcs(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
