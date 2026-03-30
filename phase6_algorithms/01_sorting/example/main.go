package main

import "fmt"

func main() {
	// ========================================
	// マージソート
	// ========================================
	fmt.Println("=== [MergeSort] マージソート ===")
	arr1 := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Printf("[MergeSort] 入力: %v\n", arr1)
	sorted1 := mergeSort(arr1)
	fmt.Printf("[MergeSort] 結果: %v\n", sorted1)

	// ========================================
	// クイックソート
	// ========================================
	fmt.Println("\n=== [QuickSort] クイックソート ===")
	arr2 := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Printf("[QuickSort] 入力: %v\n", arr2)
	quickSort(arr2, 0, len(arr2)-1)
	fmt.Printf("[QuickSort] 結果: %v\n", arr2)
}

// mergeSort は配列を半分に分割 → 再帰的にソート → マージする
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

// merge は2つのソート済み配列を1つにマージする
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// quickSort はピボットを基準に分割して再帰的にソート
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

// partition はピボット（末尾）を基準に要素を並び替える
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
