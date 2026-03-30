package main

import "fmt"

func main() {
	// ========================================
	// 隣接リストによるグラフの表現
	// ========================================
	fmt.Println("=== [Graph] グラフの構築 ===")
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 4},
		3: {1, 4, 5},
		4: {2, 3},
		5: {3},
	}

	for node, neighbors := range graph {
		fmt.Printf("[Graph] ノード%d → %v\n", node, neighbors)
	}

	// ========================================
	// BFS（幅優先探索）
	// ========================================
	fmt.Println("\n=== [BFS] 幅優先探索 ===")
	fmt.Print("[BFS] 探索順: ")
	bfs(graph, 0)
	fmt.Println()

	// ========================================
	// DFS（深さ優先探索）
	// ========================================
	fmt.Println("\n=== [DFS] 深さ優先探索 ===")
	fmt.Print("[DFS] 探索順: ")
	visited := make(map[int]bool)
	dfs(graph, 0, visited)
	fmt.Println()
}

// bfs はキューを使った幅優先探索
func bfs(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

// dfs は再帰による深さ優先探索
func dfs(graph map[int][]int, node int, visited map[int]bool) {
	visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited)
		}
	}
}
