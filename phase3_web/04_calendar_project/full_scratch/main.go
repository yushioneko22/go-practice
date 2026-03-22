package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

var (
	events = []Event{}
	mu     sync.Mutex
)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)

	case http.MethodPost:
		var e Event
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mu.Lock()
		e.ID = len(events) + 1
		events = append(events, e)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(e)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/events", eventsHandler)
	fmt.Println("[Standard] Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
