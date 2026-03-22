package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// 1. メソッドのチェック
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. クエリパラメータの取得
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/greet", greetHandler)

	fmt.Println("[Server] Listening on http://localhost:8080/greet")
	http.ListenAndServe(":8080", nil)
}
