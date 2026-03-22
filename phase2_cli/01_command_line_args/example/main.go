package main

import (
	"fmt"
	"os"
)

func main() {
	// 全体の表示
	fmt.Println("os.Args:", os.Args)

	// プログラム名
	fmt.Println("Program Name:", os.Args[0])

	// 引数がある場合のみ表示
	if len(os.Args) > 1 {
		fmt.Println("First Argument:", os.Args[1])

		// スライスでまとめて取得
		args := os.Args[1:]
		fmt.Println("Clean Args:", args)
	} else {
		fmt.Println("No arguments provided.")
	}
}
