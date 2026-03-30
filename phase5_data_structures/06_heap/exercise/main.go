package main

import (
	"container/heap"
	"fmt"
)

// IntMinHeap はint型のMinHeap（これは提供済み）
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
	// heap パッケージの使用を明示（未使用エラー回避）
	_ = heap.Init

	// 演習1: K番目に大きい要素
	// サイズKのMinHeapを使って、K番目に大きい要素を求める関数を実装してください
	// ヒント: ヒープにK個だけ保持する。K個を超えたらPopする
	fmt.Println("=== [KthLargest] K番目に大きい要素 ===")
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Printf("[KthLargest] K=%d番目に大きい要素: %d\n", k, findKthLargest(nums, k))

	// 演習2: データストリームの中央値
	// MaxHeap（下半分）とMinHeap（上半分）の2つを使って中央値を O(1) で返す
	// ヒント: MaxHeapはIntMinHeapのLessを反転させて作る
	fmt.Println("\n=== [Median] データストリームの中央値 ===")
	mf := NewMedianFinder()
	mf.AddNum(1)
	fmt.Printf("[Median] Add(1): 中央値=%.1f\n", mf.FindMedian())
	mf.AddNum(2)
	fmt.Printf("[Median] Add(2): 中央値=%.1f\n", mf.FindMedian())
	mf.AddNum(3)
	fmt.Printf("[Median] Add(3): 中央値=%.1f\n", mf.FindMedian())
}

// 演習1: findKthLargest を実装してください
func findKthLargest(nums []int, k int) int {
	// ここに実装を書いてください
	return 0
}

// 演習2用: MaxHeap（Lessを反転）
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 反転!
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 演習2: MedianFinder を実装してください
type MedianFinder struct {
	// ここにフィールドを定義してください
	// ヒント: maxHeap（下半分）と minHeap（上半分）の2つを使う
	// maxHeapの先頭 <= minHeapの先頭 を常に維持する
}

func NewMedianFinder() *MedianFinder {
	// ここに実装を書いてください
	return &MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	// ここに実装を書いてください
	// ヒント: まずmaxHeapに追加 → maxHeapの先頭をminHeapに移動
	// サイズのバランスが崩れたらminHeapからmaxHeapに戻す
}

func (mf *MedianFinder) FindMedian() float64 {
	// ここに実装を書いてください
	return 0
}
