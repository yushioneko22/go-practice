package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res := Response{Message: "Hello, Gopher!"}
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	http.ListenAndServe(":8080", nil)
}
