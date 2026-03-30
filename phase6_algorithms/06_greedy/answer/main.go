package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== [Schedule] 区間スケジューリング ===")
	meetings := [][]int{
		{1, 4}, {2, 3}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9},
	}
	count := maxMeetings(meetings)
	fmt.Printf("[Schedule] 最大会議数: %d\n", count)

	fmt.Println("\n=== [Jump] ジャンプゲーム ===")
	fmt.Printf("[Jump] [2,3,1,1,4] → %v\n", canJump([]int{2, 3, 1, 1, 4}))
	fmt.Printf("[Jump] [3,2,1,0,4] → %v\n", canJump([]int{3, 2, 1, 0, 4}))
}

func maxMeetings(meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][1] < meetings[j][1]
	})

	count := 0
	lastEnd := -1
	for _, m := range meetings {
		if m[0] >= lastEnd {
			count++
			lastEnd = m[1]
		}
	}
	return count
}

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return true
}
