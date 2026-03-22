package main

import (
	"fmt"
	"os"
)

func main() {
	// 演習1
	fmt.Println("USER:", os.Getenv("USER"))

	// 演習2
	os.Setenv("APP_DEBUG", "true")
	fmt.Println("DEBUG:", os.Getenv("APP_DEBUG"))

	// 演習3
	if val, ok := os.LookupEnv("MISSING"); !ok {
		fmt.Println("Not defined")
	} else {
		fmt.Println("Value:", val)
	}

	// 演習4
	env := os.Getenv("APP_ENV")
	if env == "production" {
		fmt.Println("本番環境")
	} else {
		fmt.Println("開発環境")
	}
}
