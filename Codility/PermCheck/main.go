package main

import "fmt"

/*
A non-empty array A consisting of N integers is given.

A permutation is a sequence containing each element from 1 to N once, and only once.

For example, array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
is a permutation, but array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
is not a permutation, because value 2 is missing.

The goal is to check whether array A is a permutation.

Write a function:

func Solution(A []int) int

that, given an array A, returns 1 if array A is a permutation and 0 if it is not.

For example, given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
the function should return 1.

Given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..1,000,000,000].
*/

func main() {
	A := []int{2, 3, 4, 1}
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// Implement your solution here
	min, max := A[0], A[0]
	mem := make(map[int]int)
	for _, v := range A {
		if _, ok := mem[v]; ok {
			return 0
		}
		mem[v] = 1
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	if min > 1 {
		return 0
	}
	if max > len(A) {
		return 0
	}

	return 1
}
