package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("=== [MergeLists] Merge Two Sorted Lists ===")
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	merged := mergeTwoLists(l1, l2)
	fmt.Print("[MergeLists] ")
	printList(merged)

	fmt.Println("\n=== [Stock] Best Time to Buy and Sell Stock ===")
	fmt.Printf("[Stock] [7,1,5,3,6,4] → %d\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("[Stock] [7,6,4,3,1] → %d\n", maxProfit([]int{7, 6, 4, 3, 1}))

	fmt.Println("\n=== [Anagram] Valid Anagram ===")
	fmt.Printf("[Anagram] \"anagram\",\"nagaram\" → %v\n", isAnagram("anagram", "nagaram"))
	fmt.Printf("[Anagram] \"rat\",\"car\" → %v\n", isAnagram("rat", "car"))
}

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
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return dummy.Next
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d→", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
