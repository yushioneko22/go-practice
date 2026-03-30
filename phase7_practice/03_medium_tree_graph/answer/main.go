package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("=== [ValidBST] Validate Binary Search Tree ===")
	validBST := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	invalidBST := &TreeNode{5,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}},
	}
	fmt.Printf("[ValidBST] [2,1,3] → %v\n", isValidBST(validBST))
	fmt.Printf("[ValidBST] [5,1,4,3,6] → %v\n", isValidBST(invalidBST))

	fmt.Println("\n=== [Islands] Number of Islands ===")
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Printf("[Islands] 島の数: %d\n", numIslands(grid))

	fmt.Println("\n=== [Course] Course Schedule ===")
	fmt.Printf("[Course] 2, [[1,0]] → %v\n", canFinish(2, [][]int{{1, 0}}))
	fmt.Printf("[Course] 2, [[1,0],[0,1]] → %v\n", canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func isValidBST(root *TreeNode) bool {
	return validateBST(root, math.MinInt64, math.MaxInt64)
}

func validateBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return validateBST(node.Left, min, node.Val) && validateBST(node.Right, node.Val, max)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
		inDegree[pre[0]]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	processed := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		processed++
		for _, next := range adj[node] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return processed == numCourses
}
