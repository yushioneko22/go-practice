package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// 演習1
	js := `{"id": 101, "name": "Pen"}`
	var item Item
	json.Unmarshal([]byte(js), &item)
	fmt.Println("Item Name:", item.Name)

	// 演習2
	newItem := Item{ID: 200, Name: "Book"}
	data, _ := json.Marshal(newItem)
	fmt.Println("JSON String:", string(data))

	// 演習4
	jsList := `[{"id": 1}, {"id": 2}]`
	var items []Item
	json.Unmarshal([]byte(jsList), &items)
	sum := 0
	for _, v := range items {
		sum += v.ID
	}
	fmt.Println("Sum of IDs:", sum)
}
