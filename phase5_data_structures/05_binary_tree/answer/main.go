package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println("=== [MaxDepth] 木の最大深さ ===")
	fmt.Printf("[MaxDepth] 最大深さ: %d\n", maxDepth(root))

	fmt.Println("\n=== [Invert] 木の反転 ===")
	fmt.Print("[Invert] 反転前 (中順): ")
	inorder(root)
	fmt.Println()
	invertTree(root)
	fmt.Print("[Invert] 反転後 (中順): ")
	inorder(root)
	fmt.Println()

	invertTree(root)
	fmt.Println("\n=== [LevelOrder] レベル順走査 ===")
	levels := levelOrder(root)
	fmt.Printf("[LevelOrder] %v\n", levels)
}

// maxDepth は再帰的に木の最大深さを求める
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// invertTree は二分木の全ノードの左右を反転する
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// levelOrder はBFSでレベルごとのノード値を返す
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	fmt.Printf("%d ", node.Val)
	inorder(node.Right)
}
