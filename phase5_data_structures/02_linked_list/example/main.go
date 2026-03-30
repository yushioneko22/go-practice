package main

import "fmt"

// ListNode は単方向連結リストのノード
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// ========================================
	// 連結リストの構築
	// ========================================
	fmt.Println("=== [LinkedList] 連結リストの構築 ===")
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Print("[LinkedList] 元のリスト: ")
	printList(head)

	// ========================================
	// リストの反転
	// ========================================
	fmt.Println("\n=== [Reverse] リストの反転 ===")
	reversed := reverseList(head)
	fmt.Print("[Reverse] 反転後: ")
	printList(reversed)
}

// reverseList は連結リストを反転する
// prev, current, next の3ポインタを使う
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next // 次のノードを退避
		current.Next = prev  // リンクの向きを逆転
		prev = current       // prevを進める
		current = next       // currentを進める
	}
	return prev // 新しいheadはprev
}

// printList はリストを見やすく表示する
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d → ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
