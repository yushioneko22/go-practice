package main

import "fmt"

func main() {
	// ========================================
	// BFS: グリッド上の最短経路
	// ========================================
	fmt.Println("=== [BFS] グリッド上の最短経路 ===")
	// 0=通路, 1=壁
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}

	dist := bfsGrid(grid, [2]int{0, 0}, [2]int{4, 3})
	fmt.Printf("[BFS] (0,0) → (4,3) の最短距離: %d\n", dist)

	// ========================================
	// DFS: 到達可能なセルの列挙
	// ========================================
	fmt.Println("\n=== [DFS] 到達可能なセルの列挙 ===")
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	count := dfsCount(grid, 0, 0, visited)
	fmt.Printf("[DFS] (0,0) から到達可能なセル数: %d\n", count)
}

type Point struct {
	row, col, dist int
}

func bfsGrid(grid [][]int, start, end [2]int) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	queue := []Point{{start[0], start[1], 0}}
	visited[start[0]][start[1]] = true

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.row == end[0] && p.col == end[1] {
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
	return -1 // 到達不可能
}

func dfsCount(grid [][]int, row, col int, visited [][]bool) int {
	rows, cols := len(grid), len(grid[0])
	if row < 0 || row >= rows || col < 0 || col >= cols ||
		visited[row][col] || grid[row][col] == 1 {
		return 0
	}

	visited[row][col] = true
	count := 1

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	for i := 0; i < 4; i++ {
		count += dfsCount(grid, row+dx[i], col+dy[i], visited)
	}
	return count
}
