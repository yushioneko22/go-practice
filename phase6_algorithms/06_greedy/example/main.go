package main

import (
	"fmt"
	"sort"
)

func main() {
	// ========================================
	// コインの貪欲法
	// ========================================
	fmt.Println("=== [Coin] コインの貪欲法 ===")
	coins := []int{500, 100, 50, 10, 5, 1}
	amount := 1234
	result := greedyCoin(coins, amount)
	fmt.Printf("[Coin] %d円のお釣り: %v (計%d枚)\n", amount, result, countCoins(result))

	// ========================================
	// 区間スケジューリング
	// ========================================
	fmt.Println("\n=== [Schedule] 区間スケジューリング ===")
	meetings := [][]int{
		{1, 3}, {2, 4}, {3, 5}, {0, 6}, {5, 7}, {3, 8}, {5, 9}, {6, 10}, {8, 11}, {8, 12},
	}
	selected := intervalSchedule(meetings)
	fmt.Printf("[Schedule] 最大会議数: %d\n", len(selected))
	for _, m := range selected {
		fmt.Printf("[Schedule] 会議: %d:00〜%d:00\n", m[0], m[1])
	}
}

func greedyCoin(coins []int, amount int) map[int]int {
	result := make(map[int]int)
	for _, coin := range coins {
		if amount >= coin {
			result[coin] = amount / coin
			amount %= coin
		}
	}
	return result
}

func countCoins(m map[int]int) int {
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}

// intervalSchedule は終了時刻が早い順にソートして貪欲に選ぶ
func intervalSchedule(meetings [][]int) [][]int {
	// 終了時刻でソート
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][1] < meetings[j][1]
	})

	var selected [][]int
	lastEnd := -1

	for _, m := range meetings {
		if m[0] >= lastEnd { // 前の会議が終わってから始まる
			selected = append(selected, m)
			lastEnd = m[1]
		}
	}
	return selected
}
