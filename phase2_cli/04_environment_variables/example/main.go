package main

import (
	"fmt"
	"os"
)

func main() {
	// 取得
	user := os.Getenv("USER")
	fmt.Println("User:", user)

	// セット
	os.Setenv("MY_APP_NAME", "GoLearning")
	fmt.Println("App Name:", os.Getenv("MY_APP_NAME"))

	// 存在確認
	val, ok := os.LookupEnv("DB_PASS")
	if !ok {
		fmt.Println("DB_PASS is not set")
	} else {
		fmt.Println("DB_PASS:", val)
	}
}
