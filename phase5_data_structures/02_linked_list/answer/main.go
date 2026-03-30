package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("=== [Cycle] サイクル検出 ===")
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Printf("[Cycle] サイクルが検出されました: %v\n", hasCycle(node1))

	fmt.Println("\n=== [Merge] ソート済みリストのマージ ===")
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
	merged := mergeTwoLists(l1, l2)
	fmt.Print("[Merge] 結果: ")
	printList(merged)
}

// hasCycle はFloyd's Algorithmでサイクルを検出する
// slow: 1歩ずつ進む、fast: 2歩ずつ進む
// サイクルがあれば必ずslowとfastが出会う（陸上トラックで速い人が遅い人に追いつくのと同じ）
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// mergeTwoLists は2つのソート済みリストをマージする
// ダミーヘッドを使うことで、先頭ノードの特別扱いを避ける
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// 残りを繋ぐ
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d → ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
