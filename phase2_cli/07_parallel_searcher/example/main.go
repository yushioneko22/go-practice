package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(name string, wg *sync.WaitGroup) {
	defer wg.Done() // 終わったら報告
	fmt.Printf("Hello %s starting...\n", name)
	time.Sleep(1 * time.Second) // 時間のかかる処理のシミュレーション
	fmt.Printf("Hello %s finished!\n", name)
}

func main() {
	var wg sync.WaitGroup
	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		wg.Add(1)              // 一つ追加
		go sayHello(name, &wg) // 裏側で実行！
	}

	fmt.Println("Waiting for all greetings...")
	wg.Wait() // 全員終わるまでここで待機
	fmt.Println("All done!")
}
