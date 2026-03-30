package main

import "fmt"

func main() {
	// 演習1: 配列の回転
	// 配列を右にk個分だけ回転させる関数を実装してください
	// 例: [1,2,3,4,5,6,7], k=3 → [5,6,7,1,2,3,4]
	// ヒント: reverseを3回使うテクニックがあります
	fmt.Println("=== [Rotate] 配列の回転 ===")
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums1, k)
	fmt.Printf("[Rotate] k=%d → %v\n", k, nums1)

	// 演習2: ソート済み配列の重複除去
	// ソート済み配列からin-placeで重複を除去し、新しい長さを返す関数を実装してください
	// 例: [1,1,2,2,3,4,4,5] → [1,2,3,4,5], 長さ5
	fmt.Println("\n=== [RemoveDup] 重複の除去 ===")
	nums2 := []int{1, 1, 2, 2, 3, 4, 4, 5}
	newLen := removeDuplicates(nums2)
	fmt.Printf("[RemoveDup] 結果=%v, 新しい長さ=%d\n", nums2[:newLen], newLen)
}

// 演習1: rotate を実装してください
func rotate(nums []int, k int) {
	// ここに実装を書いてください
	// ヒント1: k が配列長より大きい場合を考慮する（k %= len(nums)）
	// ヒント2: 配列全体をreverse → 先頭k個をreverse → 残りをreverse
}

// 演習2: removeDuplicates を実装してください
func removeDuplicates(nums []int) int {
	// ここに実装を書いてください
	// ヒント: 2つのポインタ（slow, fast）を使う
	// slowは「次に書き込む位置」、fastは「次に読む位置」
	return 0
}
