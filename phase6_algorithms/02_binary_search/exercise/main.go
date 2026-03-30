package main

import "fmt"

func main() {
	// 演習1: 回転ソート配列の探索
	// [0,1,2,4,5,6,7] が [4,5,6,7,0,1,2] に回転された配列でtargetを探す
	// ヒント: midで配列を2つに分けると、少なくとも片方はソート済み
	fmt.Println("=== [RotatedSearch] 回転ソート配列の探索 ===")
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Printf("[RotatedSearch] target=0 → index=%d\n", searchRotated(nums, 0))   // 4
	fmt.Printf("[RotatedSearch] target=3 → index=%d\n", searchRotated(nums, 3))   // -1
	fmt.Printf("[RotatedSearch] target=5 → index=%d\n", searchRotated(nums, 5))   // 1

	// 演習2: 整数の平方根
	// x の平方根の整数部分を返す（小数は切り捨て）
	// ヒント: 0〜xの範囲で「mid*mid <= x」を満たす最大のmidを二分探索
	fmt.Println("\n=== [Sqrt] 整数の平方根 ===")
	fmt.Printf("[Sqrt] sqrt(8) = %d\n", mySqrt(8))    // 2
	fmt.Printf("[Sqrt] sqrt(16) = %d\n", mySqrt(16))  // 4
	fmt.Printf("[Sqrt] sqrt(1) = %d\n", mySqrt(1))    // 1
}

// 演習1: searchRotated を実装してください
func searchRotated(nums []int, target int) int {
	// ここに実装を書いてください
	// ヒント: low, highを設定し、midで2分割
	// 左半分がソート済みか右半分がソート済みかを判定
	// targetがソート済み半分の範囲内にあればそちらを探索
	return -1
}

// 演習2: mySqrt を実装してください
func mySqrt(x int) int {
	// ここに実装を書いてください
	// ヒント: low=0, high=x として二分探索
	// mid*mid <= x を満たす最大のmidを探す
	return 0
}
