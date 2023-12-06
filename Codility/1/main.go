package main

import "fmt"

/*

 */

func main() {
	A := []int{2, -2, 3, 0, 4, -7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(Solution(A))

}

var total = 0
var cache = make(map[int]int)

func Solution(A []int) int {
	s := sum(A)
	if s == 0 {
		total++
	} else {
		if _, ok := cache[-s]; ok {
			total++
		} else {
			cache[s] = 1
		}
	}

	N := len(A)
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			s := sum(A[i:j])
			if s == 0 {
				total++
			} else {
				if _, ok := cache[-s]; ok {
					total++
				} else {
					cache[s] = 1
				}
			}
		}
	}
	if total > 1000000000 {
		total = -1
	}
	return total
}

func sum(A []int) int {
	fmt.Println(A)
	s := 0
	for _, v := range A {
		s += v
	}
	return s
}
