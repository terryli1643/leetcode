package main

import "fmt"

func main() {
	input := "III"
	x := romanToInt(input)
	fmt.Println(x)
}

var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func romanToInt(s string) int {
	result := m[string(s[len(s)-1:])]
	for i := 0; i < len(s)-1; i++ {
		li := string(s[i])
		li1 := string(s[i+1])
		if li == "I" && li1 == "X" ||
			li == "I" && li1 == "V" ||
			li == "X" && li1 == "L" ||
			li == "X" && li1 == "C" ||
			li == "C" && li1 == "D" ||
			li == "C" && li1 == "M" {
			result = result - m[li]
		} else {
			result = result + m[li]
		}
	}
	return result
}
