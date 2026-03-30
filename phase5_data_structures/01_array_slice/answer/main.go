package main

import "fmt"

func main() {
	fmt.Println("=== [Rotate] 配列の回転 ===")
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums1, k)
	fmt.Printf("[Rotate] k=%d → %v\n", k, nums1)

	fmt.Println("\n=== [RemoveDup] 重複の除去 ===")
	nums2 := []int{1, 1, 2, 2, 3, 4, 4, 5}
	newLen := removeDuplicates(nums2)
	fmt.Printf("[RemoveDup] 結果=%v, 新しい長さ=%d\n", nums2[:newLen], newLen)
}

// reverse はスライスの要素を反転する（in-place）
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// rotate は配列を右にk個分回転させる
// 手法: 全体反転 → 先頭k個反転 → 残り反転
// 例: [1,2,3,4,5,6,7] k=3
//
//	全体反転: [7,6,5,4,3,2,1]
//	先頭3個: [5,6,7,4,3,2,1]
//	残り:    [5,6,7,1,2,3,4]
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n // kがnより大きい場合に対応
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

// removeDuplicates はソート済み配列からin-placeで重複を除去する
// Two Pointersテクニック: slowが書き込み位置、fastが読み取り位置
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
