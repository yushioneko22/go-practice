package main

import (
	"fmt"
	"sort"
)

func main() {
	// ========================================
	// mapの基本操作
	// ========================================
	fmt.Println("=== [Map] 基本操作 ===")
	m := map[string]int{
		"apple":  3,
		"banana": 5,
		"cherry": 2,
	}

	// 追加・更新
	m["date"] = 7
	fmt.Printf("[Map] 追加後: %v\n", m)

	// 検索（comma ok idiom）
	if val, ok := m["banana"]; ok {
		fmt.Printf("[Map] banana = %d\n", val)
	}
	if _, ok := m["grape"]; !ok {
		fmt.Println("[Map] grape は存在しません")
	}

	// 削除
	delete(m, "cherry")
	fmt.Printf("[Map] cherry削除後: %v\n", m)

	// ========================================
	// 頻度カウント（面接超頻出パターン）
	// ========================================
	fmt.Println("\n=== [Freq] 頻度カウント ===")
	word := "programming"
	freq := make(map[rune]int)
	for _, ch := range word {
		freq[ch]++
	}

	// ソートして表示（mapの反復順序は不定のため）
	keys := make([]string, 0, len(freq))
	for ch := range freq {
		keys = append(keys, string(ch))
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("[Freq] '%s' → %d回\n", k, freq[rune(k[0])])
	}

	// ========================================
	// 重複検出
	// ========================================
	fmt.Println("\n=== [Dup] 重複検出 ===")
	nums := []int{1, 3, 5, 3, 7, 5, 9}
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			fmt.Printf("[Dup] 重複発見: %d\n", n)
		}
		seen[n] = true
	}
}
