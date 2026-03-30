package main

import "fmt"

func main() {
	// 演習1: Longest Increasing Subsequence
	// 最長増加部分列の長さを返す（部分列は連続でなくてもよい）
	// dp[i] = nums[i]で終わるLISの長さ
	fmt.Println("=== [LIS] Longest Increasing Subsequence ===")
	fmt.Printf("[LIS] [10,9,2,5,3,7,101,18] → %d\n",
		lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Printf("[LIS] [0,1,0,3,2,3] → %d\n",
		lengthOfLIS([]int{0, 1, 0, 3, 2, 3})) // 4

	// 演習2: Word Break
	// 文字列sが辞書の単語に分割可能かを判定
	// dp[i] = s[:i]が分割可能か
	fmt.Println("\n=== [WordBreak] Word Break ===")
	fmt.Printf("[WordBreak] \"leetcode\", [\"leet\",\"code\"] → %v\n",
		wordBreak("leetcode", []string{"leet", "code"})) // true
	fmt.Printf("[WordBreak] \"applepenapple\", [\"apple\",\"pen\"] → %v\n",
		wordBreak("applepenapple", []string{"apple", "pen"})) // true
	fmt.Printf("[WordBreak] \"catsandog\", [\"cats\",\"dog\",\"sand\",\"and\",\"cat\"] → %v\n",
		wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false

	// 演習3: Unique Paths
	// m x n グリッドの左上から右下へ（右か下のみ）行く経路数
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	fmt.Println("\n=== [UniquePaths] Unique Paths ===")
	fmt.Printf("[UniquePaths] 3x7 → %d\n", uniquePaths(3, 7)) // 28
	fmt.Printf("[UniquePaths] 3x3 → %d\n", uniquePaths(3, 3)) // 6
}

// 演習1: lengthOfLIS を実装してください
func lengthOfLIS(nums []int) int {
	// ヒント: dp[i] = nums[i]で終わるLISの長さ（初期値1）
	// j < i かつ nums[j] < nums[i] なら dp[i] = max(dp[i], dp[j]+1)
	return 0
}

// 演習2: wordBreak を実装してください
func wordBreak(s string, wordDict []string) bool {
	// ヒント: dp[i] = s[:i]が分割可能か
	// dp[0] = true、各位置iについて、すべての単語wを試す
	// dp[i-len(w)] が true かつ s[i-len(w):i] == w なら dp[i] = true
	return false
}

// 演習3: uniquePaths を実装してください
func uniquePaths(m, n int) int {
	// ヒント: dp[i][j] = dp[i-1][j] + dp[i][j-1]、端は1
	return 0
}
