package main

import (
	"container/heap"
	"fmt"
)

// IntMinHeap はint型のMinHeap
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	// ========================================
	// MinHeapの基本操作
	// ========================================
	fmt.Println("=== [Heap] MinHeapの基本操作 ===")
	h := &IntMinHeap{}
	heap.Init(h)

	values := []int{5, 3, 8, 1, 9, 2}
	for _, v := range values {
		heap.Push(h, v)
		fmt.Printf("[Heap] Push(%d): heap=%v\n", v, *h)
	}

	fmt.Println("\n=== [Heap] 最小値から順に取り出し ===")
	for h.Len() > 0 {
		min := heap.Pop(h).(int)
		fmt.Printf("[Heap] Pop: %d, 残り=%v\n", min, *h)
	}

	// ========================================
	// K番目に大きい要素（デモ）
	// ========================================
	fmt.Println("\n=== [KthLargest] K番目に大きい要素 ===")
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := findKthLargest(nums, k)
	fmt.Printf("[KthLargest] nums=%v, K=%d → %d番目に大きい要素: %d\n", nums, k, k, result)
}

// findKthLargest はサイズKのMinHeapを使ってK番目に大きい要素を求める
// MinHeapにK個だけ入れておくと、先頭が「K番目に大きい値」になる
func findKthLargest(nums []int, k int) int {
	h := &IntMinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h) // K個を超えたら最小値を捨てる
		}
	}
	return (*h)[0] // MinHeapの先頭がK番目に大きい値
}
