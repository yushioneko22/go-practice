package main

import (
	"encoding/json"
	"net/http"
)

type Event struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	e := Event{
		Title: "Go Web Practice",
		Date:  "2026-03-24",
	}

	json.NewEncoder(w).Encode(e)
}

func main() {
	http.HandleFunc("/event", eventHandler)
	http.ListenAndServe(":8080", nil)
}
