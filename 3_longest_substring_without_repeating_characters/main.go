package main

import "fmt"

//Given a string s, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	scaned := make(map[byte]int)
	scaned[s[0]] = 0
	length := 0
	for i, j := 0, 1; j < len(s); j++ {
		if p, ok := scaned[s[j]]; ok {
			if i < p+1 {
				i = p + 1
			}
		}
		scaned[s[j]] = j
		if (j - i + 1) > length {
			length = j - i + 1
		}
	}
	return length
}

func main() {
	s := "abba"
	i := lengthOfLongestSubstring(s)
	fmt.Println(i)
}
