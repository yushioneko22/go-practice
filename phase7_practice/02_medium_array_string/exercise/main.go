package main

import "fmt"

func main() {
	// 演習1: Longest Substring Without Repeating Characters
	// 重複文字を含まない最長の部分文字列の長さを返す
	fmt.Println("=== [LongestSub] Longest Substring Without Repeating Characters ===")
	fmt.Printf("[LongestSub] \"abcabcbb\" → %d\n", lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Printf("[LongestSub] \"bbbbb\" → %d\n", lengthOfLongestSubstring("bbbbb"))       // 1
	fmt.Printf("[LongestSub] \"pwwkew\" → %d\n", lengthOfLongestSubstring("pwwkew"))     // 3

	// 演習2: 3Sum
	// 合計が0になる3つの数のユニークな組み合わせをすべて返す
	fmt.Println("\n=== [3Sum] 3Sum ===")
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("[3Sum] %v → %v\n", nums, threeSum(nums))

	// 演習3: Group Anagrams
	// アナグラム同士をグループ化する
	fmt.Println("\n=== [GroupAnagram] Group Anagrams ===")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	for _, g := range groups {
		fmt.Printf("[GroupAnagram] %v\n", g)
	}
}

// 演習1: lengthOfLongestSubstring を実装してください
func lengthOfLongestSubstring(s string) int {
	// ヒント: スライディングウィンドウ + map[byte]int
	return 0
}

// 演習2: threeSum を実装してください
func threeSum(nums []int) [][]int {
	// ヒント: ソート → 1つ固定 → Two Pointers → 重複スキップ
	return nil
}

// 演習3: groupAnagrams を実装してください
func groupAnagrams(strs []string) [][]string {
	// ヒント: 各文字列をソートしてキーにする → mapでグループ化
	return nil
}
