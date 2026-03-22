package main

import "fmt"

func main() {
	// 演習1
	defer fmt.Println("さようなら")
	fmt.Println("こんにちは")

	// 演習2
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

	// 演習3
	// panic("Stop!")
}
