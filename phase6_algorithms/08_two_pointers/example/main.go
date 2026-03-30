package main

import "fmt"

func main() {
	// ========================================
	// ソート済み配列のTwo Sum
	// ========================================
	fmt.Println("=== [TwoSum] ソート済み配列のTwo Sum ===")
	nums := []int{2, 7, 11, 15}
	target := 9
	left, right := twoSumSorted(nums, target)
	fmt.Printf("[TwoSum] nums=%v, target=%d → [%d, %d]\n", nums, target, left, right)

	nums2 := []int{1, 3, 4, 5, 7, 10, 11}
	target2 := 9
	l2, r2 := twoSumSorted(nums2, target2)
	fmt.Printf("[TwoSum] nums=%v, target=%d → [%d, %d]\n", nums2, target2, l2, r2)

	// ========================================
	// 回文判定
	// ========================================
	fmt.Println("\n=== [Palindrome] 回文判定 ===")
	words := []string{"racecar", "hello", "madam", "level", "world"}
	for _, w := range words {
		fmt.Printf("[Palindrome] \"%s\" → %v\n", w, isPalindrome(w))
	}
}

// twoSumSorted はソート済み配列で Two Pointers を使う
func twoSumSorted(nums []int, target int) (int, int) {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return left, right
		} else if sum < target {
			left++ // 合計が小さすぎる → 左を進める
		} else {
			right-- // 合計が大きすぎる → 右を戻す
		}
	}
	return -1, -1
}

// isPalindrome は Two Pointers で回文を判定する
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
