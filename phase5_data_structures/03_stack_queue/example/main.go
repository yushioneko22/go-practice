package main

import "fmt"

func main() {
	// ========================================
	// スタック（LIFO）
	// ========================================
	fmt.Println("=== [Stack] スタック操作 ===")
	stack := []int{}
	stack = append(stack, 1) // push
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Printf("[Stack] Push後: %v\n", stack)

	top := stack[len(stack)-1]          // peek
	stack = stack[:len(stack)-1]        // pop
	fmt.Printf("[Stack] Pop: %d, 残り: %v\n", top, stack)

	// ========================================
	// キュー（FIFO）
	// ========================================
	fmt.Println("\n=== [Queue] キュー操作 ===")
	queue := []int{}
	queue = append(queue, 10) // enqueue
	queue = append(queue, 20)
	queue = append(queue, 30)
	fmt.Printf("[Queue] Enqueue後: %v\n", queue)

	front := queue[0]    // peek
	queue = queue[1:]     // dequeue
	fmt.Printf("[Queue] Dequeue: %d, 残り: %v\n", front, queue)

	// ========================================
	// 有効な括弧の判定
	// ========================================
	fmt.Println("\n=== [Brackets] 有効な括弧の判定 ===")
	tests := []string{"()", "()[]{}", "(]", "({[]})", "((()"}
	for _, s := range tests {
		fmt.Printf("[Brackets] \"%s\" → %v\n", s, isValid(s))
	}
}

// isValid は括弧の組み合わせが正しいかを判定する
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch) // 開き括弧はpush
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false // スタックが空 or 対応しない
			}
			stack = stack[:len(stack)-1] // 対応する開き括弧をpop
		}
	}
	return len(stack) == 0 // スタックが空なら全て対応済み
}
