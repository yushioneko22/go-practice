package main

import "fmt"

func main() {
	// 基本的なfor
	for i := 0; i < 3; i++ {
		fmt.Println("count:", i)
	}

	// rangeを使ったスライスの反復処理
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// マップの反復処理
	caps := map[string]string{"Japan": "Tokyo", "USA": "Washington"}
	for key, val := range caps {
		fmt.Printf("Country: %s, Capital: %s\n", key, val)
	}
}
