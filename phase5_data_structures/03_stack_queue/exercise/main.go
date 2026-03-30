package main

import "fmt"

func main() {
	// 演習1: 有効な括弧の判定
	// スタックを使って括弧の対応をチェックする関数を自分で実装してください
	fmt.Println("=== [Brackets] 有効な括弧 ===")
	tests := []string{"({[]})", "({[})", "", "((()", "()[]{}"}
	for _, s := range tests {
		fmt.Printf("[Brackets] \"%s\" → %v\n", s, isValid(s))
	}

	// 演習2: MinStack
	// Push, Pop, Top, GetMin のすべてが O(1) で動作するスタックを実装してください
	// ヒント: メインのスタックとは別に、最小値を追跡するスタックを持つ
	fmt.Println("\n=== [MinStack] 最小値スタック ===")
	ms := NewMinStack()
	ms.Push(5)
	ms.Push(3)
	ms.Push(7)
	fmt.Printf("[MinStack] Top=%d, Min=%d\n", ms.Top(), ms.GetMin()) // Top=7, Min=3
	ms.Pop()
	fmt.Printf("[MinStack] Pop後: Top=%d, Min=%d\n", ms.Top(), ms.GetMin()) // Top=3, Min=3
	ms.Pop()
	fmt.Printf("[MinStack] Pop後: Top=%d, Min=%d\n", ms.Top(), ms.GetMin()) // Top=5, Min=5
}

// 演習1: isValid を実装してください
func isValid(s string) bool {
	// ここに実装を書いてください
	return false
}

// 演習2: MinStack を実装してください
type MinStack struct {
	// ここにフィールドを定義してください
	// ヒント: stack []int と minStack []int の2つが必要
}

func NewMinStack() *MinStack {
	// ここに実装を書いてください
	return &MinStack{}
}

func (ms *MinStack) Push(val int) {
	// ここに実装を書いてください
}

func (ms *MinStack) Pop() {
	// ここに実装を書いてください
}

func (ms *MinStack) Top() int {
	// ここに実装を書いてください
	return 0
}

func (ms *MinStack) GetMin() int {
	// ここに実装を書いてください
	return 0
}
