package main

import (
	"fmt"
	"sort"
)

func main() {
	// 演習1: アナグラムのグループ化
	// 同じ文字の組み合わせからなる単語同士をグループにまとめてください
	// ヒント: 各単語をソートしたものをキーにしてmapでグループ化する
	fmt.Println("=== [Anagram] アナグラムのグループ化 ===")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	for _, group := range groups {
		sort.Strings(group)
		fmt.Printf("[Anagram] %v\n", group)
	}

	// 演習2: 最初の重複しない文字
	// 文字列中で最初に1回しか出現しない文字のインデックスを返してください
	// 見つからない場合は -1 を返す
	// ヒント: 2パス（1回目で頻度カウント、2回目で最初の1回文字を探す）
	fmt.Println("\n=== [FirstUniq] 最初の重複しない文字 ===")
	fmt.Printf("[FirstUniq] \"leetcode\" → %d\n", firstUniqChar("leetcode"))         // 0
	fmt.Printf("[FirstUniq] \"loveleetcode\" → %d\n", firstUniqChar("loveleetcode")) // 2
	fmt.Printf("[FirstUniq] \"aabb\" → %d\n", firstUniqChar("aabb"))                 // -1
}

// 演習1: groupAnagrams を実装してください
func groupAnagrams(strs []string) [][]string {
	// ここに実装を書いてください
	// ヒント: 文字列をソートしてキーにする
	// "eat" → "aet", "tea" → "aet" → 同じグループ
	return nil
}

// 演習2: firstUniqChar を実装してください
func firstUniqChar(s string) int {
	// ここに実装を書いてください
	// ヒント: map[rune]int で頻度カウント → 再度文字列を走査して頻度1の文字を探す
	return -1
}
