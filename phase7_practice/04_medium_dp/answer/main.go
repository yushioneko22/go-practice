package main

import "fmt"

func main() {
	fmt.Println("=== [LIS] Longest Increasing Subsequence ===")
	fmt.Printf("[LIS] [10,9,2,5,3,7,101,18] → %d\n",
		lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Printf("[LIS] [0,1,0,3,2,3] → %d\n",
		lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))

	fmt.Println("\n=== [WordBreak] Word Break ===")
	fmt.Printf("[WordBreak] \"leetcode\", [\"leet\",\"code\"] → %v\n",
		wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Printf("[WordBreak] \"applepenapple\", [\"apple\",\"pen\"] → %v\n",
		wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Printf("[WordBreak] \"catsandog\", [\"cats\",\"dog\",\"sand\",\"and\",\"cat\"] → %v\n",
		wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

	fmt.Println("\n=== [UniquePaths] Unique Paths ===")
	fmt.Printf("[UniquePaths] 3x7 → %d\n", uniquePaths(3, 7))
	fmt.Printf("[UniquePaths] 3x3 → %d\n", uniquePaths(3, 3))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			wLen := len(w)
			if i >= wLen && dp[i-wLen] && s[i-wLen:i] == w {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func uniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
