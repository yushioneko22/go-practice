package main

import "fmt"

func main() {
	// スライス
	fruits := []string{"apple", "banana"}
	fruits = append(fruits, "cherry") // 要素の追加
	fmt.Println("Fruits:", fruits)

	// マップ
	scores := make(map[string]int)
	scores["Alice"] = 100
	scores["Bob"] = 90
	fmt.Println("Alice's score:", scores["Alice"])

	// 存在チェック
	val, ok := scores["Charlie"]
	fmt.Println("Charlie's score:", val, "Exists?", ok)
}
