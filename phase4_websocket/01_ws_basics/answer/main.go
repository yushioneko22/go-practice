package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		mt, p, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// 演習回答: [Echo] を付けて返す
		msg := fmt.Sprintf("[Echo] %s", string(p))
		if err := conn.WriteMessage(mt, []byte(msg)); err != nil {
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handler)
	http.ListenAndServe(":8080", nil)
}
