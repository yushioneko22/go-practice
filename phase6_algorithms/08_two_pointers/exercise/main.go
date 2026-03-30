package main

import "fmt"

func main() {
	// 演習1: 3Sum
	// 合計が0になる3つの数の組み合わせをすべて見つけてください（重複なし）
	// ヒント: ソート → 1つ固定 → 残り2つをTwo Pointersで探す
	fmt.Println("=== [3Sum] 3Sum ===")
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Printf("[3Sum] %v → %v\n", nums, result)

	// 演習2: Container With Most Water
	// 高さの配列が与えられ、2本の線と底辺で最大の水量を求めてください
	// 面積 = min(height[left], height[right]) * (right - left)
	// ヒント: 両端からTwo Pointers、低い方を移動させる
	fmt.Println("\n=== [Water] Container With Most Water ===")
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("[Water] %v → %d\n", height, maxArea(height)) // 49
}

// 演習1: threeSum を実装してください
func threeSum(nums []int) [][]int {
	// ここに実装を書いてください
	// ヒント:
	// 1. ソートする
	// 2. 各要素iを固定し、i+1とlen-1でTwo Pointers
	// 3. 重複をスキップする（同じ値のiは飛ばす、left/rightも同様）
	return nil
}

// 演習2: maxArea を実装してください
func maxArea(height []int) int {
	// ここに実装を書いてください
	// ヒント: left=0, right=len-1 から始めて
	// 面積 = min(height[left], height[right]) * (right-left)
	// 低い方のポインタを移動させる
	return 0
}
