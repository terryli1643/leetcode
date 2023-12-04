package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) int {
	s := fmt.Sprint(x)
	sb := strings.Builder{}
	if x < 0 {
		sb.Write([]byte("-"))
		for i := len(s) - 1; i >= 1; i-- {
			sb.Write([]byte{s[i]})
		}
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			sb.Write([]byte{s[i]})
		}
	}
	r, _ := strconv.ParseInt(sb.String(), 10, 64)
	if r != int64(int32(r)) {
		return 0
	}
	return int(r)
}

func main() {
	fmt.Println(reverse(-123))
}
