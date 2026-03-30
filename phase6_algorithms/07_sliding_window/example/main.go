package main

import "fmt"

func main() {
	// ========================================
	// 固定サイズのスライディングウィンドウ
	// ========================================
	fmt.Println("=== [FixedWindow] サイズKの最大部分配列和 ===")
	nums := []int{1, 4, 2, 10, 2, 3, 1, 0, 20}
	k := 4
	maxSum := maxSumSubarray(nums, k)
	fmt.Printf("[FixedWindow] nums=%v, k=%d → 最大和=%d\n", nums, k, maxSum)

	// ========================================
	// 可変サイズのスライディングウィンドウ
	// ========================================
	fmt.Println("\n=== [VarWindow] 合計がtarget以上の最短部分配列 ===")
	nums2 := []int{2, 3, 1, 2, 4, 3}
	target := 7
	minLen := minSubarrayLen(target, nums2)
	fmt.Printf("[VarWindow] target=%d → 最短長=%d\n", target, minLen)
}

// maxSumSubarray は固定サイズKのウィンドウで最大和を求める
func maxSumSubarray(nums []int, k int) int {
	// 最初のウィンドウの和を計算
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum

	// ウィンドウをスライド
	for i := k; i < len(nums); i++ {
		windowSum += nums[i]     // 右端を追加
		windowSum -= nums[i-k]   // 左端を削除
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

// minSubarrayLen は合計がtarget以上の最短部分配列の長さを返す
func minSubarrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := len(nums) + 1

	for right := 0; right < len(nums); right++ {
		sum += nums[right] // 右端を拡張

		for sum >= target { // 条件を満たす間、左端を縮小
			length := right - left + 1
			if length < minLen {
				minLen = length
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
