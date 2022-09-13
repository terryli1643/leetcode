package main

import "fmt"

/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.
*/
func main() {
	r := isPalindrome(1211)
	fmt.Println(r)
}

func isPalindrome(x int) bool {
	s := fmt.Sprint(x)
	for i := range s {
		if i < len(s)/2 && i < len(s)-i-1 && s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
