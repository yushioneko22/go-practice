package main

import "fmt"

func main() {
	// ========================================
	// Product of Array Except Self
	// 各要素について「自分以外の全要素の積」を返す（除算を使わずに O(n)）
	// ========================================
	fmt.Println("=== [Product] Product of Array Except Self ===")
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Printf("[Product] %v → %v\n", nums, result) // [24,12,8,6]

	nums2 := []int{-1, 1, 0, -3, 3}
	result2 := productExceptSelf(nums2)
	fmt.Printf("[Product] %v → %v\n", nums2, result2) // [0,0,9,0,0]
}

// productExceptSelf は左からの累積積と右からの累積積を使う
// step1: result[i] = nums[0] * nums[1] * ... * nums[i-1]（左からの積）
// step2: result[i] *= nums[i+1] * ... * nums[n-1]（右からの積を掛ける）
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// 左からの累積積
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// 右からの累積積を掛ける
	rightProduct := 1
	for i := n - 2; i >= 0; i-- {
		rightProduct *= nums[i+1]
		result[i] *= rightProduct
	}

	return result
}
