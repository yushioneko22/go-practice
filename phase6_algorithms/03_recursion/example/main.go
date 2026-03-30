package main

import (
	"fmt"
	"strings"
)

func main() {
	// ========================================
	// 階乗
	// ========================================
	fmt.Println("=== [Factorial] 階乗 ===")
	for _, n := range []int{0, 1, 5, 10} {
		fmt.Printf("[Factorial] %d! = %d\n", n, factorial(n))
	}

	// ========================================
	// フィボナッチ（素朴な再帰 vs メモ化）
	// ========================================
	fmt.Println("\n=== [Fibonacci] フィボナッチ数列 ===")
	fmt.Println("[Fibonacci] 素朴な再帰:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("  fib(%d) = %d\n", i, fibNaive(i))
	}

	fmt.Println("[Fibonacci] メモ化版:")
	memo := make(map[int]int)
	for i := 0; i <= 40; i += 10 {
		fmt.Printf("  fib(%d) = %d\n", i, fibMemo(i, memo))
	}

	// ========================================
	// 再帰の可視化
	// ========================================
	fmt.Println("\n=== [Visualize] 再帰の深さを可視化 ===")
	factorialVisualize(5, 0)
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// fibNaive は素朴な再帰（O(2^n) — 遅い！）
func fibNaive(n int) int {
	if n <= 1 {
		return n
	}
	return fibNaive(n-1) + fibNaive(n-2)
}

// fibMemo はメモ化再帰（O(n)）
func fibMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

func factorialVisualize(n int, depth int) int {
	indent := strings.Repeat("  ", depth)
	fmt.Printf("%s[Visualize] factorial(%d) 呼び出し\n", indent, n)
	if n <= 1 {
		fmt.Printf("%s[Visualize] → 1 (ベースケース)\n", indent)
		return 1
	}
	result := n * factorialVisualize(n-1, depth+1)
	fmt.Printf("%s[Visualize] → %d\n", indent, result)
	return result
}
