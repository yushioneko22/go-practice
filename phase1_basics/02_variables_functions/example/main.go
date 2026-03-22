package main

import "fmt"

func main() {
	// 変数宣言
	var name string = "Alice"
	age := 20 // 型推論

	// 関数の呼び出し
	hello(name, age)

	// 条件分岐
	if age >= 20 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are young.")
	}
}

func hello(name string, age int) {
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
