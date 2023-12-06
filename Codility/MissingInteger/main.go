package main

import "fmt"

/*
This is a demo task.

Write a function:

class Solution { public int solution(int[] A); }

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
*/
func main() {
	A := []int{1, 3, 6, 4, 1, 2}
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// Implement your solution here
	m := make(map[int]int)
	for _, v := range A {
		if v <= 0 {
			continue
		}
		m[v] = v
	}
	for i := 1; i <= 1000000; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 0
}
