package main

import (
	"fmt"
	"strings"
)

func myAtoi(s string) int {
	re, neg := readNum(s)
	var num int = 0
	for i, j := len(re)-1, 1; i >= 0; i, j = i-1, j*10 {
		if string(re[i]) == "0" {
			continue
		}
		num = num + int(re[i]-48)*j
	}
	if neg {
		num = ^num + 1
	}
	return int(int32(num))
}

func readNum(s string) (re []rune, neg bool) {
	m := map[string]bool{
		"+": true, "-": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "0": true,
	}
	s = strings.TrimSpace(s)
	for _, v := range s {
		if _, ok := m[string(v)]; ok {
			re = append(re, v)
		} else {
			break
		}
	}
	if len(re) == 0 {
		return []rune{48}, false
	}
	neg = string(re[0]) == "-"
	if string(re[0]) == "-" || string(re) == "+" {
		re = re[1:]
	}
	return
}

func main() {
	i := myAtoi("words and 987")
	fmt.Println(i)
}
