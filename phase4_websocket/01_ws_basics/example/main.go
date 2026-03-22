package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// 1. アップグレーダーの設定
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 2. HTTPをWebSocketへアップグレード
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("[Server] Client connected!")

	for {
		// 3. メッセージの受信
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		fmt.Printf("[Server] Received: %s\n", string(p))

		// 4. メッセージの送信
		response := []byte("Processed: " + string(p))
		if err := conn.WriteMessage(messageType, response); err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	fmt.Println("[Server] WebSocket server starting on :8080/ws")
	http.ListenAndServe(":8080", nil)
}
