package main

import "fmt"

func main() {
	// 演習1: 重複なし最長部分文字列
	// 重複する文字を含まない最長の部分文字列の長さを返してください
	// ヒント: mapで各文字の最後の出現位置を記録 + 可変ウィンドウ
	fmt.Println("=== [NoRepeat] 重複なし最長部分文字列 ===")
	fmt.Printf("[NoRepeat] \"abcabcbb\" → %d\n", lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Printf("[NoRepeat] \"bbbbb\" → %d\n", lengthOfLongestSubstring("bbbbb"))       // 1
	fmt.Printf("[NoRepeat] \"pwwkew\" → %d\n", lengthOfLongestSubstring("pwwkew"))     // 3

	// 演習2: 最小被覆部分文字列（Minimum Window Substring）
	// 文字列sの中で、文字列tの全文字を含む最短の部分文字列を返してください
	// ヒント: tの文字頻度をカウント → ウィンドウを拡張して全文字を含むまで → 左端を縮小
	fmt.Println("\n=== [MinWindow] 最小被覆部分文字列 ===")
	fmt.Printf("[MinWindow] s=\"ADOBECODEBANC\", t=\"ABC\" → \"%s\"\n",
		minWindow("ADOBECODEBANC", "ABC")) // "BANC"
	fmt.Printf("[MinWindow] s=\"a\", t=\"a\" → \"%s\"\n",
		minWindow("a", "a")) // "a"
}

// 演習1: lengthOfLongestSubstring を実装してください
func lengthOfLongestSubstring(s string) int {
	// ここに実装を書いてください
	// ヒント: left ポインタと map[byte]int（文字→最後の位置）を使う
	return 0
}

// 演習2: minWindow を実装してください
func minWindow(s, t string) string {
	// ここに実装を書いてください
	// ヒント: need map（tの文字頻度）, have変数（条件を満たした文字数）
	// 右端拡張 → haveがtの文字種類数に達したら → 左端縮小
	return ""
}
