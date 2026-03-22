package main

import (
	"net/http"
)

// 演習: Title と Date を持つ Event 構造体を定義してください
// JSONタグを付けて、小文字で出力されるようにしてください

func main() {
	// 演習: /event にアクセスしたら、Event構造体をJSONで返すハンドラを設定してください
	// ステップ:
	// 1. Header に Content-Type: application/json をセット
	// 2. 構造体を作成
	// 3. json.NewEncoder(w).Encode(e) で送信
}
