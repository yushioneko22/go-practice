package main

import "fmt"

func main() {
	// 演習1: 区間スケジューリング
	// 会議の [開始, 終了] のリストから、重ならない最大数の会議を選んでください
	// ヒント: 終了時刻が早い順にソート → 前の会議の終了後に始まる会議を選ぶ
	fmt.Println("=== [Schedule] 区間スケジューリング ===")
	meetings := [][]int{
		{1, 4}, {2, 3}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9},
	}
	count := maxMeetings(meetings)
	fmt.Printf("[Schedule] 最大会議数: %d\n", count) // 期待値: 4

	// 演習2: ジャンプゲーム
	// nums[i] は位置iから最大何マス進めるかを表す
	// 最後のインデックスに到達できるかを判定してください
	// ヒント: 「現在到達可能な最遠地点」を貪欲に更新する
	fmt.Println("\n=== [Jump] ジャンプゲーム ===")
	fmt.Printf("[Jump] [2,3,1,1,4] → %v\n", canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Printf("[Jump] [3,2,1,0,4] → %v\n", canJump([]int{3, 2, 1, 0, 4})) // false
}

// 演習1: maxMeetings を実装してください
func maxMeetings(meetings [][]int) int {
	// ここに実装を書いてください
	// ヒント: sort.Slice で終了時刻順にソート
	return 0
}

// 演習2: canJump を実装してください
func canJump(nums []int) bool {
	// ここに実装を書いてください
	// ヒント: maxReach変数を用意し、各位置で max(maxReach, i+nums[i]) を更新
	// もし i > maxReach なら到達不可能
	return false
}
