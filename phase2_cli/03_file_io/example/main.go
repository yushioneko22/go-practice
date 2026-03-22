package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "hello.txt"
	content := []byte("Hello, File I/O in Go!")

	// 書き込み
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	// 読み込み
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("File content:", string(data))
}
