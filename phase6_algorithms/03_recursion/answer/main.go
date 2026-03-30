package main

import "fmt"

func main() {
	fmt.Println("=== [Power] べき乗の高速計算 ===")
	fmt.Printf("[Power] 2^10 = %d\n", power(2, 10))
	fmt.Printf("[Power] 3^5 = %d\n", power(3, 5))
	fmt.Printf("[Power] 5^0 = %d\n", power(5, 0))

	fmt.Println("\n=== [Permute] 文字列の全順列 ===")
	result := permute("abc")
	fmt.Printf("[Permute] \"abc\" の全順列: %v\n", result)
	fmt.Printf("[Permute] 順列数: %d\n", len(result))
}

// power は繰り返し二乗法で O(log n)
func power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp%2 == 0 {
		half := power(base, exp/2)
		return half * half
	}
	return base * power(base, exp-1)
}

// permute はバックトラッキングで全順列を生成
func permute(s string) []string {
	runes := []rune(s)
	var result []string
	used := make([]bool, len(runes))
	current := make([]rune, 0, len(runes))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(runes) {
			result = append(result, string(current))
			return
		}
		for i := 0; i < len(runes); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			current = append(current, runes[i])
			backtrack()
			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}
