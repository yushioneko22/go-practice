package main

import (
	"flag"
	"fmt"
)

func main() {
	// フラグの定義 (ポインタが返る)
	name := flag.String("name", "World", "name to greet")

	// 変数に直接格納
	var age int
	flag.IntVar(&age, "age", 0, "age of the person")

	// 解析の実行 (必須！)
	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("Age: %d\n", age)

	// フラグ以外の引数
	fmt.Println("Tail args:", flag.Args())
}
