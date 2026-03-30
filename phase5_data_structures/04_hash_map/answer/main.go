package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== [Anagram] アナグラムのグループ化 ===")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	for _, group := range groups {
		sort.Strings(group)
		fmt.Printf("[Anagram] %v\n", group)
	}

	fmt.Println("\n=== [FirstUniq] 最初の重複しない文字 ===")
	fmt.Printf("[FirstUniq] \"leetcode\" → %d\n", firstUniqChar("leetcode"))
	fmt.Printf("[FirstUniq] \"loveleetcode\" → %d\n", firstUniqChar("loveleetcode"))
	fmt.Printf("[FirstUniq] \"aabb\" → %d\n", firstUniqChar("aabb"))
}

// groupAnagrams は文字列をソートしたものをキーにしてグループ化する
// "eat" → "aet", "tea" → "aet" なので同じグループに入る
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		// 文字列をソートしてキーにする
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		key := string(runes)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

// firstUniqChar は2パスで最初のユニーク文字を見つける
// 1パス目: 各文字の出現回数をカウント
// 2パス目: 最初に出現回数1の文字のインデックスを返す
func firstUniqChar(s string) int {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	for i, ch := range s {
		if freq[ch] == 1 {
			return i
		}
	}
	return -1
}
