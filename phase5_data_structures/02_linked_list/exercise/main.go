package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 演習1: サイクル検出（Floyd's Tortoise and Hare Algorithm）
	// リストにサイクル（循環）があるかどうかを検出する関数を実装してください
	// ヒント: slow（1歩ずつ）とfast（2歩ずつ）の2ポインタを使う
	fmt.Println("=== [Cycle] サイクル検出 ===")
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // サイクル: 4 → 2

	fmt.Printf("[Cycle] サイクルが検出されました: %v\n", hasCycle(node1))

	// 演習2: 2つのソート済みリストのマージ
	// 2つのソート済みリストを1つのソート済みリストにマージする関数を実装してください
	// ヒント: ダミーヘッドを使うと先頭の処理が楽になる
	fmt.Println("\n=== [Merge] ソート済みリストのマージ ===")
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
	merged := mergeTwoLists(l1, l2)
	fmt.Print("[Merge] 結果: ")
	printList(merged)
}

// 演習1: hasCycle を実装してください
func hasCycle(head *ListNode) bool {
	// ここに実装を書いてください
	// ヒント: slowとfastの2つのポインタを用意
	// slowは1歩ずつ、fastは2歩ずつ進む
	// もしサイクルがあれば、いつか必ずslow == fastになる
	return false
}

// 演習2: mergeTwoLists を実装してください
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	// ここに実装を書いてください
	// ヒント: dummy := &ListNode{} を作り、currentポインタで繋いでいく
	// l1とl2のValを比較して小さい方をcurrent.Nextに繋ぐ
	return nil
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d → ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
