package main

import (
	"fmt"
)

/*
We are given two strings P and Q, each consisting of N lowercase English letters. For each position in the strings, we have to choose one letter from either P or Q, in order to construct a new string S, such that the number of distinct letters in S is minimal. Our task is to find the number of distinct letters in the resulting string S.

For example, if P = "ca" and Q = "ab", S can be equal to: "ca", "cb", "aa" or "ab". String "aa" has only one distinct letter ('a'), so the answer is 1 (which is minimal among those strings).

Write a function:

class Solution { public int solution(String P, String Q); }

that, given two strings P and Q, both of length N, returns the minimum number of distinct letters of a string S, that can be constructed from P and Q as described above.

Examples:

1. Given P = "abc", Q = "bcd", your function should return 2. All possible strings S that can be constructed are: "abc", "abd", "acc", "acd", "bbc", "bbd", "bcc", "bcd". The minimum number of distinct letters is 2, which be obtained by constructing the following strings: "acc", "bbc", "bbd", "bcc".

2. Given P = "axxz", Q = "yzwy", your function should return 2. String S must consist of at least two distinct letters in this case. We can construct S = "yxxy", where the number of distinct letters is equal to 2, and this is the only optimal solution.

3. Given P = "bacad", Q = "abada", your function should return 1. We can choose the letter 'a' in each position, so S can be equal to "aaaaa".

4. Given P = "amz", Q = "amz", your function should return 3. The input strings are identical, so the only possible S that can be constructed is "amz", and its number of distinct letters is 3.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..200,000];
strings P and Q are both of length N;
strings P and Q are made only of lowercase letters (a−z);
strings P and Q contain a total of at most 20 distinct letters.

*/

func main() {
	fmt.Println(Solution("axxz", "yzwy"))
}

func Solution(P string, Q string) int {
	// Implement your solution here
	return A(P, Q, 0, "")
}

func A(P string, Q string, i int, r string) int {
	if i == len(P)-1 {
		c1 := count(r + string(P[i]))
		c2 := count(r + string(Q[i]))
		if c1 > c2 {
			return c1
		} else {
			return c2
		}
	}

	r1 := A(P, Q, i+1, r+string(P[i]))
	r2 := A(P, Q, i+1, r+string(Q[i]))

	if r1 < r2 {
		return r1
	} else {
		return r2
	}
}

func count(s string) int {
	min := 0
	m := make(map[rune]int)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			min++
			m[v] = 1
		}
	}
	return min
}
