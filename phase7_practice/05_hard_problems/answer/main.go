package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("=== [Rain] Trapping Rain Water ===")
	fmt.Printf("[Rain] [0,1,0,2,1,0,1,3,2,1,2,1] → %d\n",
		trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Printf("[Rain] [4,2,0,3,2,5] → %d\n",
		trap([]int{4, 2, 0, 3, 2, 5}))

	fmt.Println("\n=== [MinWindow] Minimum Window Substring ===")
	fmt.Printf("[MinWindow] \"ADOBECODEBANC\",\"ABC\" → \"%s\"\n",
		minWindow("ADOBECODEBANC", "ABC"))

	fmt.Println("\n=== [Serialize] Serialize and Deserialize Binary Tree ===")
	root := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
	}
	data := serialize(root)
	fmt.Printf("[Serialize] 文字列: %s\n", data)
	decoded := deserialize(data)
	fmt.Printf("[Serialize] 復元後: %s\n", serialize(decoded))
}

// trap はTwo Pointersで雨水量を計算（O(n)時間、O(1)空間）
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}

func minWindow(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	required := len(need)
	have := 0
	window := make(map[byte]int)
	bestLeft, bestLen := 0, len(s)+1
	left := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		window[ch]++
		if need[ch] > 0 && window[ch] == need[ch] {
			have++
		}
		for have == required {
			if right-left+1 < bestLen {
				bestLen = right - left + 1
				bestLeft = left
			}
			lch := s[left]
			window[lch]--
			if need[lch] > 0 && window[lch] < need[lch] {
				have--
			}
			left++
		}
	}
	if bestLen == len(s)+1 {
		return ""
	}
	return s[bestLeft : bestLeft+bestLen]
}

// serialize は前順走査で文字列化
func serialize(root *TreeNode) string {
	var parts []string
	var build func(node *TreeNode)
	build = func(node *TreeNode) {
		if node == nil {
			parts = append(parts, "#")
			return
		}
		parts = append(parts, strconv.Itoa(node.Val))
		build(node.Left)
		build(node.Right)
	}
	build(root)
	return strings.Join(parts, ",")
}

// deserialize は文字列から二分木を復元
func deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")
	idx := 0

	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(parts) || parts[idx] == "#" {
			idx++
			return nil
		}
		val, _ := strconv.Atoi(parts[idx])
		idx++
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}

	return build()
}
