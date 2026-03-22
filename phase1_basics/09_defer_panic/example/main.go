package main

import "fmt"

func main() {
	defer fmt.Println("End of main") // 最後に実行される

	fmt.Println("Start of main")

	// 複数の defer はスタック（逆順）で実行される
	defer fmt.Println("First defer (but last to run)")
	defer fmt.Println("Second defer (ran before first)")

	// panic の例 (コメントアウトして動作を確認してください)
	// tryPanic()
}

func tryPanic() {
	panic("Critical Error!") // プログラムが止まる
}
