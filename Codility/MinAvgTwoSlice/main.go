package main

import (
	"fmt"
)

/*
A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice contains at least two elements). The average of a slice (P, Q) is the sum of A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

For example, array A such that:

	A[0] = 4
	A[1] = 2
	A[2] = 2
	A[3] = 5
	A[4] = 1
	A[5] = 5
	A[6] = 8

contains the following example slices:

slice (1, 2), whose average is (2 + 2) / 2 = 2;
slice (3, 4), whose average is (5 + 1) / 2 = 3;
slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.
The goal is to find the starting position of a slice whose average is minimal.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the starting position of the slice with the minimal average. If there is more than one slice with a minimal average, you should return the smallest starting position of such a slice.

For example, given array A such that:

	A[0] = 4
	A[1] = 2
	A[2] = 2
	A[3] = 5
	A[4] = 1
	A[5] = 5
	A[6] = 8

the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−10,000..10,000].
*/
func main() {

	// A := []int{4, 2, 2, 5, 1, 5, 8}
	// A := []int{-3, -5, -8, -4, -10}
	// A := []int{5, 6, 3, 4, 9}
	A := []int{10, 10, -1, 2, 4, -1, 2, -1}
	// A := []int{10000, -10000}
	fmt.Println(Solution(A))
}

// 不对
// func Solution(A []int) int {
// 	N := len(A)
// 	var dp = make([]int, N+1)

// 	for p := 1; p < N+1; p++ {
// 		sum := dp[p-1] + A[p-1]
// 		dp[p] = sum
// 	}
// 	fmt.Println(dp)

// 	min := float64(10000)
// 	pos := 0
// 	for sliceLen := 2; sliceLen <= 3; sliceLen++ {
// 		for i := 0; i < N-sliceLen+1; i++ {
// 			avg := float64(dp[i+sliceLen]-dp[i]) / float64(2)
// 			if min > avg {
// 				min = avg
// 				pos = i
// 			}
// 		}
// 	}
// 	return pos
// }

// 暴力求解
func Solution(A []int) int {
	N := len(A)
	mem := make(map[int]int)
	min := float64(10000)
	start := 0
	for p := 0; p < N-1; p++ {
		sum := A[p]
		for q := p + 1; q < N; q++ {
			sum = sum + A[q]
			avg := float64(sum) / float64(q-p+1)
			if min > avg {
				min = avg
				start = p
				mem[start] = q
			}
		}
	}
	return start
}
