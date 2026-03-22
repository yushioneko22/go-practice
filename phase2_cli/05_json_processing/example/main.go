package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"user_id"`
	Name string `json:"full_name"`
	Job  string `json:"job,omitempty"`
}

func main() {
	// JSON -> Struct (Unmarshal)
	js := `{"user_id": 1, "full_name": "Gopher"}`
	var u User
	json.Unmarshal([]byte(js), &u)
	fmt.Printf("Struct: %+v\n", u)

	// Struct -> JSON (Marshal)
	u2 := User{ID: 2, Name: "Alice"}
	data, _ := json.Marshal(u2)
	fmt.Println("JSON:", string(data))
}
