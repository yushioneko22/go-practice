package main

import "fmt"

func main() {
	// 演習1: スライス作成
	nums := []int{1, 2, 3}
	fmt.Println(nums)

	// 演習2: 要素追加
	nums = append(nums, 4)
	fmt.Println(nums)

	// 演習3: マップ作成と追加
	scores := make(map[string]int)
	scores["English"] = 80
	fmt.Println(scores)

	// 演習4: マップ検索
	val, ok := scores["Math"]
	fmt.Printf("Math score: %d Exists: %t\n", val, ok)
}
