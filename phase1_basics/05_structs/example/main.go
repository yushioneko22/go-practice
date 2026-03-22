package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	// 初期化
	u1 := User{ID: 1, Name: "Alice"}

	// フィールド更新
	u1.Name = "Bob"

	fmt.Println("User ID:", u1.ID)
	fmt.Println("User Name:", u1.Name)

	// ポインタ経由（Goでは頻繁に使います）
	u2 := &User{ID: 2, Name: "Charlie"}
	fmt.Println("Pointer User Name:", u2.Name)
}
