package main

import (
	"fmt"
	"sort"
)

// Person はソート対象の構造体
type Person struct {
	Name string
	Age  int
}

func main() {
	// 演習1: バブルソートの実装
	// 隣接する要素を比較・交換して、最大値を末尾に「泡のように」浮かび上がらせる
	fmt.Println("=== [BubbleSort] バブルソート ===")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("[BubbleSort] 入力: %v\n", arr)
	bubbleSort(arr)
	fmt.Printf("[BubbleSort] 結果: %v\n", arr)

	// 演習2: sort.Interface の実装
	// sort.Sort() を使うために、Len, Less, Swap の3メソッドを実装してください
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

// 演習1: bubbleSort を実装してください
func bubbleSort(arr []int) {
	// ここに実装を書いてください
	// ヒント: 外側ループはlen(arr)-1回
	// 内側ループで隣同士を比較して、大きい方を右にswap
}

// 演習2: ByAge に sort.Interface の3メソッドを実装してください
type ByAge []Person

func (a ByAge) Len() int {
	// ここに実装を書いてください
	return 0
}

func (a ByAge) Less(i, j int) bool {
	// ここに実装を書いてください（年齢の昇順）
	return false
}

func (a ByAge) Swap(i, j int) {
	// ここに実装を書いてください
}
