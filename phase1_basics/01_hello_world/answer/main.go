package main

import "fmt"

func main() {
	// 演習1: 自分の名前を表示してみましょう
	fmt.Println("My name is Gopher")

	// 演習2: fmt.Printf を使って「私は 20 歳です」と表示してみましょう
	age := 20
	fmt.Printf("私は %d 歳です\n", age)

	// 演習3: 2行に分けてメッセージを表示してみてください
	fmt.Println("Line 1")
	fmt.Println("Line 2")

	// 演習4: Hello と World の間にスペースを入れずに表示してみてください
	fmt.Print("Hello", "World\n")
	// または
	// fmt.Println("Hello" + "World")
}
