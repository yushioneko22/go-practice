package main

import "fmt"

func main() {
	// 演習1: 0/1ナップサック問題
	// 各アイテムは1つだけ使える。容量内で価値を最大化せよ
	// ヒント: dp[i][w] = i番目までのアイテムで容量wの最大価値
	// 遷移: dp[i][w] = max(dp[i-1][w], dp[i-1][w-weight[i]] + value[i])
	fmt.Println("=== [Knapsack] 0/1ナップサック問題 ===")
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	capacity := 7
	fmt.Printf("[Knapsack] 最大価値: %d\n", knapsack(weights, values, capacity)) // 9

	// 演習2: 最長共通部分列（LCS）
	// 2つの文字列の共通部分列のうち最長のものの長さを求めよ
	// ヒント: dp[i][j] = s1[:i]とs2[:j]のLCSの長さ
	// s1[i-1]==s2[j-1]なら dp[i][j] = dp[i-1][j-1] + 1
	// そうでなければ dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	fmt.Println("\n=== [LCS] 最長共通部分列 ===")
	fmt.Printf("[LCS] \"abcde\" と \"ace\" → 長さ: %d\n", lcs("abcde", "ace"))     // 3
	fmt.Printf("[LCS] \"abc\" と \"def\" → 長さ: %d\n", lcs("abc", "def"))           // 0
	fmt.Printf("[LCS] \"abcdef\" と \"acbcf\" → 長さ: %d\n", lcs("abcdef", "acbcf")) // 4
}

// 演習1: knapsack を実装してください
func knapsack(weights, values []int, capacity int) int {
	// ここに実装を書いてください
	return 0
}

// 演習2: lcs を実装してください
func lcs(s1, s2 string) int {
	// ここに実装を書いてください
	return 0
}
