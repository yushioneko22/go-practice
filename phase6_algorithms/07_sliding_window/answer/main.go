package main

import "fmt"

func main() {
	fmt.Println("=== [NoRepeat] 重複なし最長部分文字列 ===")
	fmt.Printf("[NoRepeat] \"abcabcbb\" → %d\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("[NoRepeat] \"bbbbb\" → %d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("[NoRepeat] \"pwwkew\" → %d\n", lengthOfLongestSubstring("pwwkew"))

	fmt.Println("\n=== [MinWindow] 最小被覆部分文字列 ===")
	fmt.Printf("[MinWindow] s=\"ADOBECODEBANC\", t=\"ABC\" → \"%s\"\n",
		minWindow("ADOBECODEBANC", "ABC"))
	fmt.Printf("[MinWindow] s=\"a\", t=\"a\" → \"%s\"\n",
		minWindow("a", "a"))
}

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int) // 文字 → 最後の出現インデックス
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, ok := charIndex[s[right]]; ok && idx >= left {
			left = idx + 1 // 重複文字の次の位置に左端を移動
		}
		charIndex[s[right]] = right
		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func minWindow(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// tの文字頻度を記録
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required := len(need) // 満たすべき文字の種類数
	have := 0             // 現在満たしている文字の種類数
	window := make(map[byte]int)

	bestLeft, bestLen := 0, len(s)+1
	left := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		window[ch]++

		if need[ch] > 0 && window[ch] == need[ch] {
			have++
		}

		// 条件を満たしている間、左端を縮小
		for have == required {
			length := right - left + 1
			if length < bestLen {
				bestLen = length
				bestLeft = left
			}

			leftCh := s[left]
			window[leftCh]--
			if need[leftCh] > 0 && window[leftCh] < need[leftCh] {
				have--
			}
			left++
		}
	}

	if bestLen == len(s)+1 {
		return ""
	}
	return s[bestLeft : bestLeft+bestLen]
}
