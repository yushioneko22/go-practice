package main

import "fmt"

type Point struct {
	row, col, dist int
}

func main() {
	fmt.Println("=== [Maze] 迷路の最短経路 ===")
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
	}
	dist := shortestPath(maze)
	fmt.Printf("[Maze] 最短距離: %d\n", dist)

	fmt.Println("\n=== [AllPaths] 全経路の列挙 ===")
	paths := allPaths(3, 3)
	for _, path := range paths {
		fmt.Printf("[AllPaths] %v\n", path)
	}
	fmt.Printf("[AllPaths] 総経路数: %d\n", len(paths))
}

func shortestPath(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[rows-1][cols-1] == 1 {
		return -1
	}

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	queue := []Point{{0, 0, 0}}
	visited[0][0] = true

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.row == rows-1 && p.col == cols-1 {
			return p.dist
		}

		for i := 0; i < 4; i++ {
			nr, nc := p.row+dx[i], p.col+dy[i]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
				!visited[nr][nc] && grid[nr][nc] == 0 {
				visited[nr][nc] = true
				queue = append(queue, Point{nr, nc, p.dist + 1})
			}
		}
	}
	return -1
}

func allPaths(rows, cols int) [][2]int {
	var result [][2]int
	var path [][2]int

	var dfs func(r, c int)
	dfs = func(r, c int) {
		path = append(path, [2]int{r, c})

		if r == rows-1 && c == cols-1 {
			// ゴール到達 → パスのコピーを結果に追加
			result = append(result, [2]int{len(path), 0}) // 経路数カウント用
			pathCopy := make([][2]int, len(path))
			copy(pathCopy, path)
			// 経路全体を表示
			fmt.Printf("  経路: %v\n", pathCopy)
		} else {
			// 右に移動
			if c+1 < cols {
				dfs(r, c+1)
			}
			// 下に移動
			if r+1 < rows {
				dfs(r+1, c)
			}
		}

		path = path[:len(path)-1] // バックトラック
	}

	dfs(0, 0)
	return result
}
