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
	// math パッケージの使用を明示
	_ = math.MinInt64

	// 演習1: Validate Binary Search Tree
	// BSTの条件を満たしているかを判定（左部分木全体 < 親 < 右部分木全体）
	fmt.Println("=== [ValidBST] Validate Binary Search Tree ===")
	validBST := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	invalidBST := &TreeNode{5,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}},
	}
	fmt.Printf("[ValidBST] [2,1,3] → %v\n", isValidBST(validBST))   // true
	fmt.Printf("[ValidBST] [5,1,4,3,6] → %v\n", isValidBST(invalidBST)) // false

	// 演習2: Number of Islands
	// '1'=陸、'0'=水のグリッドで島の数を数える
	fmt.Println("\n=== [Islands] Number of Islands ===")
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Printf("[Islands] 島の数: %d\n", numIslands(grid)) // 3

	// 演習3: Course Schedule
	// 前提条件付きのコースをすべて履修できるか判定
	fmt.Println("\n=== [Course] Course Schedule ===")
	fmt.Printf("[Course] 2, [[1,0]] → %v\n", canFinish(2, [][]int{{1, 0}}))         // true
	fmt.Printf("[Course] 2, [[1,0],[0,1]] → %v\n", canFinish(2, [][]int{{1, 0}, {0, 1}})) // false
}

// 演習1: isValidBST を実装してください
func isValidBST(root *TreeNode) bool {
	// ヒント: helper(node, min, max) で上限/下限を渡す再帰
	// 初期値は math.MinInt64, math.MaxInt64
	return false
}

// 演習2: numIslands を実装してください
func numIslands(grid [][]byte) int {
	// ヒント: '1'を見つけたらDFSで隣接する'1'を'0'にする
	return 0
}

// 演習3: canFinish を実装してください
func canFinish(numCourses int, prerequisites [][]int) bool {
	// ヒント: トポロジカルソート（入次数 + BFS）
	return false
}
