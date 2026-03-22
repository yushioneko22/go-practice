package main

import "fmt"

func main() {
	x := 10
	p := &x // xのアドレスを代入

	fmt.Println("Address:", p)
	fmt.Println("Value via pointer:", *p)

	*p = 20 // ポインタ経由で値を更新
	fmt.Println("Updated x:", x)

	// 値渡しとポインタ渡しの違い
	y := 100
	updateValue(y)
	fmt.Println("After updateValue:", y) // 変化なし

	updatePointer(&y)
	fmt.Println("After updatePointer:", y) // 変化あり
}

func updateValue(n int) {
	n = 0
}

func updatePointer(n *int) {
	*n = 0
}
