package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://www.google.com"

	// タイムアウトを設定したクライアントを作成
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	fmt.Printf("Checking: %s ...\n", url)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}
	defer resp.Body.Close() // 必須！

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("System is HEALTHY")
	} else {
		fmt.Println("System is UNHEALTHY")
	}
}
