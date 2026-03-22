package main

import (
	"net/http"
)

// 演習: Message string を持つ Response 構造体を定義してください

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// 演習: GETメソッド以外を弾く
	// 演習: Content-Type を application/json に設定
	// 演習: JSONで { "message": "Hello, Gopher!" } を返却
}

func main() {
	// 演習: /greet にハンドラを登録
	// 演習: 8080ポートで起動
}
