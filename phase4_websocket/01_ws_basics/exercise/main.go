package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. Upgrade を実行してください
	// 2. メッセージをループで受信(Read)してください
	// 3. 受信した内容の頭に "[Echo] " を付けて送信(Write)してください
}

func main() {
	http.HandleFunc("/ws", handler)
	http.ListenAndServe(":8080", nil)
}
