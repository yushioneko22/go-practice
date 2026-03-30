package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== [LongestSub] Longest Substring Without Repeating Characters ===")
	fmt.Printf("[LongestSub] \"abcabcbb\" → %d\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("[LongestSub] \"bbbbb\" → %d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("[LongestSub] \"pwwkew\" → %d\n", lengthOfLongestSubstring("pwwkew"))

	fmt.Println("\n=== [3Sum] 3Sum ===")
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("[3Sum] %v → %v\n", []int{-1, 0, 1, 2, -1, -4}, threeSum(nums))

	fmt.Println("\n=== [GroupAnagram] Group Anagrams ===")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	for _, g := range groups {
		fmt.Printf("[GroupAnagram] %v\n", g)
	}
}

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLen, left := 0, 0
	for right := 0; right < len(s); right++ {
		if idx, ok := charIndex[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		charIndex[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
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

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		key := string(runes)
		groups[key] = append(groups[key], s)
	}
	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}
