package main

import "fmt"

/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.
*/
func main() {
	r := isPalindrome(1213121)
	fmt.Println(r)
}

func isPalindrome(x int) bool {
	s := fmt.Sprint(x)
	n := len(s)
	for i := range s {
		if i < n/2 && i < n-i-1 && s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
