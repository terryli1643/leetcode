package main

import "fmt"

// Given a string s, return the longest palindromic substring in s.

func longestPalindrome(s string) string {
	for n := len(s); n > 0; n-- {
		for i, j := 0, n; j <= len(s); i, j = i+1, j+1 {
			if isPalindrome(s[i:j]) {
				return s[i:j]
			}
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func main() {
	sub := longestPalindrome("aba")
	fmt.Println(sub)
}
