package main

import "fmt"

func main() {
	// 演習1: Trapping Rain Water
	// 高さの配列で、雨水が溜まる量を求める
	// ヒント: 各位置で「左側の最大高さ」と「右側の最大高さ」の小さい方 - 自分の高さ
	// Two Pointersアプローチ: O(n)時間、O(1)空間
	fmt.Println("=== [Rain] Trapping Rain Water ===")
	fmt.Printf("[Rain] [0,1,0,2,1,0,1,3,2,1,2,1] → %d\n",
		trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Printf("[Rain] [4,2,0,3,2,5] → %d\n",
		trap([]int{4, 2, 0, 3, 2, 5})) // 9

	// 演習2: Minimum Window Substring
	// sの中でtの全文字を含む最短部分文字列を返す
	// ヒント: スライディングウィンドウ + 2つのmap
	fmt.Println("\n=== [MinWindow] Minimum Window Substring ===")
	fmt.Printf("[MinWindow] \"ADOBECODEBANC\",\"ABC\" → \"%s\"\n",
		minWindow("ADOBECODEBANC", "ABC")) // "BANC"

	// 演習3: Serialize and Deserialize Binary Tree
	// 二分木を文字列に変換し、文字列から二分木を復元する
	// ヒント: 前順走査で文字列化、nullは"#"で表現
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 演習1: trap を実装してください
func trap(height []int) int {
	// ヒント: Two Pointers。left, rightから中央に向かう
	// leftMax, rightMaxを追跡し、低い方の水量を計算
	return 0
}

// 演習2: minWindow を実装してください
func minWindow(s, t string) string {
	// ヒント: needマップ(tの文字頻度)、windowマップ、have/required変数
	return ""
}

// 演習3: serialize / deserialize を実装してください
func serialize(root *TreeNode) string {
	// ヒント: 前順走査、nilは"#"、カンマ区切り
	return ""
}

func deserialize(data string) *TreeNode {
	// ヒント: 文字列をカンマで分割、インデックスを再帰で進める
	return nil
}
