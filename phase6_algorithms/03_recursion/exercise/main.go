package main

import "fmt"

func main() {
	// 演習1: べき乗の高速計算（繰り返し二乗法）
	// base^exp を O(log n) で計算する関数を実装してください
	// ヒント: exp が偶数なら base^exp = (base^(exp/2))^2
	//         exp が奇数なら base^exp = base * base^(exp-1)
	fmt.Println("=== [Power] べき乗の高速計算 ===")
	fmt.Printf("[Power] 2^10 = %d\n", power(2, 10))  // 1024
	fmt.Printf("[Power] 3^5 = %d\n", power(3, 5))    // 243
	fmt.Printf("[Power] 5^0 = %d\n", power(5, 0))    // 1

	// 演習2: 文字列の全順列（Permutations）
	// 与えられた文字列のすべての並び替えを返す関数を実装してください
	// ヒント: バックトラッキング。各位置に残りの文字を1つずつ配置して再帰
	fmt.Println("\n=== [Permute] 文字列の全順列 ===")
	result := permute("abc")
	fmt.Printf("[Permute] \"abc\" の全順列: %v\n", result)
	fmt.Printf("[Permute] 順列数: %d\n", len(result)) // 6
}

// 演習1: power を実装してください
func power(base, exp int) int {
	// ここに実装を書いてください
	// ベースケース: exp == 0 → 1
	// exp が偶数: half = power(base, exp/2) → half * half
	// exp が奇数: base * power(base, exp-1)
	return 0
}

// 演習2: permute を実装してください
func permute(s string) []string {
	// ここに実装を書いてください
	// ヒント: バックトラッキングで全パターンを生成
	// 使用済み文字を記録する used []bool を使う
	return nil
}
