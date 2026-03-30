package main

import (
	"container/heap"
	"fmt"
)

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println("=== [KthLargest] K番目に大きい要素 ===")
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Printf("[KthLargest] K=%d番目に大きい要素: %d\n", k, findKthLargest(nums, k))

	fmt.Println("\n=== [Median] データストリームの中央値 ===")
	mf := NewMedianFinder()
	mf.AddNum(1)
	fmt.Printf("[Median] Add(1): 中央値=%.1f\n", mf.FindMedian())
	mf.AddNum(2)
	fmt.Printf("[Median] Add(2): 中央値=%.1f\n", mf.FindMedian())
	mf.AddNum(3)
	fmt.Printf("[Median] Add(3): 中央値=%.1f\n", mf.FindMedian())
}

func findKthLargest(nums []int, k int) int {
	h := &IntMinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

// MedianFinder は2つのヒープで中央値を管理する
// maxHeap: 下半分（大きい値が先頭）
// minHeap: 上半分（小さい値が先頭）
// maxHeapの先頭 <= minHeapの先頭 を常に維持
type MedianFinder struct {
	maxHeap *IntMaxHeap // 下半分
	minHeap *IntMinHeap // 上半分
}

func NewMedianFinder() *MedianFinder {
	maxH := &IntMaxHeap{}
	minH := &IntMinHeap{}
	heap.Init(maxH)
	heap.Init(minH)
	return &MedianFinder{maxHeap: maxH, minHeap: minH}
}

func (mf *MedianFinder) AddNum(num int) {
	// まずmaxHeap（下半分）に追加
	heap.Push(mf.maxHeap, num)
	// maxHeapの最大値をminHeap（上半分）に移動して順序を保証
	heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	// サイズのバランスを保つ（maxHeap >= minHeap）
	if mf.maxHeap.Len() < mf.minHeap.Len() {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	}
	return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2.0
}
