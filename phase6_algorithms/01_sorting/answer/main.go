package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== [BubbleSort] バブルソート ===")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("[BubbleSort] 入力: %v\n", arr)
	bubbleSort(arr)
	fmt.Printf("[BubbleSort] 結果: %v\n", arr)

	fmt.Println("\n=== [CustomSort] カスタムソート ===")
	people := ByAge{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Dave", 20},
	}
	sort.Sort(people)
	for _, p := range people {
		fmt.Printf("[CustomSort] %s (%d歳)\n", p.Name, p.Age)
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
