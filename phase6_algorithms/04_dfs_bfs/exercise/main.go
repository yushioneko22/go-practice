package main

import "fmt"

func main() {
	// 演習1: 迷路の最短経路（BFS）
	// 0=通路、1=壁のグリッドで、左上から右下への最短距離を求めてください
	fmt.Println("=== [Maze] 迷路の最短経路 ===")
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}
	dist := shortestPath(maze)
	fmt.Printf("[Maze] 最短距離: %d\n", dist) // 期待値: 7

	// 演習2: 全経路の列挙（DFS + バックトラッキング）
	// (0,0) から (2,2) へのすべての経路を列挙してください（右と下のみ移動可）
	fmt.Println("\n=== [AllPaths] 全経路の列挙 ===")
	paths := allPaths(3, 3)
	for _, path := range paths {
		fmt.Printf("[AllPaths] %v\n", path)
	}
	fmt.Printf("[AllPaths] 総経路数: %d\n", len(paths))
}

// 演習1: shortestPath を実装してください
func shortestPath(grid [][]int) int {
	// ここに実装を書いてください
	// ヒント: BFS（キュー）を使う
	// (0,0) からスタート、(rows-1, cols-1) がゴール
	// 4方向移動: dx := []int{0,0,1,-1}, dy := []int{1,-1,0,0}
	return -1
}

// 演習2: allPaths を実装してください
func allPaths(rows, cols int) [][2]int {
	// ここに実装を書いてください
	// ヒント: DFS + バックトラッキング
	// 右(0,1)と下(1,0)のみ移動可能
	// 各経路を [][2]int で記録...ではなく、全経路のリストを返す
	// （簡略化のため、経路数だけ返してもOK）
	return nil
}
