package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "test.txt"

	// 演習1: 作成
	os.WriteFile(filename, []byte("Hello Go File\nSecond Line\n"), 0644)

	// 演習2: 読み込み
	data, _ := os.ReadFile(filename)
	fmt.Println("Content:", string(data))

	// 演習3: 存在チェック
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Not Found")
	} else {
		fmt.Println("File is here!")
	}

	// 演習4: 一行ずつ
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
}
