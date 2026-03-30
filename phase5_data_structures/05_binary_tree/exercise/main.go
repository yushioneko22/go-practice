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

	// 演習1: 木の最大深さ
	// ルートから最も遠い葉ノードまでの辺の数を返す関数を実装してください
	// ヒント: 再帰で左右の深さの大きい方+1を返す
	fmt.Println("=== [MaxDepth] 木の最大深さ ===")
	fmt.Printf("[MaxDepth] 最大深さ: %d\n", maxDepth(root))

	// 演習2: 木の反転（ミラー化）
	// すべてのノードの左右の子を入れ替える関数を実装してください
	// ヒント: 再帰的に各ノードのLeftとRightをswapする
	fmt.Println("\n=== [Invert] 木の反転 ===")
	fmt.Print("[Invert] 反転前 (中順): ")
	inorder(root)
	fmt.Println()
	invertTree(root)
	fmt.Print("[Invert] 反転後 (中順): ")
	inorder(root)
	fmt.Println()

	// 演習3: レベル順走査
	// キュー（BFS）を使ってレベルごとにノードの値を集める関数を実装してください
	fmt.Println("\n=== [LevelOrder] レベル順走査 ===")
	// 反転を戻してからテスト
	invertTree(root)
	levels := levelOrder(root)
	fmt.Printf("[LevelOrder] %v\n", levels)
}

// 演習1: maxDepth を実装してください
func maxDepth(root *TreeNode) int {
	// ここに実装を書いてください
	// ヒント: nilなら0、それ以外はmax(左の深さ, 右の深さ)+1
	return 0
}

// 演習2: invertTree を実装してください
func invertTree(root *TreeNode) *TreeNode {
	// ここに実装を書いてください
	// ヒント: node.Left, node.Right = node.Right, node.Left を再帰的に
	return root
}

// 演習3: levelOrder を実装してください
func levelOrder(root *TreeNode) [][]int {
	// ここに実装を書いてください
	// ヒント: キュー（スライス）を使い、各レベルのサイズ分だけノードを処理する
	return nil
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	fmt.Printf("%d ", node.Val)
	inorder(node.Right)
}
