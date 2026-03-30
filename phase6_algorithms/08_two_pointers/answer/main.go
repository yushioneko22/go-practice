package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== [3Sum] 3Sum ===")
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Printf("[3Sum] %v → %v\n", []int{-1, 0, 1, 2, -1, -4}, result)

	fmt.Println("\n=== [Water] Container With Most Water ===")
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("[Water] %v → %d\n", height, maxArea(height))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		// 重複スキップ
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 重複スキップ
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxWater {
			maxWater = area
		}

		// 低い方を移動（高い方を動かしても面積が減るだけ）
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxWater
}
