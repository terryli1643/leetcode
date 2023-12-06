package main

import (
	"fmt"
)

/*
For a given array A of N integers and a sequence S of N integers from the set {−1, 1}, we define val(A, S) as follows:

val(A, S) = |sum{ A[i]*S[i] for i = 0..N−1 }|

(Assume that the sum of zero elements equals zero.)

For a given array A, we are looking for such a sequence S that minimizes val(A,S).

Write a function:

class Solution { public int solution(int[] A); }

that, given an array A of N integers, computes the minimum value of val(A,S) from all possible values of val(A,S) for all possible sequences S of N integers from the set {−1, 1}.

For example, given array:

  A[0] =  1
  A[1] =  5
  A[2] =  2
  A[3] = -2
your function should return 0, since for S = [−1, 1, −1, 1], val(A, S) = 0, which is the minimum possible value.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..20,000];
each element of array A is an integer within the range [−100..100].
*/

func main() {
	// A := []int{1, 5, 2, -2}
	A := []int{2, 3, 4}
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// 初始化 DP 数组
	N := len(A)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}

	// 初始条件
	dp[0][0] = A[0]

	// 状态转移
	for i := 1; i < N; i++ {
		for j := 0; j < N; j++ {
			dp[i][j] = minPos(dp[i-1][j-1]+abs(A[i]-(-1)*j), dp[i-1][j]+abs(A[i]), dp[i][j-1]+abs(A[i]))
		}
	}

	// 返回答案
	return dp[N-1][N-1]
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// func Solution(A []int) int {
// 	// Implement your solution here
// 	n := len(A)
// 	dp := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		dp[i] = make([]int, 2)
// 	}

// 	dp[0][0] = A[0]  //表示正数最小值
// 	dp[0][1] = -A[0] //表示负数最大值

// 	result := Abs(dp[0][0])
// 	for i := 1; i < n; i++ {
// 		a, b, c, d := 0, 0, 0, 0
// 		a = dp[i-1][0] - A[i]
// 		b = dp[i-1][0] + A[i]
// 		c = dp[i-1][1] - A[i]
// 		d = dp[i-1][1] + A[i]
// 		dp[i][0] = minPos(a, b, c, d)
// 		dp[i][1] = maxNeg(a, b, c, d)

// 		result = minPos(Abs(dp[i][0]), Abs(dp[i][1]))
// 	}

// 	return result
// }

// func Abs(i int) int {
// 	if i < 0 {
// 		return -i
// 	}
// 	return i
// }

func minPos(x ...int) int {
	min := -1
	for i := 0; i < len(x); i++ {
		if x[i] < 0 {
			continue
		}
		if min == -1 {
			min = x[i]
			continue
		}
		if min > x[i] {
			min = x[i]
		}
	}
	return min
}

// func maxNeg(x ...int) int {
// 	max := 1
// 	for i := 0; i < len(x); i++ {
// 		if x[i] > 0 {
// 			continue
// 		}
// 		if max == 1 {
// 			max = x[i]
// 			continue
// 		}
// 		if max < x[i] {
// 			max = x[i]
// 		}
// 	}
// 	return max
// }
