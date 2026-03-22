package main

import "fmt"

func main() {
	// 演習1: string型の変数 `city` を宣言
	var city string = "Tokyo"
	fmt.Println("City:", city)

	// 演習2: plus関数の呼び出し
	result := plus(5, 7)
	fmt.Println("Sum:", result)

	// 演習3: if文
	x := 15
	if x > 10 {
		fmt.Println("Large")
	} else {
		fmt.Println("Small")
	}

	// 演習4: 短縮宣言 :=
	pi := 3.14
	fmt.Printf("Pi is %f\n", pi)
}

func plus(a int, b int) int {
	return a + b
}
