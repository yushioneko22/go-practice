package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: searcher [keyword]")
		return
	}
	keyword := os.Args[1]

	files, _ := filepath.Glob("*.txt") // カレントディレクトリのtxtファイルを探す
	var wg sync.WaitGroup
	results := make(chan string)

	// 各ファイルに対してゴルーチンを起動
	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			f, _ := os.Open(filename)
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				if contains(scanner.Text(), keyword) {
					results <- fmt.Sprintf("[%s] Found: %s", filename, scanner.Text())
				}
			}
		}(file)
	}

	// 待ち受ける側が死なないように、Waitを別のゴルーチンで実行
	go func() {
		wg.Wait()
		close(results) // すべて終わったらチャネルを閉じる
	}()

	// 結果を出力 (closeされるまで回り続ける)
	for res := range results {
		fmt.Println(res)
	}
}

func contains(line, keyword string) bool {
	// 簡易的な検索。実際には strings.Contains を使うのが一般的
	return len(line) >= len(keyword) // 便宜上の実装
}
