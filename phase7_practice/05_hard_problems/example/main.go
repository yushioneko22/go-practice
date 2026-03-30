package main

import (
	"container/heap"
	"fmt"
)

// ListNode は連結リストのノード
type ListNode struct {
	Val  int
	Next *ListNode
}

// NodeHeap はヒープに使うListNodeのスライス
type NodeHeap []*ListNode

func (h NodeHeap) Len() int            { return len(h) }
func (h NodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	// ========================================
	// Merge K Sorted Lists（ヒープ使用）
	// K個のソート済みリストを1つにマージする
	// ========================================
	fmt.Println("=== [MergeK] Merge K Sorted Lists ===")
	lists := []*ListNode{
		{1, &ListNode{4, &ListNode{5, nil}}},
		{1, &ListNode{3, &ListNode{4, nil}}},
		{2, &ListNode{6, nil}},
	}

	merged := mergeKLists(lists)
	fmt.Print("[MergeK] ")
	printList(merged) // 1→1→2→3→4→4→5→6→nil
}

// mergeKLists はMinHeapで各リストの先頭を管理
// 常に最小のノードを取り出し、そのnextをヒープに追加
func mergeKLists(lists []*ListNode) *ListNode {
	h := &NodeHeap{}
	heap.Init(h)

	// 各リストの先頭をヒープに入れる
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	dummy := &ListNode{}
	current := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		current.Next = node
		current = current.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d→", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
