package main

import "fmt"

// TreeNode は二分木のノード
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// ========================================
	// 二分木の構築
	// ========================================
	//       1
	//      / \
	//     2   3
	//    / \ / \
	//   4  5 6  7
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

	// ========================================
	// 3種の走査
	// ========================================
	fmt.Println("=== [Traversal] 二分木の走査 ===")

	fmt.Print("[Traversal] 前順 (Pre-order):  ")
	preorder(root)
	fmt.Println()

	fmt.Print("[Traversal] 中順 (In-order):   ")
	inorder(root)
	fmt.Println()

	fmt.Print("[Traversal] 後順 (Post-order): ")
	postorder(root)
	fmt.Println()

	// ========================================
	// レベル順走査（BFS）
	// ========================================
	fmt.Println("\n=== [BFS] レベル順走査 ===")
	levels := levelOrder(root)
	for i, level := range levels {
		fmt.Printf("[BFS] レベル%d: %v\n", i, level)
	}
}

func preorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val) // 自分 → 左 → 右
	preorder(node.Left)
	preorder(node.Right)
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	fmt.Printf("%d ", node.Val) // 左 → 自分 → 右
	inorder(node.Right)
}

func postorder(node *TreeNode) {
	if node == nil {
		return
	}
	postorder(node.Left)
	postorder(node.Right)
	fmt.Printf("%d ", node.Val) // 左 → 右 → 自分
}

// levelOrder はキューを使ったBFS（レベル順走査）
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
