package main

import (
	"fmt"
	"strings"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

*/

const (
	up   = -1
	down = 1
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := make([][]string, numRows)
	move(s, 0, 0, down, numRows, result)
	sb := strings.Builder{}
	for i := range result {
		for j := range result[i] {
			sb.Write([]byte(result[i][j]))
		}
	}
	return sb.String()
}

func move(s string, p int, i int, dir int, numRows int, result [][]string) {
	if p >= len(s) {
		return
	}
	result[i] = append(result[i], string(s[p]))
	if i == 0 {
		move(s, p+1, i+down, down, numRows, result)
	} else if i == numRows-1 {
		move(s, p+1, i+up, up, numRows, result)
	} else {
		move(s, p+1, i+dir, dir, numRows, result)
	}
}

func main() {
	re := convert("AB", 1)
	fmt.Println(re)
}
