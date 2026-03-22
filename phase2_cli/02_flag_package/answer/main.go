package main

import (
	"flag"
	"fmt"
)

func main() {
	// 演習用フラグ定義
	msg := flag.String("msg", "Hello", "message to print")
	count := flag.Int("count", 1, "number of times to print")
	verbose := flag.Bool("v", false, "verbose mode")

	flag.Parse()

	if *verbose {
		fmt.Println("Verbose mode is ON")
	}

	for i := 0; i < *count; i++ {
		fmt.Printf("Message: %s\n", *msg)
	}
}
