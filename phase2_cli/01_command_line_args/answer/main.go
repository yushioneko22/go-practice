package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 演習1: 引数の数とスライスを表示
	args := os.Args[1:]
	fmt.Printf("引数の数: %d\n", len(args))
	fmt.Printf("引数リスト: %v\n", args)

	// 演習2: 名前の表示
	name := "Guest"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s!\n", name)

	// 演習3: 2つの数値の合計
	if len(os.Args) > 2 {
		n1, _ := strconv.Atoi(os.Args[1])
		n2, _ := strconv.Atoi(os.Args[2])
		fmt.Printf("合計: %d\n", n1+n2)
	}

	// 演習4: プログラム名
	fmt.Printf("Program Name: %s\n", os.Args[0])
}
