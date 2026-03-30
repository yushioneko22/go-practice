package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 演習1: Merge Two Sorted Lists
	fmt.Println("=== [MergeLists] Merge Two Sorted Lists ===")
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	merged := mergeTwoLists(l1, l2)
	fmt.Print("[MergeLists] ")
	printList(merged) // 1→1→2→3→4→4→nil

	// 演習2: Best Time to Buy and Sell Stock
	// 株価の配列から、1回の売買で得られる最大利益を求めてください
	// ヒント: 「これまでの最安値」を追跡し、各日の利益を計算
	fmt.Println("\n=== [Stock] Best Time to Buy and Sell Stock ===")
	fmt.Printf("[Stock] [7,1,5,3,6,4] → %d\n", maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Printf("[Stock] [7,6,4,3,1] → %d\n", maxProfit([]int{7, 6, 4, 3, 1}))       // 0

	// 演習3: Valid Anagram
	// 2つの文字列がアナグラム（同じ文字の並び替え）かどうかを判定
	// ヒント: 各文字の出現回数をカウントして比較
	fmt.Println("\n=== [Anagram] Valid Anagram ===")
	fmt.Printf("[Anagram] \"anagram\",\"nagaram\" → %v\n", isAnagram("anagram", "nagaram")) // true
	fmt.Printf("[Anagram] \"rat\",\"car\" → %v\n", isAnagram("rat", "car"))                 // false
}

// 演習1: mergeTwoLists を実装してください
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	// ここに実装を書いてください
	return nil
}

// 演習2: maxProfit を実装してください
func maxProfit(prices []int) int {
	// ここに実装を書いてください
	return 0
}

// 演習3: isAnagram を実装してください
func isAnagram(s, t string) bool {
	// ここに実装を書いてください
	return false
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d→", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
