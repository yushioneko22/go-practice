package main

import "fmt"

func main() {
	// ========================================
	// 基本的な二分探索
	// ========================================
	fmt.Println("=== [BinarySearch] 基本の二分探索 ===")
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	targets := []int{7, 4, 15, 0}
	for _, target := range targets {
		idx := binarySearch(arr, target)
		fmt.Printf("[BinarySearch] arr=%v, target=%d → index=%d\n", arr, target, idx)
	}

	// ========================================
	// Lower Bound（最初の出現位置）
	// ========================================
	fmt.Println("\n=== [LowerBound] 最初の出現位置 ===")
	arr2 := []int{1, 2, 2, 2, 3, 4, 5}
	idx := lowerBound(arr2, 2)
	fmt.Printf("[LowerBound] arr=%v, target=2 → first index=%d\n", arr2, idx)
}

// binarySearch はソート済み配列からtargetのインデックスを返す（見つからなければ-1）
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2 // オーバーフロー防止
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// lowerBound はtarget以上の最初のインデックスを返す
func lowerBound(arr []int, target int) int {
	low, high := 0, len(arr)

	for low < high {
		mid := low + (high-low)/2
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
