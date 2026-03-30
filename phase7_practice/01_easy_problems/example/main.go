package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("=== [Palindrome] Valid Palindrome ===")
	tests := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"",
		" ",
	}
	for _, s := range tests {
		fmt.Printf("[Palindrome] \"%s\" → %v\n", s, isPalindrome(s))
	}
}

// isPalindrome は英数字のみを対象に回文判定する
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		// 英数字以外をスキップ
		for left < right && !isAlphaNum(rune(s[left])) {
			left++
		}
		for left < right && !isAlphaNum(rune(s[right])) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
