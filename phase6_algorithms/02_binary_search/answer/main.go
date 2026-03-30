package main

import "fmt"

func main() {
	fmt.Println("=== [RotatedSearch] 回転ソート配列の探索 ===")
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Printf("[RotatedSearch] target=0 → index=%d\n", searchRotated(nums, 0))
	fmt.Printf("[RotatedSearch] target=3 → index=%d\n", searchRotated(nums, 3))
	fmt.Printf("[RotatedSearch] target=5 → index=%d\n", searchRotated(nums, 5))

	fmt.Println("\n=== [Sqrt] 整数の平方根 ===")
	fmt.Printf("[Sqrt] sqrt(8) = %d\n", mySqrt(8))
	fmt.Printf("[Sqrt] sqrt(16) = %d\n", mySqrt(16))
	fmt.Printf("[Sqrt] sqrt(1) = %d\n", mySqrt(1))
}

func searchRotated(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		// 左半分がソート済み
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// 右半分がソート済み
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	low, high := 0, x
	ans := 0

	for low <= high {
		mid := low + (high-low)/2
		if mid <= x/mid { // mid*mid <= x（オーバーフロー防止）
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
