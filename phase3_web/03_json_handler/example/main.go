package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	// 1. レスポンスをJSONに設定
	w.Header().Set("Content-Type", "application/json")

	// 2. データの準備
	e := Event{ID: 1, Title: "Learn Go Web"}

	// 3. JSONに変換して書き出し
	fmt.Println("[Server] Sending JSON response...")
	json.NewEncoder(w).Encode(e)
}

func main() {
	http.HandleFunc("/event", eventHandler)
	fmt.Println("[Server] Listening on http://localhost:8080/event")
	http.ListenAndServe(":8080", nil)
}
