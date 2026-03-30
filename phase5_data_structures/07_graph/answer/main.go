package main

import "fmt"

func main() {
	fmt.Println("=== [Islands] 島の数 ===")
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Printf("[Islands] 島の数: %d\n", numIslands(grid))

	fmt.Println("\n=== [Course] コース履修可能判定 ===")
	fmt.Printf("[Course] 2コース, [[1,0],[0,1]] → %v\n",
		canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Printf("[Course] 2コース, [[1,0]] → %v\n",
		canFinish(2, [][]int{{1, 0}}))
}

// numIslands はDFSで連結する陸地を沈めて島を数える
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
		grid[r][c] = '0' // 訪問済みとして沈める
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

// canFinish はカーンのアルゴリズム（トポロジカルソート）でサイクル検出する
// サイクルがなければすべてのコースを履修可能
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 隣接リストと入次数を構築
	adj := make(map[int][]int)
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		adj[prereq] = append(adj[prereq], course)
		inDegree[course]++
	}

	// 入次数0のノードをキューに入れる
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS
	processed := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		processed++

		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return processed == numCourses
}
