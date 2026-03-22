package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// [Step 1] メソッドのチェック
	if r.Method != http.MethodGet {
		http.Error(w, "GET only", http.StatusMethodNotAllowed)
		return
	}

	// [Step 2] レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json")

	// [Step 3] データの生成
	u := User{ID: 100, Name: "Gopher"}

	// [Step 4] JSONエンコードして送信
	json.NewEncoder(w).Encode(u)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("[Server] Starting at http://localhost:8080/user")
	http.ListenAndServe(":8080", nil)
}
