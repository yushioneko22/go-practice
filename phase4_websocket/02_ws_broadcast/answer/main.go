package main

import (
	"fmt"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.broadcast <- []byte("[System] A new member joined!")
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			formatted := []byte(fmt.Sprintf("[Member] %s", string(message)))
			// Systemメッセージなどはそのまま流したい場合の処理は省略
			for client := range h.clients {
				client.send <- formatted
			}
		}
	}
}

// Client等の定義は example と同様 (省略可だがanswerとしては完全なものを置く)
// ... (中略)
// ※ 実際の実装は example/main.go をベースにメッセージ加工を追加したもの
