package main

import "fmt"

func main() {
	fmt.Println("=== [Brackets] 有効な括弧 ===")
	tests := []string{"({[]})", "({[})", "", "((()", "()[]{}"}
	for _, s := range tests {
		fmt.Printf("[Brackets] \"%s\" → %v\n", s, isValid(s))
	}

	fmt.Println("\n=== [MinStack] 最小値スタック ===")
	ms := NewMinStack()
	ms.Push(5)
	ms.Push(3)
	ms.Push(7)
	fmt.Printf("[MinStack] Top=%d, Min=%d\n", ms.Top(), ms.GetMin())
	ms.Pop()
	fmt.Printf("[MinStack] Pop後: Top=%d, Min=%d\n", ms.Top(), ms.GetMin())
	ms.Pop()
	fmt.Printf("[MinStack] Pop後: Top=%d, Min=%d\n", ms.Top(), ms.GetMin())
}

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// MinStack はすべての操作が O(1) のスタック
// メインスタックと並行して、その時点の最小値を記録するスタックを持つ
type MinStack struct {
	stack    []int
	minStack []int // 各時点での最小値を追跡
}

func NewMinStack() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	// minStackが空、または新しい値が現在の最小値以下なら追加
	if len(ms.minStack) == 0 || val <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, val)
	}
}

func (ms *MinStack) Pop() {
	top := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	// popした値が最小値と同じなら、minStackからも削除
	if top == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}
