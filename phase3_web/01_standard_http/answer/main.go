package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current time is: %s", now)
}

func main() {
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
