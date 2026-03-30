package main

import "fmt"

func main() {
	// ========================================
	// Two Sum問題: O(n) 解法
	// ========================================
	fmt.Println("=== [TwoSum] Two Sum問題 ===")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("[TwoSum] nums=%v, target=%d → indices=%v\n", nums, target, result)

	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Printf("[TwoSum] nums=%v, target=%d → indices=%v\n", nums2, target2, result2)

	// ========================================
	// スライスの内部構造
	// ========================================
	fmt.Println("\n=== [Slice] スライスの内部構造 ===")
	s := make([]int, 0, 3)
	fmt.Printf("[Slice] 初期状態: len=%d, cap=%d\n", len(s), cap(s))

	for i := 1; i <= 10; i++ {
		s = append(s, i)
		fmt.Printf("[Slice] append(%d): len=%d, cap=%d\n", i, len(s), cap(s))
	}
}

// twoSum はハッシュマップを使ったO(n)解法
// 各要素を見るたびに「targetとの差分」がすでに見つかっているかチェックする
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // 値 → インデックス
	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}
