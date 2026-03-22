package main

import (
	"fmt"
	"net/http"
)

// 1. ハンドラ関数の定義
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "[Server] Request received!")
	fmt.Fprint(w, "\nHello, Go Web World!")
}

func main() {
	// 2. ルーティングの設定 (ルートパス "/" に対する処理)
	http.HandleFunc("/", helloHandler)

	fmt.Println("[Server] Starting server at http://localhost:8080")
	
	// 3. サーバーの起動 (ポート 8080)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("[Error] Server failed to start: %v\n", err)
	}
}
